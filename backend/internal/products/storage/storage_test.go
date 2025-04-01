package storage_test

import (
	"context"
	proto "pim-sys/gen/go/products"
	"pim-sys/internal/products/storage"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestStorage_CreateProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.ProductInfo
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1, "status", 1, "name", 1.0, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: context.TODO(),
				content: &proto.ProductInfo{
					CategoryId: 1,
					BranchId:   1,
					Status:     "status",
					Name:       "name",
					Price:      1.0,
					Amount:     1,
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).CreateProduct(tt.args.ctx, tt.args.content)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStorage_AlterAttributes(t *testing.T) {
	type args struct {
		ctx       context.Context
		productId int32
		attr      *proto.Attribute
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
					WithArgs(1, 1, "text", 1, true).
					WillReturnRows()

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1, 1, "text", 1, true).
					WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:       context.TODO(),
				productId: 1,
				attr: &proto.Attribute{
					Id:          1,
					ValueText:   "text",
					ValueNumber: 1,
					ValueBool:   true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).AlterAttributes(tt.args.ctx, tt.args.productId, tt.args.attr)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_AlterProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.ProductInfoWithId
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
					WithArgs("status", 1, "name", 1, 1.0, 1).
					WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: context.TODO(),
				content: &proto.ProductInfoWithId{
					ProductId: 1,
					Product: &proto.ProductInfo{
						CategoryId: 1,
						BranchId:   1,
						Status:     "status",
						Name:       "name",
						Price:      1.0,
						Amount:     1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).AlterProduct(tt.args.ctx, tt.args.content)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_DeleteProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.DeleteProductRequest
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
					WithArgs(1).
					WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: context.TODO(),
				content: &proto.DeleteProductRequest{
					ProductId: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).DeleteProduct(tt.args.ctx, tt.args.content)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_ListProducts(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []*proto.ProductInfoWithId
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
					WillReturnRows(sqlmock.NewRows([]string{"product.id", "category_id", "status", "branch_id", "name", "price", "amount"}).
						AddRow(1, 1, "status", 1, "name", 1.0, 1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id", "value_text", "value_number", "value_boolean"}).
						AddRow(1, "text", 1, true))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			want: []*proto.ProductInfoWithId{
				{
					ProductId: 1,
					Product: &proto.ProductInfo{
						CategoryId: 1,
						BranchId:   1,
						Status:     "status",
						Name:       "name",
						Price:      1.0,
						Amount:     1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).ListProducts(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			for idx, res := range got {
				requireProductInfoWithIdEqual(t, tt.want[idx], res)
			}
		})
	}
}

func requireProductInfoWithIdEqual(t *testing.T, expected *proto.ProductInfoWithId, actual *proto.ProductInfoWithId) {
	t.Helper()

	require.Equal(t, expected.ProductId, actual.ProductId)

	requireProductInfoEqual(t, expected.GetProduct(), actual.GetProduct())
}

func requireProductInfoEqual(t *testing.T, expected *proto.ProductInfo, actual *proto.ProductInfo) {
	t.Helper()

	require.Equal(t, expected.Amount, actual.Amount)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.BranchId, actual.BranchId)
	require.Equal(t, expected.CategoryId, actual.CategoryId)
	require.Equal(t, expected.Price, actual.Price)
	require.Equal(t, expected.Status, actual.Status)
}

func TestStorage_SellProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.SellProductRequest
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
					WithArgs(1, 1).
					WillReturnRows()

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(sqlmock.AnyArg(), 1, 1).
					WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: context.TODO(),
				content: &proto.SellProductRequest{
					ProductId: 1,
					Amount:    1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).SellProduct(tt.args.ctx, tt.args.content)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_GetAccessableBranchIds(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []int32
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
					WillReturnRows(sqlmock.NewRows([]string{"branch.id"}).
						AddRow(1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				userId: 1,
			},
			want:    []int32{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).GetAccessableBranchIds(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
