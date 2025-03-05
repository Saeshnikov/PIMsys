package storage_test

import (
	"context"
	"errors"
	"fmt"
	proto "pim-sys/gen/go/shop"
	"pim-sys/internal/shop/storage"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

// 14 tests

func TestStorage_CreateShop(t *testing.T) {
	type args struct {
		ctx         context.Context
		userID      int
		name        string
		description string
		url         string
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
				mockManager.ExpectQuery(".*").WithArgs("name", "url", "description", 1).WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userID:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("prepare error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userID:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").WithArgs("name", "url", "description", 1).WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userID:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).CreateShop(tt.args.ctx, tt.args.userID, tt.args.name, tt.args.description, tt.args.url)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_AlterShop(t *testing.T) {
	type args struct {
		ctx         context.Context
		shopId      int32
		name        string
		description string
		url         string
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
				mockManager.ExpectQuery(".*").WithArgs("name", "url", "description", 1).WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("prepare error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").WithArgs("name", "url", "description", 1).WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId:      1,
				name:        "name",
				description: "description",
				url:         "url",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).AlterShop(tt.args.ctx, tt.args.shopId, tt.args.name, tt.args.description, tt.args.url)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

		})
	}
}

func TestStorage_DeleteShop(t *testing.T) {
	type args struct {
		ctx    context.Context
		shopId int32
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
				mockManager.ExpectQuery(".*").WithArgs(1).WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId: 1,
			},
			wantErr: false,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("prepare error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId: 1,
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").WithArgs(1).WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shopId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).DeleteShop(tt.args.ctx, tt.args.shopId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_ListShops(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []*proto.ShopInfo
		wantErr bool
	}{
		{
			name: "default empty",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"shop.id", "name", "description", "avatar_url"}))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userId: 1,
			},
			wantErr: false,
		},
		{
			name: "default with row",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"shop.id", "name", "description", "avatar_url"}).
						AddRow(1, "name", "description", "url"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userId: 1,
			},
			want:    []*proto.ShopInfo{{ShopId: 1, Name: "name", Description: "description", Url: "url"}},
			wantErr: false,
		},
		{
			name: "invalid rows",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"shop.id", "name", "description", "avatar_url"}).
						AddRow("invalid", 1, 1, 1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userId: 1,
			},
			wantErr: true,
		},
		{
			name: "prepare error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("prepare error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
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
				mockManager.ExpectQuery(".*").WithArgs(1).WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				userId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).ListShops(tt.args.ctx, tt.args.userId)
			fmt.Println(got)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			for idx, res := range got {
				requireShopInfoEqual(t, tt.want[idx], res)
			}

		})
	}
}

func requireShopInfoEqual(t *testing.T, expected *proto.ShopInfo, actual *proto.ShopInfo) {
	t.Helper()

	require.Equal(t, expected.ShopId, actual.ShopId)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.Description, actual.Description)
	require.Equal(t, expected.Url, actual.Url)
}
