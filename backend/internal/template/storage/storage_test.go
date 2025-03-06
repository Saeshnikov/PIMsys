package template_storage_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/template"
	template_storage "pim-sys/internal/template/storage"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name DB --dir . --outpkg template_storage_test --output .

// 18 tests
func TestStorage_CreateTemplate(t *testing.T) {
	type args struct {
		ctx         context.Context
		branch_id   int32
		name        string
		description string
		attributes  []*proto.AttributeInfo
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *template_storage.Storage
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("a", true, true, "abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id:   1,
				name:        "abcd",
				description: "abcd",
				attributes: []*proto.AttributeInfo{{
					Type:            "a",
					IsValueRequired: true,
					IsUnique:        true,
					Name:            "abcd",
					Description:     "abcd",
				}},
			},
			wantErr: false,
		},
		{
			name: "CreateTemplate error on first query",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("1", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("a", true, true, "abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id:   1,
				name:        "abcd",
				description: "abcd",
				attributes: []*proto.AttributeInfo{{
					Type:            "a",
					IsValueRequired: true,
					IsUnique:        true,
					Name:            "abcd",
					Description:     "abcd",
				}},
			},
			wantErr: true,
		},
		{
			name: "CreateTemplate error on first prepare",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").
					WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id:   1,
				name:        "abcd",
				description: "abcd",
				attributes: []*proto.AttributeInfo{{
					Type:            "a",
					IsValueRequired: true,
					IsUnique:        true,
					Name:            "abcd",
					Description:     "abcd",
				}},
			},
			wantErr: true,
		},
		{
			name: "CreateTemplate error on second prepare",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*").
					WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id:   1,
				name:        "abcd",
				description: "abcd",
				attributes: []*proto.AttributeInfo{{
					Type:            "a",
					IsValueRequired: true,
					IsUnique:        true,
					Name:            "abcd",
					Description:     "abcd",
				}},
			},
			wantErr: true,
		},
		{
			name: "CreateTemplate error on second query",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("1", true, true, "abcd", "abcd", 1).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id:   1,
				name:        "abcd",
				description: "abcd",
				attributes: []*proto.AttributeInfo{{
					Type:            "a",
					IsValueRequired: true,
					IsUnique:        true,
					Name:            "abcd",
					Description:     "abcd",
				}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s(t).CreateTemplate(
				tt.args.ctx,
				tt.args.branch_id,
				tt.args.name,
				tt.args.description,
				tt.args.attributes); (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteTemplate(t *testing.T) {
	type args struct {
		ctx        context.Context
		templateId int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *template_storage.Storage
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows()

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				templateId: 1,
			},
			wantErr: false,
		},
		{
			name: "DeleteTemplate error on prepare",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").
					WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				templateId: 1,
			},
			wantErr: true,
		},
		{
			name: "DeleteTemplate error on query",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(0).
					WillReturnRows()

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				templateId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s(t).DeleteTemplate(
				tt.args.ctx,
				tt.args.templateId); (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ListTemplate(t *testing.T) {
	type args struct {
		ctx       context.Context
		branch_id int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *template_storage.Storage
		args    args
		want    []*proto.TemplateInfo
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(3))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow("a", true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: false,
			want: []*proto.TemplateInfo{
				{
					Attributes: []*proto.AttributeInfo{{
						Id:              3,
						Type:            "a",
						IsValueRequired: true,
						IsUnique:        true,
						Name:            "abcd",
						Description:     "abcd",
					}},
					Name:        "abcd",
					Description: "abcd",
					TemplateId:  2}},
		},
		{
			name: "ListTemplates error on prepare1",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query1",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("name").
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(1, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on prepare2",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on prepare3",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query2",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs("name").
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(3))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on prepare4",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(3))

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Error!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query3",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(3))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs("Name").
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow(true, true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query3_2",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(1, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(3))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow(true, true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query3_4",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow(2, "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow("name"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow("a", true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query3_4",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow("name", "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow("a", true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
		{
			name: "ListTemplates error on query3_5",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"category.id", "name", "description"}).
						AddRow("name", "abcd", "abcd"))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectPrepare(".*")

				mockManager.ExpectQuery(".*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"attribute_id"}).
						AddRow(1))

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery("SELECT type.*").
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{
						"type",
						"is_value_required",
						"is_unique",
						"name",
						"description"}).
						AddRow("a", true, true, "abcd", "abcd"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).ListTemplates(
				tt.args.ctx,
				tt.args.branch_id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			for idx, res := range got {
				requireTemplateInfoEqual(t, tt.want[idx], res)
			}
		})
	}
}

func requireTemplateInfoEqual(t *testing.T, expected *proto.TemplateInfo, actual *proto.TemplateInfo) {
	t.Helper()

	requireAttributeInfoEqual(t, expected.Attributes, actual.Attributes)
	require.Equal(t, expected.Description, actual.Description)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.TemplateId, actual.TemplateId)
}

func requireAttributeInfoEqual(t *testing.T, expected []*proto.AttributeInfo, actual []*proto.AttributeInfo) {
	t.Helper()

	for idx, res := range expected {
		require.Equal(t, res.Description, actual[idx].Description)
		require.Equal(t, res.Name, actual[idx].Name)
		require.Equal(t, res.IsUnique, actual[idx].IsUnique)
		require.Equal(t, res.IsValueRequired, actual[idx].IsValueRequired)
		require.Equal(t, res.Type, actual[idx].Type)
	}
}

func TestStorage_GetUserListBranchesTemplate(t *testing.T) {
	type args struct {
		ctx     context.Context
		user_id int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *template_storage.Storage
		want    []int32
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"b.id"}).
						AddRow(1))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				user_id: 1,
			},
			wantErr: false,
			want:    []int32{1},
		},
		{
			name: "GetUserListBranches error on prepare",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Errors!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				user_id: 1,
			},
			wantErr: true,
			want:    []int32{1},
		},
		{
			name: "GetUserListBranches error on query",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("name").
					WillReturnRows(sqlmock.NewRows([]string{"b.id"}).
						AddRow(1))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				user_id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).GetUserListBranches(tt.args.ctx, tt.args.user_id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestStorage_GetBranchIdFromTemplateId(t *testing.T) {
	type args struct {
		ctx         context.Context
		template_id int32
	}
	tests := []struct {
		name    string
		s       func(t *testing.T) *template_storage.Storage
		want    int32
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"branch_id"}).
						AddRow(1))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				template_id: 1,
			},
			wantErr: false,
			want:    1,
		},
		{
			name: "GetUserListBranches error on prepare",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*").WillReturnError(errors.New("Errors!"))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				template_id: 1,
			},
			wantErr: true,
			want:    1,
		},
		{
			name: "GetUserListBranches error on query",
			s: func(t *testing.T) *template_storage.Storage {
				db, mockManager, err := sqlmock.New()
				require.NoError(t, err)

				mockManager.ExpectPrepare(".*")
				mockManager.ExpectQuery(".*").
					WithArgs("name").
					WillReturnRows(sqlmock.NewRows([]string{"b.id"}).
						AddRow(1))

				return &template_storage.Storage{
					DB: db,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(),
					metadata.New(map[string]string{"user_id": "1"})),
				template_id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s(t).GetBranchIdFromTemplateId(tt.args.ctx, tt.args.template_id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			require.Equal(t, tt.want, got)
		})
	}
}
