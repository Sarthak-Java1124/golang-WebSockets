package users

import (
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
	err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
	}
	 res , err := h.s.CreateUser(c.Request.Context() , &u)
	 if err != nil {
		log.Fatal("The error in create user service is : " , err)
	 }
c.JSON(http.StatusOK , gin.H{"data": res})

}
