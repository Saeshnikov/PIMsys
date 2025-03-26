package branch_service

import (
	"context"
	"fmt"
	proto "pim-sys/gen/go/branch"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	proto.UnimplementedBranchServer // Хитрая штука, о ней ниже
	Branch                          Branch
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Branch interface {
	NewBranch(
		ctx context.Context,
		name string,
		shop_id int32,
		description string,
		address string,
		site string,
		branchtype string,
	) error
	AlterBranch(
		ctx context.Context,
		branchId int32,
		name string,
		description string,
		address string,
		site string,
	) error
	DeleteBranch(
		ctx context.Context,
		branchId int32,
	) error
	ListBranches(
		ctx context.Context,
		shop_id int32,
	) (
		[]*proto.BranchInfo,
		error,
	)
}

func Register(gRPCServer *grpc.Server, branch Branch) {
	proto.RegisterBranchServer(gRPCServer, &ServerAPI{Branch: branch})
}

func (s *ServerAPI) NewBranch(
	ctx context.Context,
	in *proto.NewBranchRequest,
) (*proto.NewBranchResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if in.ShopId == 0 {
		return nil, status.Error(codes.InvalidArgument, "shop id is required")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "description is required")
	}
	if in.Address == "" && in.Site == "" {
		return nil, status.Error(codes.InvalidArgument, "address or site is required")
	}
	err := s.Branch.NewBranch(ctx, in.GetName(), in.GetShopId(), in.GetDescription(), in.GetAddress(), in.GetSite(), in.GetBranchType())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create new branch: %w", err).Error())
	}

	return &proto.NewBranchResponse{}, nil
}

func (s *ServerAPI) AlterBranch(
	ctx context.Context,
	in *proto.AlterBranchRequest,
) (*proto.AlterBranchResponse, error) {
	if in.BranchId == 0 {
		return nil, status.Error(codes.InvalidArgument, "branch id is required")
	}

	err := s.Branch.AlterBranch(ctx, in.GetBranchId(), in.BranchInfo.GetName(), in.BranchInfo.GetDescription(), in.BranchInfo.Address, in.BranchInfo.Site)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to alter branch: %w", err).Error())
	}

	return &proto.AlterBranchResponse{}, nil
}

func (s *ServerAPI) DeleteBranch(
	ctx context.Context,
	in *proto.DeleteBranchRequest,
) (*proto.DeleteBranchResponse, error) {
	if in.BranchId == 0 {
		return nil, status.Error(codes.InvalidArgument, "branch id is required")
	}

	err := s.Branch.DeleteBranch(ctx, in.GetBranchId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete branch: %w", err).Error())
	}

	return &proto.DeleteBranchResponse{}, nil
}

func (s *ServerAPI) ListBranches(
	ctx context.Context,
	in *proto.ListBranchesRequest,
) (*proto.ListBranchesResponse, error) {
	branchInfo, err := s.Branch.ListBranches(ctx, in.ShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to list branches: %w", err).Error())
	}

	return &proto.ListBranchesResponse{Info: branchInfo}, nil
}
