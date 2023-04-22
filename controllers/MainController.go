package controllers

import (
	"go/todo/migrate"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	migrate.Migrate()
}
