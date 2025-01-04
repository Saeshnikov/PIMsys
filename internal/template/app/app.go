package shop_app

import (
	"context"
	"fmt"
	"html/template"
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
	attribute_id []int32,
) error {

	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	err := template.userMustHaveAccess(ctx, user_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "user don't have permissions", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	return template.templateStorage.CreateTemplate(ctx, name, description, attribute_id)
}

func (template *Template) AlterTemplate(
	ctx context.Context,
	shopId int32,
	name string,
	description string,
	url string,
) error {
	err := template.userMustHaveAccess(ctx, shopId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.templateStorage.AlterTemplate(ctx, shopId, name, description, url)
}

func (template *Template) DeleteTemplate(
	ctx context.Context,
	templateId int32,
) error {

	err := template.userMustHaveAccess(ctx, shopId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.templateStorage.DeleteShop(ctx, shopId)
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

	templateStorage, err := template.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerShop := func(gRPCServer *grpc.Server) {
		template_service.Register(
			gRPCServer,
			&Template{
				templateStorage: templateStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerShop, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (template *Template) userMustHaveAccess(ctx context.Context, template_id int32) error {
	availableTemplates, err := template.ListTemplates(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user's available shops: ", err)
	}

	for _, available := range availableTemplates {
		if template_id == available.GetTemplateId() {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or template does not exist")
}
