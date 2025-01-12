package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
	ErrMethodNotAllowed    = errors.New("Method not allowed")
	ErrConflict            = errors.New("Resource already exists")
	ErrUnprocessableEntity = errors.New("Invalid data")
	ErrNotFound            = errors.New("Resource not found")
	ErrInternalServerError = errors.New("Internal server error")
)

func ErrorHandlingMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			log.Println(e)
		}
		lastErr := c.Errors.Last().Err
		statusCode := c.Writer.Status()
		c.IndentedJSON(statusCode, gin.H{"error": lastErr.Error()})
	}
}
