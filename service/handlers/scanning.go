package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/models"
)

func (s ServiceHandler) CreateScanning(c *gin.Context) {
	uid := uuid.New().String()
	scanSignal := &models.Signal{
		ScanID: uid,
		Status: models.QUEUED.String(),
	}
	err := s.cache.Publish(context.Background(), config.Values.Worker.Topic, common.Stringify(scanSignal)).Err()
	message := "ok"
	status := http.StatusOK
	if err != nil {
		message = err.Error()
		status = http.StatusInternalServerError
	}
	c.JSON(status, gin.H{
		"message": message,
	})
}

func (s ServiceHandler) ViewResult(c *gin.Context) {
	log := common.GetLogger("ViewResult", config.Values.Env)
	log.Info("query: ", c.Query("name"))
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	s.logicLayer.GetRepoContent("sonnht1409", "capstone-etours")
}

func (s ServiceHandler) StopScanning(c *gin.Context) {
	log := common.GetLogger("StopScanning", config.Values.Env)
	log.Info("param: ", c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
