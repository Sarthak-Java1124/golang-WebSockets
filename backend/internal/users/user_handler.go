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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with username, email, password
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserReq true "User data"
// @Success 200 {object} CreateUserRes
// @Failure 400 {object} map[string]string
// @Router /signup [post]
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

// LoginHandler godoc
// @Summary Login user
// @Description Login with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body LoginUserReq true "Login data"
// @Success 200 {object} LoginUserRes
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *Handlers) LoginHandler(c *gin.Context) {
	var u LoginUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.s.Login(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("access_token", user.accessToken, 3600, "/", "localhost", false, true)
	res := LoginUserRes{
		accessToken: user.accessToken,
		Username:    user.Username,
		ID:          user.ID,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "data": res})

}
