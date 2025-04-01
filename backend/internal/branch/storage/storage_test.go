package storage_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/branch"
	"pim-sys/internal/branch/storage"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

//22 tests

func TestStorage_CreateBranch(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		shopID      int32
		description string
		address     string
		site        string
		branch_type string
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
				mockManager.ExpectQuery(".*").WithArgs("test-name", "test-description", "test-address", "test-site", "online", 1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "test-name",
				shopID:      1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
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
				ctx:         context.TODO(),
				name:        "test-name",
				shopID:      1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").WithArgs("test-name", "test-description", "test-address", "test-site", "online", 1).
					WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "test-name",
				shopID:      1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).CreateBranch(tt.args.ctx, tt.args.name, tt.args.shopID, tt.args.description, tt.args.address, tt.args.site, tt.args.branch_type)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_AlterBranch(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		branchId    int32
		description string
		address     string
		site        string
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
				mockManager.ExpectQuery(".*").WithArgs("test-name", "test-description", "test-address", "test-site", 1).WillReturnRows()

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branchId:    1,
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
				ctx:         context.TODO(),
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branchId:    1,
			},
			wantErr: true,
		},
		{
			name: "query error",
			s: func(t *testing.T) *storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").WithArgs("test-name", "test-description", "test-address", "test-site", 1).WillReturnRows().WillReturnError(errors.New("query error"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branchId:    1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.s(t).AlterBranch(tt.args.ctx, tt.args.name, tt.args.branchId, tt.args.description, tt.args.address, tt.args.site)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_DeleteBranch(t *testing.T) {

	type args struct {
		ctx      context.Context
		branchId int32
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
				ctx:      context.TODO(),
				branchId: 1,
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
				ctx:      context.TODO(),
				branchId: 1,
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
				ctx:      context.TODO(),
				branchId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s(t).DeleteBranch(tt.args.ctx, tt.args.branchId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestStorage_ListBranches(t *testing.T) {

	type args struct {
		ctx    context.Context
		shopId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *storage.Storage
		args    args
		want    []*proto.BranchInfo
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
					WillReturnRows(sqlmock.NewRows([]string{"branch.id", "name", "description", "address", "site", "type"}))
				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				shopId: 1,
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
					WillReturnRows(sqlmock.NewRows([]string{"branch.id", "name", "description", "address", "site", "type"}).
						AddRow(1, "test-name", "test-description", "test-address", "test-site", "online"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				shopId: 1,
			},
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
					WillReturnRows(sqlmock.NewRows([]string{"branch.id", "name", "description", "address", "site", "type"}).
						AddRow("invalid", "test-name", "test-description", "test-address", "test-site", "online"))

				return &storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx:    context.TODO(),
				shopId: 1,
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
				ctx:    context.TODO(),
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
				ctx:    context.TODO(),
				shopId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s(t).ListBranches(tt.args.ctx, tt.args.shopId)
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
		want    []int32
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
				ctx:    context.TODO(),
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
					WillReturnRows(sqlmock.NewRows([]string{"shop.id"}).
						AddRow(1))

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
				ctx:    context.TODO(),
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
				mockManager.ExpectQuery(".*").WithArgs(1).WillReturnRows().WillReturnError(errors.New("query error"))

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
			_, err := tt.s(t).ListShops(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

		})
	}
}

func TestStorage_GetShopId(t *testing.T) {
	type args struct {
		ctx      context.Context
		branchId int32
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
				mockManager.ExpectQuery(".*").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"shop_id"}).
						AddRow(1))
				return &storage.Storage{
					DB: db,
				}
			},

			args: args{
				ctx:      context.TODO(),
				branchId: 1,
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
				ctx:      context.TODO(),
				branchId: 1,
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
				ctx:      context.TODO(),
				branchId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s(t).GetShopId(tt.args.ctx, tt.args.branchId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
