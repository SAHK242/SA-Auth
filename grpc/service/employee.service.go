package grpcservice

import (
	"auth/config"
	"auth/ent"
	mapper "auth/grpc/mapper"
	grpcutil "auth/grpc/util"
	validator "auth/grpc/validator"
	"auth/proto/auth"
	"auth/proto/gcommon"
	"auth/repository"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strings"
)

type employeeGrpcService struct {
	config             config.Config
	client             *ent.Client
	logger             *zap.SugaredLogger
	redis              *redis.Client
	authGrpcMapper     mapper.AuthGrpcMapper
	authGrpcValidator  validator.AuthValidator
	employeeRepository repository.EmployeeRepository
}

func (s *employeeGrpcService) CreateEmployee(ctx context.Context, request *auth.CreateEmployeeRequest) (*gcommon.EmptyResponse, error) {
	err := s.authGrpcValidator.ValidateCreateEmployeeRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to validate CreateEmployee: %v", err)
	}

	err = s.createEmployee(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee: %v", err)
	}
	err = grpcutil.SendEmailWhenCreateNewEmployee(request.Email, request.LastName+" "+request.FirstName, strings.Split(request.Email, "@")[0])

	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}
	return &gcommon.EmptyResponse{}, nil
}

func (s *employeeGrpcService) createEmployee(ctx context.Context, request *auth.CreateEmployeeRequest) error {
	err := withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.employeeRepository.CreateEmployee(ctx, tx, request)
	})

	if err != nil {
		return fmt.Errorf("error while creating employee: %v", err)
	}
	return nil
}

func (s *employeeGrpcService) ListEmployee(ctx context.Context, request *auth.ListEmployeeRequest) (*auth.ListEmployeeResponse, error) {
	employees, total, err := s.employeeRepository.FindAllEmployee(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding employees: %v", err)
	}

	return &auth.ListEmployeeResponse{
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Employees:    s.authGrpcMapper.ConvertUserSlice(employees),
	}, nil

}

func NewEmployeeGrpcService(
	redis *redis.Client,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	authGrpcMapper mapper.AuthGrpcMapper,
	authGrpcValidator validator.AuthValidator,
	employeeRepository repository.EmployeeRepository,
) EmployeeGrpcService {
	return &employeeGrpcService{
		redis:              redis,
		client:             client,
		logger:             logger,
		config:             config,
		authGrpcMapper:     authGrpcMapper,
		authGrpcValidator:  authGrpcValidator,
		employeeRepository: employeeRepository,
	}
}
