package auth

import (
	dbcontext "Auction/internal/adapters/db"
	"Auction/internal/core/domain/entity"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func init() {
	gob.Register(webauthn.SessionData{})
}

type PassKeyService struct {
	db      *gorm.DB
	webAuth *webauthn.WebAuthn
	jwt     *JWTServices
}

func NewPasskeyService(db *dbcontext.PgContext, jwt *JWTServices) *PassKeyService {
	webauthn, err := webauthn.New(&webauthn.Config{
		RPDisplayName: "Auction",
		RPID:          "localhost",
		RPOrigins:     []string{"http://10.20.0.73", "http://localhost:8080"},
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &PassKeyService{
		db.Context(),
		webauthn,
		jwt,
	}
}

func (pks PassKeyService) BeginRegistration(c *gin.Context) {
	var username = c.Param("username")
	var user entity.User
	pks.db.First(&user, "email = ?", username)
	options, sessionData, err := pks.webAuth.BeginRegistration(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("sessionData", sessionData)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, options)
}

func (pks PassKeyService) FinishRegistration(c *gin.Context) {
	var username = c.Param("username")
	var user entity.User
	pks.db.First(&user, "email = ?", username)
	// Загружаем sessionData из сессии
	session := sessions.Default(c)
	fmt.Println("Session keys:", session.ID())

	sessionData, ok := session.Get("sessionData").(webauthn.SessionData)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session data not found"})
		return
	}

	credential, err := pks.webAuth.FinishRegistration(user, sessionData, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Credentials = append(user.Credentials, *credential)
	pks.db.Save(&user)
	// Очищаем sessionData из сессии
	session.Delete("sessionData")
	c.JSON(http.StatusOK, gin.H{"message": "Passkey registered!"})
}

func (pks PassKeyService) BeginLogin(c *gin.Context) {
	var username = c.Param("username")
	var user entity.User
	pks.db.Preload("Credentials").First(&user, "email = ?", username)

	options, sessionData, err := pks.webAuth.BeginLogin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("sessionData", sessionData)
	err = session.Save()
	c.JSON(http.StatusOK, options)
}

func (pks PassKeyService) FinishLogin(c *gin.Context) {
	var username = c.Param("username")
	var user entity.User
	pks.db.First(&user, "email = ?", username)
	session := sessions.Default(c)
	sessionData := session.Get("sessionData").(webauthn.SessionData)
	_, err := pks.webAuth.FinishLogin(user, sessionData, c.Request)
	token, err := pks.jwt.GenerateAccessToken(user.ID)
	refreshToken, err := pks.jwt.GenerateRefreshToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}
