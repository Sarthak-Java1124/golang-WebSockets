package users

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	s Service
}

func NewHandler(s Service) *Handlers {
	return &Handlers{
		s: s,
	}
}

func (h *Handlers) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.s.CreateUser(context.Background(), &u)
	if err != nil {
		log.Fatal("The error in create user service is : ", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": res})

}
