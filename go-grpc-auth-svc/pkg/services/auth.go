package services

import (
	"context"
	"net/http"

	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/db"
	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/models"
	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/pb"
	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/utils"
	"gorm.io/gorm"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer.
// func (*Server) mustEmbedUnimplementedAuthServiceServer() {
// 	panic("unimplemented")
// }

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer.
// func (*Server) mustEmbedUnimplementedAuthServiceServer() {
// 	panic("unimplemented")
// }

// // mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer.
// func (*Server) mustEmbedUnimplementedAuthServiceServer() {
// 	panic("unimplemented")
// }

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil

	}
	if result.Error == nil && user.Id != 0 {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Record Already Exist",
		}, nil
	}
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
