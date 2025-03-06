package auth

import (
	"Auction/domain/configurations"
	"Auction/domain/entity"
	"Auction/services/Configuration"
	"Auction/services/dbcontext"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

type JWTServices struct {
	cfg *configurations.JWTConfig
	db  *gorm.DB
}

func NewJWTServices(db *dbcontext.PgContext) *JWTServices {
	var cfg *configurations.MainConfig
	Configuration.ReadFile(&cfg)
	return &JWTServices{cfg: &cfg.JWTConfig, db: db.Context()}
}

func (service JWTServices) GenerateAccessToken(userId string) (string, error) {
	var user entity.User
	err := service.db.Preload("Roles").First(&user, "ID = ?", userId).Error
	if err != nil {
		log.Println("Error fetching user:", err)
		return "", err
	}
	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * time.Duration(service.cfg.ExpirationDateInMinutes)).Unix(), // Токен истекает через 24 часа
		"roles": roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.cfg.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service JWTServices) GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Minute * time.Duration(service.cfg.ExpirationDateInMinutes)).Unix(), // Токен истекает через 24 часа
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.cfg.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service JWTServices) VerifyRefreshToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(service.cfg.Secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, nil, errors.New("token expired")
			}
		} else {
			return nil, nil, errors.New("invalid token format")
		}
		return token, claims, nil
	}
	return nil, nil, errors.New("invalid token")
}

func (service JWTServices) VerifyAccessToken(tokenString string, roles ...string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(service.cfg.Secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, nil, errors.New("token expired")
			}
		} else {
			return nil, nil, errors.New("invalid token format")
		}

		if len(roles) > 0 {
			var userRoles []string
			if roles, ok := claims["roles"].([]interface{}); ok {
				for _, r := range roles {
					if roleStr, valid := r.(string); valid {
						userRoles = append(userRoles, roleStr)
					}
				}
			}
			if !ok {
				return nil, nil, errors.New("role not found in token")
			}

			if !contains(roles, userRoles) {
				return nil, nil, errors.New("user does not have required role")
			}
		}
		return token, claims, nil
	}
	return nil, nil, errors.New("invalid token")
}

func (service JWTServices) AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		_, claims, err := service.VerifyAccessToken(tokenString, requiredRoles...)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if userID, ok := claims["sub"].(string); ok {
			c.Set("userID", userID)
		}

		c.Next()
	}
}

func contains(slice []string, item []string) bool {
	for _, v := range slice {
		for _, v2 := range item {
			if strings.EqualFold(v, v2) {
				return true
			}
		}
	}
	return false
}
