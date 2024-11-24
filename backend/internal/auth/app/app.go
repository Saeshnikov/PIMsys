package auth_app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	auth_interceptor "pim-sys/internal/auth-interceptor"
	auth_errors "pim-sys/internal/auth/errors"
	auth_jwt "pim-sys/internal/auth/jwt"
	auth_service "pim-sys/internal/auth/service"
	"pim-sys/internal/auth/storage"
	grpcapp "pim-sys/internal/grpc"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
		name string,
		phone string,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (storage.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

// Login checks if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, auth_errors.ErrUserNotFound) {
			a.log.Warn("user not found")

			return "", fmt.Errorf("%s: %w", op, auth_errors.ErrInvalidCredentials)
		}

		a.log.Error("failed to get user")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials")

		return "", fmt.Errorf("%s: %w", op, auth_errors.ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	token, err := auth_jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string, name string, phone string) (int64, error) {
	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
		slog.String("name", name),
		slog.String("phone", phone),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash")

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, email, passHash, name, phone)
	if err != nil {
		log.Error("failed to save user")

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// IsAdmin checks if user is admin.
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "Auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("user_id", userID),
	)

	log.Info("checking if user is admin")

	isAdmin, err := a.usrProvider.IsAdmin(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))

	return isAdmin, nil
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {
	storage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerAuth := func(gRPCServer *grpc.Server) {
		auth_service.Register(
			gRPCServer,
			&Auth{
				log:         log,
				usrSaver:    storage,
				usrProvider: storage,
				tokenTTL:    tokenTTL,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerAuth, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}
