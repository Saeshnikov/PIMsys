package shop_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/template"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	template_service "pim-sys/internal/template/service"
	"pim-sys/internal/template/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Template struct {
	templateStorage *storage.Storage
}

func (template *Template) NewTemplate(
	ctx context.Context,
	name string,
	description string,
	branch_id int32,
	attributes []*proto.AttributeInfo,
) error {
	return template.templateStorage.CreateTemplate(ctx, branch_id, name, description, attributes)
}

func (template *Template) AlterTemplate(
	ctx context.Context,
	template_id int32,
	name string,
	description string,
	attributes []*proto.AttributeInfo,
) error {
	err := template.userMustHaveAccess(ctx, template_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.templateStorage.AlterTemplate(ctx, template_id, name, description, attributes)
}

func (template *Template) DeleteTemplate(
	ctx context.Context,
	templateId int32,
) error {

	err := template.userMustHaveAccess(ctx, templateId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.templateStorage.DeleteTemplate(ctx, templateId)
}

func (shop *Template) ListTemplates(
	ctx context.Context,
) (
	[]*proto.TemplateInfo,
	error,
) {
	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	return shop.templateStorage.ListTemplates(ctx, int32(userId))
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	templateStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerTemplate := func(gRPCServer *grpc.Server) {
		template_service.Register(
			gRPCServer,
			&Template{
				templateStorage: templateStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerTemplate, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (template *Template) userMustHaveAccess(ctx context.Context, template_id int32) error {
	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return fmt.Errorf("%s", "template operations: can't take current user id")
	}

	var user_idInt int32
	result, err := strconv.Atoi(user_id)
	user_idInt = int32(result)
	if err != nil {
		return fmt.Errorf("%s", "template operations: can't take cast user_id to int32")
	}
	availibleCategories, err := template.templateStorage.GetUserListCategories(ctx, user_idInt)
	if err != nil {
		return fmt.Errorf("%s%w", "template operations: can't take user's availible categories: ", err)
	}

	for _, category_id := range availibleCategories {
		if category_id == template_id {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or template does not exist")
}
