package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	connPool *pgxpool.Pool
	router   *gin.Engine
}
