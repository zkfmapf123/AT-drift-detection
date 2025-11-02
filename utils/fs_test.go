package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPWD(t *testing.T) {
	dir, err := GetPwd()
	if err != nil {
		t.Fatalf("failed to get pwd: %v", err)
	}

	assert.NotEqual(t, dir, "")
}
