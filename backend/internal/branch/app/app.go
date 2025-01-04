package branch_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/branch"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	branch_service "pim-sys/internal/branch/service"
	"pim-sys/internal/branch/storage"
	grpcapp "pim-sys/internal/grpc"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Branch struct {
	branchStorage *storage.Storage
}

func (branch *Branch) NewBranch(
	ctx context.Context,
	name string,
	shop_id int32,
	description string,
	address string,
	site string,
	branch_type string,
) error {

	err := branch.userMustHaveAccess(ctx, shop_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return branch.branchStorage.CreateBranch(ctx, name, shop_id, description, address, site, branch_type) //sdfsaf
}

func (branch *Branch) AlterBranch(
	ctx context.Context,
	branchId int32,
	name string,
	description string,
	address string,
	site string,
) error {
	shop_id, err := branch.branchStorage.GetShopId(ctx, branchId)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting shop_id ", err)
	}
	err = branch.userMustHaveAccess(ctx, shop_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return branch.branchStorage.AlterBranch(ctx, name, branchId, description, address, site)
}

func (branch *Branch) DeleteBranch(
	ctx context.Context,
	branchId int32,
) error {

	shop_id, err := branch.branchStorage.GetShopId(ctx, branchId)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting shop_id ", err)
	}

	err = branch.userMustHaveAccess(ctx, shop_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return branch.branchStorage.DeleteBranch(ctx, branchId)
}

func (branch *Branch) ListBranches(
	ctx context.Context,
	shop_id int32,
) (
	[]*proto.BranchInfo,
	error,
) {
	return branch.branchStorage.ListBranches(ctx, shop_id)
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	branchStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerBanch := func(gRPCServer *grpc.Server) {
		branch_service.Register(
			gRPCServer,
			&Branch{
				branchStorage: branchStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerBanch, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (branch *Branch) ListShops(
	ctx context.Context,
) (
	[]int32,
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

	return branch.branchStorage.ListShops(ctx, int32(userId))
}

func (branch *Branch) userMustHaveAccess(ctx context.Context, shopId int32) error {
	availableShops, err := branch.ListShops(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user's available shops: ", err)
	}

	for _, available := range availableShops {
		if shopId == available {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or branch does not exist")
}
