package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

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

func DoHTTPRequest(ctx context.Context, host, method, body, query string) (interface{}, error) {
	var req *http.Request
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), host, strings.NewReader(body))
	if query != "" {
		req.URL.RawQuery = query
	}
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &out)

	return out, err
}
