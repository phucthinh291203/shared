package interceptor

import (
	"context"
	"log"
	"strings"
	"user-service/errors"
	protobufpb "user-service/protobuf/gen/go"
	t "user-service/token"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	UsernameKey contextKey = "username"
	NameKey     contextKey = "name"
	RoleKey     contextKey = "role"
	UserIDKey   contextKey = "userid"
)

func RoleBasedInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println(info.FullMethod)
		//Admin
		if info.FullMethod == protobufpb.UserSrv_UploadAvatar_FullMethodName {
			return JWTAuthInterceptor([]string{"Admin"})(ctx, req, info, handler)
		}
		// else if info.FullMethod == protobufpb.UserSrv_List_FullMethodName {
		// 	return JWTAuthInterceptor([]string{""})(ctx, req, info, handler) //Không cần role nhưng phải login
		// }

		return handler(ctx, req) //không cần role
	}
}

func JWTAuthInterceptor(requiredRoles []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Lấy metadata từ context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.ErrorWithMessage(errors.ErrBadRequest, "Missing metadata")
		}

		// Lấy token từ header Authorization
		token := md["authorization"]
		if len(token) == 0 {
			return nil, errors.ErrorWithMessage(errors.ErrBadRequest, "Authorization header missing")
		}

		// Xóa "Bearer " khỏi token
		tokenString := strings.TrimPrefix(token[0], "Bearer ")

		// Giải mã và kiểm tra token
		claims, err := t.ParseJWT(tokenString)
		if err != nil {
			return nil, errors.ErrorWithMessage(errors.ErrInternalServerError, "ParseJWT failed")
		}

		// Kiểm tra vai trò của người dùng
		role := claims.Role
		roleAllowed := false
		for _, requiredRole := range requiredRoles {
			if role == requiredRole || requiredRole == "" {
				roleAllowed = true
				break
			}
		}
		if !roleAllowed {
			return nil, errors.ErrorWithMessage(errors.ErrInvalidToken, "Invalid token")
		}

		// Thêm thông tin người dùng vào context
		newCtx := context.WithValue(ctx, UsernameKey, claims.Username)
		newCtx = context.WithValue(newCtx, NameKey, claims.Name)
		if claims.Role == "Teacher" || claims.Role == "Admin" {
			newCtx = context.WithValue(newCtx, RoleKey, claims.Role)
			newCtx = context.WithValue(newCtx, UserIDKey, claims.UserID)
		}

		// Gọi handler với context mới
		return handler(newCtx, req)
	}
}
