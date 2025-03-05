package auth_app_test

import (
	"context"
	"errors"
	"log/slog"
	auth_app "pim-sys/internal/auth/app"
	auth_errors "pim-sys/internal/auth/errors"
	"pim-sys/internal/auth/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// 9 тестов
//
//go:generate mockery --testonly --name UserSaver --dir . --outpkg auth_app_test --output .
//go:generate mockery --testonly --name UserProvider --dir . --outpkg auth_app_test --output .
func TestAuth_Login(t *testing.T) {
	type fields struct {
		log         *slog.Logger
		usrSaver    func() auth_app.UserSaver
		usrProvider func() auth_app.UserProvider
		tokenTTL    time.Duration
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"User",
						mock.Anything,
						"aboba123@yandex.ru",
					).Return(storage.User{
						PassHash: []byte("$2a$10$8XseJz8XOsi2JPftDhEsoeVBcIN9vDi7OHvz9lXw.xvZ4huDp24ei"),
					}, nil)
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "aboba123@yandex.ru",
				password: "adminadmin",
			},
			wantErr: false,
		},
		{
			name: "User not found",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"User",
						mock.Anything,
						"aboba123@yandex.ru",
					).Return(storage.User{
						PassHash: nil,
					}, auth_errors.ErrUserNotFound)
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "aboba123@yandex.ru",
				password: "adminadmin",
			},
			wantErr: true,
		},
		{
			name: "RaNdOm Db eRroRrrrr",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"User",
						mock.Anything,
						"aboba123@yandex.ru",
					).Return(storage.User{
						PassHash: nil,
					}, errors.New("random db error"))
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "aboba123@yandex.ru",
				password: "adminadmin",
			},
			wantErr: true,
		},
		{
			name: "Wrong password",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"User",
						mock.Anything,
						"aboba123@yandex.ru",
					).Return(storage.User{
						PassHash: nil,
					}, nil)
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "aboba123@yandex.ru",
				password: "adminadmin",
			},
			wantErr: true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auth_app.Auth{
				Log:         tt.fields.log,
				UsrSaver:    tt.fields.usrSaver(),
				UsrProvider: tt.fields.usrProvider(),
				TokenTTL:    tt.fields.tokenTTL,
			}
			_, err := a.Login(tt.args.ctx, tt.args.email, tt.args.password)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestAuth_RegisterNewUser(t *testing.T) {
	type fields struct {
		log         *slog.Logger
		usrSaver    func() auth_app.UserSaver
		usrProvider func() auth_app.UserProvider
		tokenTTL    time.Duration
	}
	type args struct {
		ctx   context.Context
		email string
		pass  string
		name  string
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					uS := &UserSaver{}
					uS.Mock.On(
						"SaveUser",
						mock.Anything,
						"admin",
						mock.Anything,
						"test-name",
						"+79999999999",
					).Return(int64(1), nil)
					return uS
				},
				usrProvider: func() auth_app.UserProvider {
					return &UserProvider{}
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:   context.TODO(),
				email: "admin",
				pass:  "123123",
				name:  "test-name",
				phone: "+79999999999",
			},
			wantErr: false,
		},
		{
			name: "failed generate passhash",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					uS := &UserSaver{}
					uS.Mock.On(
						"SaveUser",
						mock.Anything,
						"admin",
						mock.Anything,
						"test-name",
						"+79999999999",
					).Return(int64(1), nil)
					return uS
				},
				usrProvider: func() auth_app.UserProvider {
					return &UserProvider{}
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:   context.TODO(),
				email: "admin",
				pass:  "123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123123",
				name:  "test-name",
				phone: "+79999999999",
			},
			wantErr: true,
		},
		{
			name: "failed to save user",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					uS := &UserSaver{}
					uS.Mock.On(
						"SaveUser",
						mock.Anything,
						"admin",
						mock.Anything,
						"test-name",
						"+79999999999",
					).Return(int64(0), errors.New("random save user error"))
					return uS
				},
				usrProvider: func() auth_app.UserProvider {
					return &UserProvider{}
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:   context.TODO(),
				email: "admin",
				pass:  "123123",
				name:  "test-name",
				phone: "+79999999999",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auth_app.Auth{
				Log:         tt.fields.log,
				UsrSaver:    tt.fields.usrSaver(),
				UsrProvider: tt.fields.usrProvider(),
				TokenTTL:    tt.fields.tokenTTL,
			}
			_, err := a.RegisterNewUser(tt.args.ctx, tt.args.email, tt.args.pass, tt.args.name, tt.args.phone)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestAuth_IsAdmin(t *testing.T) {
	type fields struct {
		log         *slog.Logger
		usrSaver    func() auth_app.UserSaver
		usrProvider func() auth_app.UserProvider
		tokenTTL    time.Duration
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"IsAdmin",
						mock.Anything,
						int64(1),
					).Return(true, nil)
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:    context.TODO(),
				userID: 1,
			},
			wantErr: false,
		},
		{
			name: "isAdmin error",
			fields: fields{
				log: slog.Default(),
				usrSaver: func() auth_app.UserSaver {
					return &UserSaver{}
				},
				usrProvider: func() auth_app.UserProvider {
					uP := &UserProvider{}
					uP.Mock.On(
						"IsAdmin",
						mock.Anything,
						int64(1),
					).Return(false, errors.New("random isAdmin error"))
					return uP
				},
				tokenTTL: 1,
			},
			args: args{
				ctx:    context.TODO(),
				userID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auth_app.Auth{
				Log:         tt.fields.log,
				UsrSaver:    tt.fields.usrSaver(),
				UsrProvider: tt.fields.usrProvider(),
				TokenTTL:    tt.fields.tokenTTL,
			}
			_, err := a.IsAdmin(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
