package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jdiego0102/gobank/db/sqlc"
	"github.com/jdiego0102/gobank/token"
)

type transferRequest struct {
	FromCuentaID int64  `json:"from_cuenta_id" binding:"required,min=1"`
	ToCuentaID   int64  `json:"to_cuenta_id" binding:"required,min=1"`
	Monto        int64  `json:"monto" binding:"required,gt=0"`
	Divisa       string `json:"divisa" binding:"required,currency"`
}

func (server *Server) createTranfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromCuentaID, req.Divisa)

	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Propietario != authPayload.Username {
		err := errors.New("la cuenta no pertenece al usuario autenticado")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToCuentaID, req.Divisa)

	if !valid {
		return
	}

	arg := db.TransferTxParams{
		FromCuentaID: req.FromCuentaID,
		ToCuentaID:   req.ToCuentaID,
		Amount:       req.Monto,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Cuentum, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Divisa != currency {
		err := fmt.Errorf("la moneda de la cuenta [%d] no coincide: %s vs %s", accountID, account.Divisa, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}
