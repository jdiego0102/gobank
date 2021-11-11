package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jdiego0102/gobank/db/sqlc"
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

	if !server.validAccount(ctx, req.FromCuentaID, req.Divisa) {
		return
	}

	if !server.validAccount(ctx, req.ToCuentaID, req.Divisa) {
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

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Divisa != currency {
		err := fmt.Errorf("la moneda de la cuenta [%d] no coincide: %s vs %s", accountID, account.Divisa, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
