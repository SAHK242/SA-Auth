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
	"github.com/Luzifer/go-openssl/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type authGrpcService struct {
	config            config.Config
	client            *ent.Client
	logger            *zap.SugaredLogger
	redis             *redis.Client
	authRepository    repository.AuthRepository
	authGrpcMapper    mapper.AuthGrpcMapper
	authGrpcValidator validator.AuthValidator
}

func (s *authGrpcService) CreateEmployee(ctx context.Context, request *auth.CreateEmployeeRequest) (*gcommon.EmptyResponse, error) {
	err := s.authGrpcValidator.ValidateCreateEmployeeRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to validate CreateEmployee: %v", err)
	}

	err = s.createEmployee(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee: %v", err)
	}
	err = grpcutil.SendEmail(request.Email, request.LastName+" "+request.FirstName, strings.Split(request.Email, "@")[0])

	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}
	return &gcommon.EmptyResponse{}, nil
}

func (s *authGrpcService) createEmployee(ctx context.Context, request *auth.CreateEmployeeRequest) error {
	err := withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.authRepository.CreateEmployee(ctx, tx, request)
	})

	if err != nil {
		return fmt.Errorf("error while creating employee: %v", err)
	}
	return nil
}

func (s *authGrpcService) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*gcommon.EmptyResponse, error) {
	key, err := s.redis.Get(ctx, request.Username).Result()

	if err != nil {
		return nil, fmt.Errorf("invalid request - key not found: %v", err)
	}

	found := key != ""

	if !found {
		return nil, fmt.Errorf("invalid request")
	}

	dec, err := openSslClient.DecryptBytes(key, []byte(request.OldPassword), openssl.BytesToKeyMD5)

	if err != nil {
		return nil, fmt.Errorf("error while decrypting password: %v", err)
	}

	plainPassword := string(dec)

	user, err := s.authRepository.FindByUsername(ctx, request.Username)

	if err != nil {
		return nil, fmt.Errorf("error while finding user: %v", err)
	}

	accountState := user.State

	newPassword := request.NewPassword

	newPassDec, err := openSslClient.DecryptBytes(key, []byte(newPassword), openssl.BytesToKeyMD5)
	plainNewPassword := string(newPassDec)

	if err != nil {
		return nil, fmt.Errorf("error while decrypting new password: %v", err)
	}

	switch accountState {
	case int32(gcommon.AccountState_ACCOUNT_STATE_ACTIVE):
		break
	case int32(gcommon.AccountState_ACCOUNT_STATE_INACTIVE):
		break
	case int32(gcommon.AccountState_ACCOUNT_STATE_LOCKED):
		if user.Password == plainPassword {
			e := s.updatePassword(ctx, request.Username, plainNewPassword)
			if e != nil {
				return nil, fmt.Errorf("error while updating password: %v", e)
			}
			// Update state to active
			e = s.updateAccountState(ctx, request.Username, int32(gcommon.AccountState_ACCOUNT_STATE_ACTIVE))
			if e != nil {
				return nil, fmt.Errorf("error while updating account state: %v", e)
			}
		}
		break
	case int32(gcommon.AccountState_ACCOUNT_STATE_SUPER_ADMIN_LOCKED):
		if user.Password == plainPassword {
			e := s.updatePassword(ctx, request.Username, plainNewPassword)
			if e != nil {
				return nil, fmt.Errorf("error while updating password: %v", e)
			}
			// Update state to active
			e = s.updateAccountState(ctx, request.Username, int32(gcommon.AccountState_ACCOUNT_STATE_SUPER_ADMIN_ACTIVE))
			if e != nil {
				return nil, fmt.Errorf("error while updating account state: %v", e)
			}

			return &gcommon.EmptyResponse{}, nil
		}
		break
	case int32(gcommon.AccountState_ACCOUNT_STATE_SUPER_ADMIN_ACTIVE):
		break
	default:
		break
	}

	return &gcommon.EmptyResponse{}, status.Error(codes.Unauthenticated, "invalid username or password")
}

var openSslClient = openssl.New()

func (s *authGrpcService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	key, err := s.redis.Get(ctx, request.Username).Result()

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "pre-flight check failed, hash password with pre-flight key")
	}

	found := key != ""

	if !found {
		return nil, fmt.Errorf("invalid request")
	}

	dec, err := openSslClient.DecryptBytes(key, []byte(request.Password), openssl.BytesToKeyMD5)

	if err != nil {
		return nil, fmt.Errorf("error while decrypting password: %v", err)
	}

	plainPassword := string(dec)

	user, err := s.authRepository.FindByUsername(ctx, request.Username)

	if err != nil {
		return nil, fmt.Errorf("error while finding user: %v", err)
	}

	accountState := user.State

	switch accountState {
	case int32(gcommon.AccountState_ACCOUNT_STATE_ACTIVE), int32(gcommon.AccountState_ACCOUNT_STATE_SUPER_ADMIN_ACTIVE):
		empId := ""
		if user.Edges.Employee != nil {
			empId = user.Edges.Employee.ID.String()
		}
		if CheckPasswordHash(plainPassword, user.Password) {
			token, err := s.generateToken(ctx, user.ID.String(), empId)
			if err != nil {
				return nil, fmt.Errorf("error while generating token: %v", err)
			}
			return &auth.LoginResponse{
				Token: token,
				User:  s.authGrpcMapper.ConvertUser(user.Edges.Employee),
			}, nil
		}
		break
	case int32(gcommon.AccountState_ACCOUNT_STATE_SUPER_ADMIN_LOCKED), int32(gcommon.AccountState_ACCOUNT_STATE_LOCKED):
		if user.Password == plainPassword {
			return &auth.LoginResponse{
				NextStep: auth.NextStep_NEXT_STEP_REQUIRE_CHANGE_PASSWORD,
			}, nil
		}
		break
	default:
		return nil, fmt.Errorf("account is inactive")
	}

	return &auth.LoginResponse{}, status.Error(codes.Unauthenticated, "invalid username or password")
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *authGrpcService) updatePassword(ctx context.Context, username string, password string) error {
	encryptedPassword, err := HashPassword(password)

	if err != nil {
		return fmt.Errorf("error while hashing password: %v", err)
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.authRepository.UpdatePassword(ctx, tx, username, encryptedPassword)
	})

	if err != nil {
		return fmt.Errorf("error while updating password: %v", err)
	}
	return nil
}

func (s *authGrpcService) updateAccountState(ctx context.Context, username string, state int32) error {
	err := withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.authRepository.UpdateAccountState(ctx, tx, username, state)
	})

	if err != nil {
		return fmt.Errorf("error while updating account state: %v", err)
	}
	return nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT creates a JWT token
func (s *authGrpcService) GenerateJWT(userID string, employeeId string) (string, error) {
	// Secret key
	jwtSecret := []byte(s.config.GetString("JWT_SECRET", "secret"))
	// Create claims (payload)
	claims := jwt.MapClaims{
		"auth_id": userID,
		"emp_id":  employeeId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expiry: 24 hours
		"iat":     time.Now().Unix(),                     // Issued At
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authGrpcService) generateToken(ctx context.Context, authId string, empId string) (string, error) {
	token, err := s.GenerateJWT(authId, empId)
	if err != nil {
		return "", fmt.Errorf("error while generating token: %v", err)
	}
	return token, nil
}

func (s *authGrpcService) CreateDepartment(ctx context.Context, request *auth.CreateDepartmentRequest) (*gcommon.EmptyResponse, error) {
	if request.Name == "" {
		return nil, fmt.Errorf("department name is required")
	}

	err := withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.authRepository.CreateDepartment(ctx, tx, request)
	})
	return &gcommon.EmptyResponse{}, err
}

func (s *authGrpcService) ListEmployee(ctx context.Context, request *auth.ListEmployeeRequest) (*auth.ListEmployeeResponse, error) {
	employees, total, err := s.authRepository.FindAllEmployee(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding employees: %v", err)
	}

	return &auth.ListEmployeeResponse{
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Employees:    s.authGrpcMapper.ConvertUserSlice(employees),
	}, nil

}

func (s *authGrpcService) ListDepartment(ctx context.Context, request *auth.ListDepartmentRequest) (*auth.ListDepartmentResponse, error) {
	departments, total, err := s.authRepository.FindAllDepartment(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding departments: %v", err)
	}

	return &auth.ListDepartmentResponse{
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Departments:  s.authGrpcMapper.ConvertDepartmentSlice(departments),
	}, nil
}

func NewAuthGrpcService(
	redis *redis.Client,
	authRepository repository.AuthRepository,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	authGrpcMapper mapper.AuthGrpcMapper,
	authGrpcValidator validator.AuthValidator,
) AuthGrpcService {
	return &authGrpcService{
		redis:             redis,
		authRepository:    authRepository,
		client:            client,
		logger:            logger,
		config:            config,
		authGrpcMapper:    authGrpcMapper,
		authGrpcValidator: authGrpcValidator,
	}
}
