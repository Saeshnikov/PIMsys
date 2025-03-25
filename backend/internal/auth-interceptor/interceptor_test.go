package auth_interceptor

import (
	"context"
	auth_jwt "pim-sys/internal/auth/jwt"
	"pim-sys/internal/auth/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuthInterceptor(t *testing.T) {
	tests := []struct {
		name string
		want grpc.UnaryServerInterceptor
	}{
		{
			name: "no panic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := AuthInterceptor()
			usr := storage.User{
				ID:       1,
				Name:     "test-name",
				Email:    "test-email",
				PassHash: []byte("test-hash"),
				Phone:    "test-phone",
				IsAdmin:  false,
			}

			token, err := auth_jwt.NewToken(usr, time.Hour)
			require.NoError(t, err)

			f(metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"authorization": token})),
				&grpc.UnaryServerInfo{},
				&grpc.UnaryServerInfo{},
				func(context.Context, any) (any, error) { return nil, nil })
		})
	}
}

func TestGetFromContext(t *testing.T) {
	type args struct {
		ctx        context.Context
		entityName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				ctx:        metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				entityName: "user_id",
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFromContext(tt.args.ctx, tt.args.entityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
