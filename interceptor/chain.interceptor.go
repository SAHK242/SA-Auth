package interceptor

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

// validateToken checks if the given JWT token is valid
func validateToken(tokenString string) bool {
	// Get the secret key from environment variables
	secretKey := "sasasasa"

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		fmt.Println("Invalid token:", err)
		return false
	}

	// Check expiration time (if needed)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims")
		return false
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			fmt.Println("Token expired")
			return false
		}
	}

	return true
}

// UnaryInterceptor checks if a request has a valid token
func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Get metadata from the request
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	// Get the token from metadata
	authHeaders := md["authorization"]
	if len(authHeaders) == 0 {
		return nil, fmt.Errorf("missing authorization token")
	}

	// Extract the token (assuming format "Bearer <token>")
	tokenParts := strings.Split(authHeaders[0], " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization format")
	}

	token := tokenParts[1]

	// Validate the token
	if !validateToken(token) {
		return nil, fmt.Errorf("invalid or expired token")
	}

	// Proceed with the request
	return handler(ctx, req)
}

// NewChainInterceptor creates the interceptor for gRPC server
func NewChainInterceptor() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(UnaryInterceptor)
}
