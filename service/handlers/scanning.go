package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/models"
)

func (s ServiceHandler) CreateScanning(c *gin.Context) {
	ctx := c.Request.Context()
	log := common.GetLogger("CreateScanning", config.Values.Env)
	req := models.CreateScanningRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorln("CreateScanningError ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	lock, err := s.locker.Obtain(ctx, req.RepoName, 5*time.Second, nil)
	if err != nil {
		log.Errorln("ObtainLockErr ", err.Error())
		c.JSON(http.StatusTooManyRequests, gin.H{
			"message": "you request so fast, please wait some sec!",
		})
		return
	}
	defer lock.Release(ctx)
	queuedAt := time.Now()
	scanningModel := models.Scanning{
		RepoName:     req.RepoName,
		RepoURL:      req.RepoURL,
		ScanUniqueID: uuid.NewString(),
		Status:       models.QUEUED,
		QueuedAt:     &queuedAt,
	}

	err = scanningModel.SighUpScanning(s.db)
	if err != nil {
		log.Errorln("SighUpScanningErr ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	scanSignal := &models.Signal{
		ScanID: scanningModel.ScanUniqueID,
		Status: scanningModel.Status.String(),
	}
	err = s.cache.Publish(ctx, config.Values.Worker.Topic, common.Stringify(scanSignal)).Err()

	message := "ok"
	status := http.StatusOK

	if err != nil {
		log.Errorln("PublishRedisErr ", err.Error())
		message = "error!"
		status = http.StatusInternalServerError
	}
	c.JSON(status, gin.H{
		"message":        message,
		"scan_unique_id": scanningModel.ScanUniqueID,
	})
}

func (s ServiceHandler) ViewResult(c *gin.Context) {
	log := common.GetLogger("ViewResult", config.Values.Env)
	req := models.ViewResultRequest{}

	if err := c.BindQuery(&req); err != nil {
		log.Errorln("BindQueryErr ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid repo name",
		})
		return
	}
	if req.RepoName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid repo name",
		})
		return
	}

	scanningModel := models.Scanning{
		RepoName: req.RepoName,
	}
	result, err := scanningModel.ViewScanResultByRepoName(s.db)
	message := "ok"
	if err != nil {
		log.Errorf("ViewScanResultByRepoNameErr %s %s\n", req.RepoName, err.Error())
		message = "error!"
	}

	data := []models.ViewResultResponse{}
	for _, v := range result {
		if v.ID == 0 {
			continue
		}
		data = append(data, models.ViewResultResponse{
			ID:           v.ID,
			ScanUniqueID: v.ScanUniqueID,
			RepoName:     v.RepoName,
			RepoURL:      v.RepoURL,
			Status:       v.Status.String(),
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			QueuedAt:     v.QueuedAt,
			ScanningAt:   v.ScanningAt,
			FinishedAt:   v.FinishedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func (s ServiceHandler) RetryScan(c *gin.Context) {
	ctx := c.Request.Context()
	log := common.GetLogger("RetryScan", config.Values.Env)
	req := models.RetryScanProcessRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorln("CreateScanningError ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	scanningModel := models.Scanning{
		ScanUniqueID: req.ScanUniqueID,
	}
	exist, err := scanningModel.ViewOneScanningProcess(s.db)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "scanning not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error!",
		})
		return
	}

	if scanningModel.Status == models.IN_PROGRESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "scanning is running",
		})
		return
	}
	scanSignal := &models.Signal{
		ScanID: scanningModel.ScanUniqueID,
		Status: models.QUEUED.String(),
	}
	err = s.cache.Publish(ctx, config.Values.Worker.Topic, common.Stringify(scanSignal)).Err()
	status := http.StatusOK
	message := "ok"
	if err != nil {
		status = http.StatusInternalServerError
		message = "error!"
	}

	c.JSON(status, gin.H{
		"message": message,
	})
}

func (s ServiceHandler) ViewSpecificScanningProcess(c *gin.Context) {
	log := common.GetLogger("ViewSpecificScanningProcess", config.Values.Env)
	req := models.ViewScanningProcessRequest{}
	if err := c.BindQuery(&req); err != nil {
		log.Errorln("BindQueryErr ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid repo name",
		})
		return
	}
	if req.ScanUniqueID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid repo name",
		})
		return
	}
	scanningModel := models.Scanning{
		ScanUniqueID: req.ScanUniqueID,
	}
	exist, err := scanningModel.ViewOneScanningProcess(s.db)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "scanning not found",
		})
		return
	}
	status := http.StatusOK
	message := "ok"
	if err != nil {
		status = http.StatusInternalServerError
		message = "error!"
	}

	c.JSON(status, gin.H{
		"message": message,
		"data": models.ViewResultResponse{
			ID:           scanningModel.ID,
			ScanUniqueID: scanningModel.ScanUniqueID,
			RepoName:     scanningModel.RepoName,
			RepoURL:      scanningModel.RepoURL,
			Status:       scanningModel.Status.String(),
			CreatedAt:    scanningModel.CreatedAt,
			UpdatedAt:    scanningModel.UpdatedAt,
			QueuedAt:     scanningModel.QueuedAt,
			ScanningAt:   scanningModel.ScanningAt,
			FinishedAt:   scanningModel.FinishedAt,
		},
	})
}
