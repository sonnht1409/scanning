package common

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"
)

func GetLogger(module, env string) *zap.SugaredLogger {
	var config zap.Config
	if strings.ToUpper(env) == "PROD" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level.UnmarshalText([]byte(env)) //nolint: errcheck
	log, _ := config.Build()                //nolint: errcheck

	return log.Named(module).Sugar()
}

func Stringify(in interface{}) string {
	out, err := json.Marshal(in)
	if err != nil {
		return ""
	}
	return string(out)
}
