package users

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Sarthak-Java1124/golang-WebSockets/backend/utils"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	hashedPassword := utils.HashPassword(req.Password)
	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, err
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	u, err := s.Repository.GetUsersByEmail(ctx, req.Email)
	if err != nil {
		log.Println("The error in finding users by email is :", err)
		return nil, err
	}
	err = utils.VerifyHashPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}
	accessToken, err := utils.GenerateJWT(u.Email)
	if err != nil {
		log.Println("The error in generating the jwt is : ", err)
	}
	var loginResponse LoginUserRes
	loginResponse.accessToken = accessToken
	loginResponse.Username = u.Username
	loginResponse.ID = strconv.Itoa(int(u.ID))
	return &loginResponse, nil

}
