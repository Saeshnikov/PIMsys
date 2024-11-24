package auth_interceptor

import (
	"context"
	"fmt"
	"log"
	"pim-sys/internal/auth/secret"
	"strconv"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	NotAuthorized = iota
	User
	Admin
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> Validate jwt-token: ", info.FullMethod)

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "role", strconv.Itoa(NotAuthorized))
			return handler(ctx, req)
		}

		values := md["authorization"]
		if len(values) == 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, "role", strconv.Itoa(NotAuthorized))
			return handler(ctx, req)
		}

		accessToken := values[0]
		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret.Secret), nil
		})

		if err != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "role", strconv.Itoa(NotAuthorized))
			return handler(ctx, req)
		}

		uid, ok := claims["uid"].(float64)
		if !ok {
			return nil, status.Error(codes.Internal, "uid not found")
		}
		md.Append("user_id", strconv.Itoa(int(uid)))

		urole, ok := claims["role"].(bool)
		if !ok {
			return nil, status.Error(codes.Internal, "role not found")
		}
		role := User
		if urole {
			role = Admin
		}
		md.Append("role", strconv.Itoa(role))

		newCtx := metadata.NewIncomingContext(ctx, md)

		return handler(newCtx, req)
	}
}

func GetFromContext(ctx context.Context, entityName string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("%s", "metadata does not exist")
	}

	values := md[entityName]
	if len(values) == 0 {
		return "", fmt.Errorf("%s", "uid does not exist")
	}

	return values[0], nil
}

// func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
// 	return func(
// 		srv interface{},
// 		stream grpc.ServerStream,
// 		info *grpc.StreamServerInfo,
// 		handler grpc.StreamHandler,
// 	) error {
// 		log.Println("--> stream interceptor: ", info.FullMethod)

// 		// TODO: implement authorization

// 		return handler(srv, stream)
// 	}
// }
