package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkfmapf123/donggo"
)

func Test_HTTPCommInvalidMethod(t *testing.T) {
	atHTTP := NewATHTTP()

	url := "https://jsonplaceholder.typicode.com/posts"

	_, err := atHTTP.Comm(HTTPParams{
		Url:    url,
		Method: "PUT",
	})

	assert.Equal(t, err, errors.New("invalid method"))
}

func Test_HTTPComm(t *testing.T) {
	atHTTP := NewATHTTP()

	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := atHTTP.Comm(HTTPParams{
		Url:    url,
		Method: "POST",
		Body:   map[string]any{"title": "foo", "body": "bar", "userId": 1},
	})

	assert.NoError(t, err)

	type PostParams struct {
		ID int `json:"id"`
	}

	res := donggo.JsonParse[PostParams](resp)
	assert.NotZero(t, res.ID)
}
