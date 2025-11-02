package client

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkfmapf123/at-plan/usecase"
	"github.com/zkfmapf123/at-plan/utils"
)

func Test_ValidConfigFile(t *testing.T) {

	pwd, _ := utils.GetPwd()
	_, err := os.Create(filepath.Join(pwd, "atlantis_config_file.yaml"))
	if err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}

	defer os.Remove(filepath.Join(pwd, "atlantis_config_file.yaml"))

	at := getMock()
	err = at.ValidConfigFile()
	assert.NoError(t, err)
}

func Test_ValidRepository(t *testing.T) {

	at := getMock()
	err := at.ValidRepository()

	assert.NoError(t, err)
}

func getMock() AtlantisParams {
	return AtlantisParams{
		Request: &usecase.AtlantisRequest{
			AtlantisURL:        "https://atlantis.example.com",
			AtlantisToken:      "atlantis_token",
			AtlantisRepository: "zkfmapf123/atlantis-fargate",
			AtlantisConfigFile: "atlantis_config_file.yaml",
		},
		httpClient: utils.NewATHTTP(),
	}
}
