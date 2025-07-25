package middleware

import (
	"booking-system/internal/backend/application/dto"
	"booking-system/internal/backend/domain/enum"
	"booking-system/internal/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	secret string
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{secret: secret}
}

// verifyToken checks if the JWT token is valid

func (m *JWTMiddleware) Verifytoken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		authFields := strings.Fields(authorization)
		if len(authFields) < 2 {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Authorization header is missing or malformed"),
			})
			ctx.Abort()
			return
		}

		if authFields[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Authorization header must start with Bearer"),
			})
			ctx.Abort()
			return
		}

		claims, err := utils.UnMarshalJwt(m.secret, authFields[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Invalid token: " + err.Error()),
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()

	}
}

// allow client
func (m *JWTMiddleware) allowClient(clients []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Claims not found in context"),
			})
			ctx.Abort()
			return
		}

		jwtClaims, ok := claims.(*dto.JWTClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, dto.ResponseAPI{
				Status:  dto.ApiExternalError.WithReason("Failed to cast claims to JWTClaims"),
				Message: "Invalid claims type",
			})
			ctx.Abort()
			return
		}

		for _, client := range clients {
			if jwtClaims.Client == client {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, dto.ResponseAPI{
			Status: dto.ApiStatusForbidden.WithReason("You do not have permission to access this resource"),
		})
		ctx.Abort()
	}
}

// allowRole

func (m *JWTMiddleware) allowRole(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Claims not found in context"),
			})
			ctx.Abort()
			return
		}

		jwtClaims, ok := claims.(*dto.JWTClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, dto.ResponseAPI{
				Status:  dto.ApiExternalError.WithReason("Failed to cast claims to JWTClaims"),
				Message: "Invalid claims type",
			})
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if jwtClaims.Role == role {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, dto.ResponseAPI{
			Status: dto.ApiStatusForbidden.WithReason("You do not have permission to access this resource"),
		})
		ctx.Abort()
	}
}

// isLogin
func (m *JWTMiddleware) IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimsValue, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
				Status: dto.ApiStatusUnauthorized.WithReason("Claims not found in context"),
			})
			ctx.Abort()
			return
		}

		claims, ok := claimsValue.(*dto.JWTClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, dto.ResponseAPI{
				Status:  dto.ApiExternalError.WithReason("Failed to cast claims to JWTClaims"),
				Message: "Invalid claims type",
			})
			ctx.Abort()
			return
		}

		if claims.IsLogin {
			ctx.Next()
			return
		}

		ctx.JSON(http.StatusUnauthorized, dto.ResponseAPI{
			Status: dto.ApiStatusUnauthorized.WithReason("You must be logged in to access this resource"),
		})
		ctx.Abort()
	}
}

// inject allow client
func (m *JWTMiddleware) AllowAllClient() gin.HandlerFunc {
	return m.allowClient([]string{enum.ClientWeb, enum.ClientMobile})
}

func (m *JWTMiddleware) AllowWebClient() gin.HandlerFunc {
	return m.allowClient([]string{enum.ClientWeb})
}
func (m *JWTMiddleware) AllowMobileClient() gin.HandlerFunc {
	return m.allowClient([]string{enum.ClientMobile})
}

func (m *JWTMiddleware) AllowAllRole() gin.HandlerFunc {
	return m.allowRole([]string{enum.RoleSuperAdmin, enum.RoleAdmin, enum.RoleUser})
}
func (m *JWTMiddleware) AllowSuperAdminRole() gin.HandlerFunc {
	return m.allowRole([]string{enum.RoleSuperAdmin})
}
func (m *JWTMiddleware) AllowAdminRole() gin.HandlerFunc {
	return m.allowRole([]string{enum.RoleAdmin})
}
func (m *JWTMiddleware) AllowUserRole() gin.HandlerFunc {
	return m.allowRole([]string{enum.RoleUser})
}
