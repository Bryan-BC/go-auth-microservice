package services

import (
	"context"
	"net/http"

	"github.com/Bryan-BC/go-auth-microservice/pkg/db"
	"github.com/Bryan-BC/go-auth-microservice/pkg/models"
	"github.com/Bryan-BC/go-auth-microservice/pkg/pb"
	"github.com/Bryan-BC/go-auth-microservice/pkg/utils"
)

type Server struct {
	DBPointer *db.DB
	JWT       *utils.JWTWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.DBPointer.DataBase.Where(&models.User{Username: req.Username}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Username already exists",
		}, nil
	}

	user.Username = req.Username
	user.Password = utils.Hash(req.Password)

	if result := s.DBPointer.DataBase.Create(&user); result.Error != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error creating user",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if result := s.DBPointer.DataBase.Where(&models.User{Username: req.Username}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid username or password",
		}, nil
	}

	checkedPassword := utils.CheckPasswordHash(req.Password, user.Password)
	if !checkedPassword {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid username or password",
		}, nil
	}

	token, err := s.JWT.GenerateToken(user)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error generating token",
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  "Error validating token",
		}, nil
	}

	var user models.User

	if result := s.DBPointer.DataBase.Where(&models.User{Username: claims.Username}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Id:     user.Id,
	}, nil
}
