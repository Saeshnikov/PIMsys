package auth_interceptor

import (
	"context"
	"fmt"
	"log"
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
		const appSecret = "test-secret"

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
			return []byte(appSecret), nil
		})

		if err != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "role", strconv.Itoa(NotAuthorized))
			return handler(ctx, req)
		}

		uid, ok := claims["uid"].(float64)
		if !ok {
			str := fmt.Sprintf("uid type: %T", claims["uid"])
			return nil, status.Error(codes.Internal, "uid type has changed; "+str)
		}

		urole, ok := claims["role"].(bool)
		if !ok {
			str := fmt.Sprintf("role: %T", claims["role"])
			return nil, status.Error(codes.Internal, "role has changed; "+str)
		}
		role := User
		if urole {
			role = Admin
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "role", strconv.Itoa(role))
		ctx = metadata.AppendToOutgoingContext(ctx, "user_id", strconv.Itoa(int(uid)))
		return handler(ctx, req)
	}
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
