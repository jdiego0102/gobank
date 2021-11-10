package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jdiego0102/gobank/db/sqlc"
)

// Server entiende todas las peticiones HTTP del servidor bancario
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer crea una nueva instancia del servidor y configura todas las rutas
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Agregar rutas al enturutador
	router.POST("accounts/", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)
	router.PUT("accounts", server.updateAccount)
	router.DELETE("accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start ejecuta el servidor HTTP en la dirección específica de entrada
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
