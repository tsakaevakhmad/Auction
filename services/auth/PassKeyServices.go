package auth

import (
	"Auction/domain/entity"
	"Auction/services/dbcontext"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type PassKeyService struct {
	db      *gorm.DB
	webAuth *webauthn.WebAuthn
}

func NewPasskeyService(db *dbcontext.PgContext) *PassKeyService {
	webauthn, err := webauthn.New(&webauthn.Config{
		RPDisplayName: "Auction",
		RPID:          "webauthn",
		RPOrigins:     []string{"https://webauthn.me"},
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &PassKeyService{
		db.Context(),
		webauthn,
	}
}

type registerRequest struct {
	Name string `json:"name"`
}

func (pks PassKeyService) BeginRegistration(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var user entity.User
	pks.db.First(&user, "email = ?", req.Name)
	options, sessionData, err := pks.webAuth.BeginRegistration(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("sessionData", sessionData)

	c.JSON(http.StatusOK, options)
}

func (pks PassKeyService) FinishRegistration(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var user entity.User
	pks.db.First(&user, "email = ?", req.Name)
	// Загружаем sessionData из сессии
	session := sessions.Default(c)
	sessionData, ok := session.Get("sessionData").(*webauthn.SessionData)
	if !ok || sessionData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session data not found"})
		return
	}

	credential, err := pks.webAuth.FinishRegistration(user, *sessionData, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Credentials = append(user.Credentials, *credential)
	pks.db.Save(&user)
	// Очищаем sessionData из сессии
	session.Delete("sessionData")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Passkey registered!"})
}

func (pks PassKeyService) BeginLogin(c *gin.Context) {
	/*user := getUserFromDB("user@example.com")

	options, sessionData, err := webAuthn.BeginLogin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	saveSessionData(sessionData)
	c.JSON(http.StatusOK, options)*/
}

func (pks PassKeyService) FinishLogin(c *gin.Context) {
	/*user := getUserFromDB("user@example.com")
	sessionData := getSessionData()

	credential, err := webAuthn.FinishLogin(user, *sessionData, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})*/
}
