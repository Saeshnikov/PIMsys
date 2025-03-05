package storage_test

import (
	"context"
	"database/sql"
	"errors"
	"pim-sys/internal/auth/storage"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

// 12 тестов
func TestStorage_SaveUser(t *testing.T) {
	type fields struct {
		DB func(t *testing.T) storage.DB
	}
	type args struct {
		ctx      context.Context
		email    string
		passHash []byte
		name     string
		phone    string
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
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin", []byte("121341"), "testname", "79999999999", false).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1))
					return db
				},
			},

			args: args{
				ctx:      context.TODO(),
				email:    "admin",
				passHash: []byte("121341"),
				name:     "testname",
				phone:    "79999999999",
			},
			wantErr: false,
		},
		{
			name: "db error",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*").WillReturnError(errors.New("some db error"))

					return db
				},
			},

			args: args{
				ctx:      context.TODO(),
				email:    "admin",
				passHash: []byte("121341"),
				name:     "testname",
				phone:    "79999999999",
			},
			wantErr: true,
		},
		{
			name: "unique_violation",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin", []byte("121341"), "testname", "79999999999", false).
						WillReturnError(&pq.Error{Code: pq.ErrorCode("23505")})
					return db
				},
			},

			args: args{
				ctx:      context.TODO(),
				email:    "admin",
				passHash: []byte("121341"),
				name:     "testname",
				phone:    "79999999999",
			},
			wantErr: true,
		},
		{
			name: "some error",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin", []byte("121341"), "testname", "79999999999", false).
						WillReturnError(errors.New("some db error"))
					return db
				},
			},

			args: args{
				ctx:      context.TODO(),
				email:    "admin",
				passHash: []byte("121341"),
				name:     "testname",
				phone:    "79999999999",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage.Storage{
				DB: tt.fields.DB(t),
			}
			_, err := s.SaveUser(tt.args.ctx, tt.args.email, tt.args.passHash, tt.args.name, tt.args.phone)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_User(t *testing.T) {
	type fields struct {
		DB func(t *testing.T) storage.DB
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    storage.User
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin").
						WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "isAdmin"}).
							AddRow(1, "admin", []byte("121341"), false))
					return db
				},
			},

			args: args{
				ctx:   context.TODO(),
				email: "admin",
			},
			wantErr: false,
		},
		{
			name: "error preparing querry",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*").WillReturnError(errors.New("some db error"))

					return db
				},
			},

			args: args{
				ctx:   context.TODO(),
				email: "admin",
			},
			wantErr: true,
		},
		{
			name: "error scaning querry(user not found)",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin").
						WillReturnError(sql.ErrNoRows)
					return db
				},
			},

			args: args{
				ctx:   context.TODO(),
				email: "admin",
			},
			wantErr: true,
		},
		{
			name: "error scaning querry",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs("admin").
						WillReturnError(errors.New("some db error"))
					return db
				},
			},

			args: args{
				ctx:   context.TODO(),
				email: "admin",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage.Storage{
				DB: tt.fields.DB(t),
			}
			_, err := s.User(tt.args.ctx, tt.args.email)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_IsAdmin(t *testing.T) {
	type fields struct {
		DB func(t *testing.T) storage.DB
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
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs(1).
						WillReturnRows(sqlmock.NewRows([]string{"isAdmin"}).
							AddRow(false))
					return db
				},
			},

			args: args{
				ctx:    context.TODO(),
				userID: 1,
			},
			wantErr: false,
		},
		{
			name: "error preparing querry",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*").WillReturnError(errors.New("some db error"))
					return db
				},
			},

			args: args{
				ctx:    context.TODO(),
				userID: 1,
			},
			wantErr: true,
		},
		{
			name: "error scaning querry(user not found)",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs(1).
						WillReturnError(sql.ErrNoRows)
					return db
				},
			},

			args: args{
				ctx:    context.TODO(),
				userID: 1,
			},
			wantErr: true,
		},
		{
			name: "error scaning querry",
			fields: fields{
				DB: func(t *testing.T) storage.DB {
					db, mockManager, err := sqlmock.New()
					require.NoError(t, err)

					mockManager.ExpectPrepare(".*")
					mockManager.ExpectQuery(".*").WithArgs(1).
						WillReturnError(errors.New("some error"))
					return db
				},
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
			s := &storage.Storage{
				DB: tt.fields.DB(t),
			}
			_, err := s.IsAdmin(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
