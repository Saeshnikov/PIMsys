package auth_service

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/sso"
	auth_errors "pim-sys/internal/auth/errors"

	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	LoginOrEmailRegex = `^(?:[a-zA-Z0-9]+|[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`
	PasswordRegex     = `^[a-zA-Z!_?$#@]{8,}$`
	PhoneRegex        = `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
)

type serverAPI struct {
	proto.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
		name string,
		phone string,
	) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	proto.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *proto.LoginRequest,
) (*proto.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword())

	if err != nil {
		if errors.Is(err, auth_errors.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &proto.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *proto.RegisterRequest,
) (*proto.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	rxLoginOrEmail := regexp.MustCompile(LoginOrEmailRegex)

	if !rxLoginOrEmail.MatchString(in.Email) || !(len(in.Email) > 1 && len(in.Email) < 100) {
		return nil, status.Error(codes.InvalidArgument, "email is not valid")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	rxPassword := regexp.MustCompile(PasswordRegex)

	if !rxPassword.MatchString(in.Password) {
		return nil, status.Error(codes.InvalidArgument, "password is not valid")
	}

	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if in.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, "phone is required")
	}
	rxPhone := regexp.MustCompile(PhoneRegex)
	if !rxPhone.MatchString(in.Phone) {
		return nil, status.Error(codes.InvalidArgument, "phone is not valid")
	}

	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword(), in.GetName(), in.GetPhone())

	if err != nil {
		if errors.Is(err, auth_errors.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &proto.RegisterResponse{UserId: uid}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	in *proto.IsAdminRequest,
) (*proto.IsAdminResponse, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, auth_errors.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "failed to check admin status")
	}

	return &proto.IsAdminResponse{IsAdmin: isAdmin}, nil
}
