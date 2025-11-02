package client

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zkfmapf123/at-plan/usecase"
	"github.com/zkfmapf123/at-plan/utils"
	"github.com/zkfmapf123/donggo"
)

/*
@reference
https://www.runatlantis.io/docs/api-endpoints
*/
var (
	API_PLAN   = "/api/plan"
	API_HEALTH = "/healthz"
)

type AtlantisParams struct {
	Request    *usecase.AtlantisRequest
	httpClient *utils.ATHTTP
}

func NewAtlantisRequest(atParams usecase.AtlantisRequest) AtlantisParams {
	return AtlantisParams{
		Request:    &atParams,
		httpClient: utils.NewATHTTP(),
	}
}

func (a *AtlantisParams) ValidURL() error {

	resp, err := a.httpClient.Comm(
		utils.HTTPParams{
			Url:    a.Request.AtlantisURL + API_HEALTH,
			Method: "GET",
		},
	)

	if err != nil {
		return err
	}

	res := donggo.JsonParse[usecase.APIHealthResponse](resp)
	if res.Status != "ok" {
		return errors.New("atlantis is not running")
	}

	return nil
}

func (a *AtlantisParams) ValidRepository() error {

	gitAttr := strings.Split(a.Request.AtlantisRepository, "/")
	owner, repository := gitAttr[0], gitAttr[1]

	resp, err := a.httpClient.Comm(
		utils.HTTPParams{
			Url:    fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repository),
			Method: "GET",
		},
	)

	if err != nil {
		return err
	}

	res := donggo.JsonParse[any](resp)
	if res == nil {
		return errors.New("repository not found")
	}

	return nil
}

func (a *AtlantisParams) ValidConfigFile() error {

	pwd, _ := utils.GetPwd()
	atFilePath := filepath.Join(pwd, a.Request.AtlantisConfigFile)

	_, err := os.Stat(atFilePath)
	if os.IsNotExist(err) {
		return errors.New("atlantis config file not found")
	}

	// rewrite
	a.Request.AtlantisConfigFile = atFilePath
	return nil
}
