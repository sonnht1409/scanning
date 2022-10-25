package workers

import (
	"context"
	"encoding/json"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"

	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/connectors/cache"
	"github.com/sonnht1409/scanning/service/connectors/db"
	"github.com/sonnht1409/scanning/service/handlers/logic"
	"github.com/sonnht1409/scanning/service/models"
)

type Worker struct {
	cache      *redis.Client
	db         *gorm.DB
	locker     *redislock.Client
	logicLayer logic.IServiceLogic
}

func NewServiceWorker() Worker {
	w := Worker{
		cache:      cache.NewCache(),
		db:         db.NewDB(),
		logicLayer: logic.NewServiceLogic(),
	}
	w.locker = cache.NewCacheLock(w.cache)
	return w
}

func (w Worker) Subscribe() {
	log := common.GetLogger("Worker", config.Values.Env)
	subscriber := w.cache.Subscribe(context.Background(), config.Values.Worker.Topic)
	for {
		mess, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			log.Errorf("ReceiveMessage Error :s \n", err.Error())
		}
		log.Infof("ReceiveMessage %+v", mess.Payload)

		signal := models.Signal{}
		err = json.Unmarshal([]byte(mess.Payload), &signal)
		if err == nil {
			if signal.Status == models.QUEUED.String() {
				w.RunQueuedRepo(context.Background(), signal.ScanID)
			}
		}
	}
}

func (w Worker) RunQueuedRepo(ctx context.Context, scanUniqueID string) {
	log := common.GetLogger("RunQueuedRepo", config.Values.Env)
	scannings := models.Scanning{}
	if err := scannings.GetByScanID(w.db, scanUniqueID); err != nil {
		log.Errorln("GetByScanIDErr ", err.Error())
		return
	}

	if err := scannings.UpdateScanningStatus(w.db, models.IN_PROGRESS); err != nil {
		log.Errorln("UpdateScanningStatusErr ", err.Error())
		return
	}

	w.CheckRepoContent(ctx, scannings)
}

func (w Worker) CheckRepoContent(ctx context.Context, scannings models.Scanning) {
	log := common.GetLogger("CheckRepoContent", config.Values.Env)
	repoOwner := w.logicLayer.GetRepoOwner(scannings.RepoURL)
	log.Infoln("RepoOwner", repoOwner)
	if repoOwner == "" {
		log.Warnln("GetRepoOwnerErr, repoUrl: ", scannings.RepoURL)
		return
	}

	contents, err := w.logicLayer.GetRepoContent(repoOwner, scannings.RepoName)
	if err != nil {
		log.Errorln("GetRepoContentErr, repoName: ", scannings.RepoName)
		return
	}

	if len(scannings.Findings) == 0 {
		scannings.Findings = []models.Finding{}
	}
	isPassed := true
	scannings.Status = models.SUCCESS
	for _, content := range contents {
		isViolatedPublicRule, indexes := w.logicLayer.CheckRule(content.Content, logic.PUBLIC_RULE)
		if isViolatedPublicRule {
			for _, indx := range indexes {
				scannings.Findings = append(scannings.Findings, models.Finding{
					RuleName: logic.PUBLIC_RULE.RuleName,
					Location: models.Location{
						Path: content.Path,
						Line: indx,
					},
				})
			}
			isPassed = false
		}
		isViolatedPrivateRule, indexes := w.logicLayer.CheckRule(content.Content, logic.PRIVATE_RULE)
		if isViolatedPrivateRule {
			for _, indx := range indexes {
				scannings.Findings = append(scannings.Findings, models.Finding{
					RuleName: logic.PRIVATE_RULE.RuleName,
					Location: models.Location{
						Path: content.Path,
						Line: indx,
					},
				})
			}
			isPassed = false
		}
	}
	if !isPassed {
		scannings.Status = models.FAILURE
	}
	err = scannings.UpdateScanningStatus(w.db, scannings.Status)
	if err != nil {
		log.Errorln("UpdateScanningStatusErr, repoName: ", scannings.RepoName)
		return
	}

	err = scannings.UpdateFindings(w.db)
	if err != nil {
		log.Errorln("UpdateFindingsErr, repoName: ", scannings.RepoName)
		return
	}
}
