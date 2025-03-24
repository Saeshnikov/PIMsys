package auth_service_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/sso"
	auth_errors "pim-sys/internal/auth/errors"
	auth_service "pim-sys/internal/auth/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// 27 тестов
//go:generate mockery --testonly --name Auth --dir . --outpkg auth_service_test --output .

func Test_serverAPI_Login(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer proto.UnimplementedAuthServer
		Auth                    func() *Auth
	}
	type args struct {
		ctx context.Context
		in  *proto.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.LoginResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"Login",
						mock.Anything,
						"aboba123@yandex.ru",
						"123123",
					).Return("token", nil)
					return auth
				},
			},
			args: args{
				in: &proto.LoginRequest{
					Email:    "aboba123@yandex.ru",
					Password: "123123",
				},
			},
			wantErr: false,
		},
		{
			name: "missing email",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.LoginRequest{
					Email:    "",
					Password: "123123",
				},
			},
			wantErr: true,
		},
		{
			name: "missing password",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.LoginRequest{
					Email:    "aboba123@yandex.ru",
					Password: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid email or password",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"Login",
						mock.Anything,
						"aboba123@yandex.ru",
						"123123",
					).Return("", auth_errors.ErrInvalidCredentials)
					return auth
				},
			},
			args: args{
				in: &proto.LoginRequest{
					Email:    "aboba123@yandex.ru",
					Password: "123123",
				},
			},
			wantErr: true,
		},
		{
			name: "failed to login",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"Login",
						mock.Anything,
						"aboba123@yandex.ru",
						"123123",
					).Return("", errors.New("random login error"))
					return auth
				},
			},
			args: args{
				in: &proto.LoginRequest{
					Email:    "aboba123@yandex.ru",
					Password: "123123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &auth_service.ServerAPI{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				Auth:                    tt.fields.Auth(),
			}
			_, err := s.Login(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_Register(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer proto.UnimplementedAuthServer
		Auth                    func() *Auth
	}
	type args struct {
		ctx context.Context
		in  *proto.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.RegisterResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "email is missing",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		//email len should be >1 and <16
		{
			name: "email is not valid(len = 1)",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "a",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "email is valid(len = 2)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"aa",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "aa",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "email is valid(len = 5)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"aaaaa",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "aaaaa",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "email is valid(len = 15)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"aaaaaaaaaaaaaaa",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "aaaaaaaaaaaaaaa",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "email is not valid(len = 16)",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "aaaaaaaaaaaaaaaa",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "pass is missing",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		//pass len should be >=8
		{
			name: "pass is not valid(len = 7)",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "pass is valid(len = 8)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "pass is valid(len = 9)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aaaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "pass is not valid(contains invalid symbol)",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa'",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "pass is valid(with all symbols)",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aA!_?$#@",
						"testname",
						"79999999999",
					).Return(int64(1), nil)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aA!_?$#@",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: false,
		},
		{
			name: "name is missing",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "phone is missing",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "",
				},
			},
			wantErr: true,
		},
		{
			name: "phone is not valid",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "a79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "user already exists",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(0), auth_errors.ErrUserExists)
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
		{
			name: "failed to register user",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"RegisterNewUser",
						mock.Anything,
						"admin",
						"aaaaaaaa",
						"testname",
						"79999999999",
					).Return(int64(0), errors.New("random register error"))
					return auth
				},
			},
			args: args{
				in: &proto.RegisterRequest{
					Email:    "admin",
					Password: "aaaaaaaa",
					Name:     "testname",
					Phone:    "79999999999",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &auth_service.ServerAPI{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				Auth:                    tt.fields.Auth(),
			}
			_, err := s.Register(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_IsAdmin(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer proto.UnimplementedAuthServer
		Auth                    func() *Auth
	}
	type args struct {
		ctx context.Context
		in  *proto.IsAdminRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.IsAdminResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"IsAdmin",
						mock.Anything,
						int64(1),
					).Return(true, nil)
					return auth
				},
			},
			args: args{
				in: &proto.IsAdminRequest{
					UserId: int64(1),
				},
			},
			wantErr: false,
		},
		{
			name: "uid is required",
			fields: fields{
				Auth: func() *Auth {
					return &Auth{}
				},
			},
			args: args{
				in: &proto.IsAdminRequest{
					UserId: int64(0),
				},
			},
			wantErr: true,
		},
		{
			name: "user not found",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"IsAdmin",
						mock.Anything,
						int64(1),
					).Return(false, auth_errors.ErrUserNotFound)
					return auth
				},
			},
			args: args{
				in: &proto.IsAdminRequest{
					UserId: int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "failed to check admin status",
			fields: fields{
				Auth: func() *Auth {
					auth := &Auth{}
					auth.Mock.On(
						"IsAdmin",
						mock.Anything,
						int64(1),
					).Return(false, errors.New("random isAdmin error"))
					return auth
				},
			},
			args: args{
				in: &proto.IsAdminRequest{
					UserId: int64(1),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &auth_service.ServerAPI{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				Auth:                    tt.fields.Auth(),
			}
			_, err := s.IsAdmin(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
