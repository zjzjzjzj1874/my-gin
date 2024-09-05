package main

import (
	"context"

	myclient "github/zjzjzjzj1874/my-gin/openapi"

	"github.com/sirupsen/logrus"
)

func main() {
	//NewConfiguration
	cfg := myclient.NewConfiguration()
	cfg.Host = "localhost:28888"
	cfg.Scheme = "http"
	apiClient := myclient.NewAPIClient(cfg)

	list, resp, err := apiClient.AccountsAPI.AccountGet(context.Background()).Execute()
	if err != nil {
		logrus.Errorf("failure:%d,err:%v", resp.StatusCode, err.Error())
		return
	}

	logrus.Info(list)
}
