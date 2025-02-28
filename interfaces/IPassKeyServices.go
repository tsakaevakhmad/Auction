package interfaces

import "github.com/gin-gonic/gin"

type IPassKeyServices interface {
	BeginRegistration(context *gin.Context)
	FinishRegistration(context *gin.Context)
	BeginLogin(context *gin.Context)
	FinishLogin(context *gin.Context)
}
