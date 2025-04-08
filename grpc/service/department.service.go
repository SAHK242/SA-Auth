package grpcservice

import (
	"auth/config"
	"auth/ent"
	grpcmapper "auth/grpc/mapper"
	grpcutil "auth/grpc/util"
	validator "auth/grpc/validator"
	"auth/proto/auth"
	"auth/proto/gcommon"
	"auth/repository"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type departmentGrpcService struct {
	config                  config.Config
	client                  *ent.Client
	logger                  *zap.SugaredLogger
	redis                   *redis.Client
	departmentRepository    repository.DepartmentRepository
	departmentGrpcMapper    grpcmapper.DepartmentGrpcMapper
	departmentGrpcValidator validator.DepartmentValidator
}

func (s *departmentGrpcService) CreateDepartment(ctx context.Context, request *auth.CreateDepartmentRequest) (*gcommon.EmptyResponse, error) {
	err := s.departmentGrpcValidator.ValidateCreateDepartmentRequest(ctx, request)

	if err != nil {
		return nil, fmt.Errorf("error while validating request: %v", err)
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.departmentRepository.CreateDepartment(ctx, tx, request)
	})
	return &gcommon.EmptyResponse{}, err
}

func (s *departmentGrpcService) ListDepartment(ctx context.Context, request *auth.ListDepartmentRequest) (*auth.ListDepartmentResponse, error) {
	departments, total, err := s.departmentRepository.FindAllDepartment(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding departments: %v", err)
	}

	return &auth.ListDepartmentResponse{
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Departments:  s.departmentGrpcMapper.ConvertDepartmentSlice(departments),
	}, nil
}

func NewDepartmentGrpcService(
	redis *redis.Client,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	departmentRepository repository.DepartmentRepository,
	departmentGrpcMapper grpcmapper.DepartmentGrpcMapper,
	departmentGrpcValidator validator.DepartmentValidator,
) DepartmentGrpcService {
	return &departmentGrpcService{
		redis:                   redis,
		client:                  client,
		logger:                  logger,
		config:                  config,
		departmentRepository:    departmentRepository,
		departmentGrpcMapper:    departmentGrpcMapper,
		departmentGrpcValidator: departmentGrpcValidator,
	}
}
