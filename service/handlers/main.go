package handlers

import (
	"fmt"
	"net/http"

	"github.com/bsm/redislock"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/connectors/cache"
	"github.com/sonnht1409/scanning/service/connectors/db"
	"github.com/sonnht1409/scanning/service/handlers/logic"
	"gorm.io/gorm"
)

type ServiceHandler struct {
	r          *gin.Engine
	db         *gorm.DB
	cache      *redis.Client
	logicLayer logic.IServiceLogic
	locker     *redislock.Client
}

func NewServiceHandlers() ServiceHandler {
	s := ServiceHandler{
		r:          gin.Default(),
		db:         db.NewDB(),
		cache:      cache.NewCache(),
		logicLayer: logic.NewServiceLogic(),
	}
	s.locker = cache.NewCacheLock(s.cache)
	return s
}

func (s ServiceHandler) ApiRegister() {
	s.r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	s.r.GET("/api/v1/scanning/results", s.ViewResult)
	s.r.GET("/api/v1/scanning/result", s.ViewSpecificScanningProcess)
	s.r.POST("/api/v1/scanning", s.CreateScanning)
	s.r.POST("/api/v1/scanning/retry", s.RetryScan)

}

func (s ServiceHandler) Start() {
	if config.Values.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	s.r.Run(fmt.Sprintf(":%s", config.Values.Application.Port)) //nolint: errcheck
}
