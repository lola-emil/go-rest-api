package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)

	group := r.Group("/users")
	{
		group.GET("/", func(c *gin.Context) {
			users, err := svc.ListUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, users)
		})

		group.POST("/", func(c *gin.Context) {
			var u User
			if err := c.ShouldBindJSON(&u); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := svc.CreateUser(&u); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, u)
		})
	}
}
