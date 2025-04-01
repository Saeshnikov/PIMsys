package storage_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/logs"
	"pim-sys/internal/logs/storage"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestStorage_GetLogs(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []*proto.Log
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"log_journal.shop_id",
						"log_journal.branch_id",
						"log_journal.product_id",
						"log_journal.info",
					}).
						AddRow(1, 1, 1, "info"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("err"))
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"log_journal.shop_id",
						"log_journal.branch_id",
						"log_journal.product_id",
						"log_journal.info",
					}).
						AddRow(1, 1, 1, "info"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).WillReturnError(errors.New("err"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			wantErr: true,
		},
		{
			name: "scan err",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"log_journal.shop_id",
						"log_journal.branch_id",
						"log_journal.product_id",
						"log_journal.info",
					}).
						AddRow(1, 1, "invalid", "info"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s(t).GetLogs(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_GetSales(t *testing.T) {
	type args struct {
		ctx      context.Context
		TimeFrom int64
		TimeTo   int64
		userId   int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []*proto.Graph
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1, 1, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"sales.date",
						"sales.price",
						"sales.quantity",
					}).
						AddRow(1, 1, 1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				userId:   1,
				TimeFrom: 1,
				TimeTo:   1,
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("err"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				userId:   1,
				TimeFrom: 1,
				TimeTo:   1,
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1, 1, 1).WillReturnError(errors.New("error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				userId:   1,
				TimeFrom: 1,
				TimeTo:   1,
			},
			wantErr: true,
		},
		{
			name: "scan error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1, 1, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"sales.date",
						"sales.price",
						"sales.quantity",
					}).
						AddRow(1, 1, "invaslid"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				userId:   1,
				TimeFrom: 1,
				TimeTo:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s(t).GetSales(tt.args.ctx, tt.args.TimeFrom, tt.args.TimeTo, tt.args.userId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_GetMinDate(t *testing.T) {
	type args struct {
		ctx      context.Context
		dateFrom int64
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{
						"date",
					}).
						AddRow(1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				dateFrom: 1,
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				dateFrom: 1,
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs().WillReturnError(errors.New("error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				dateFrom: 1,
			},
			wantErr: true,
		},
		{
			name: "scan error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{
						"date",
					}).
						AddRow("asdasd"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:      context.TODO(),
				dateFrom: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).GetMinDate(tt.args.ctx, tt.args.dateFrom)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
