package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type HTTPParams struct {
	Url     string
	Method  string
	Body    map[string]any
	Headers map[string]any
}

type ATHTTP struct{}

func NewATHTTP() *ATHTTP {
	return &ATHTTP{}
}

func (a *ATHTTP) Comm(
	params HTTPParams,
) ([]byte, error) {

	var resp *http.Response
	var err error
	var body []byte

	switch params.Method {
	case "GET":
		resp, err = http.Get(params.Url)
	case "POST":
		resp, err = http.Post(params.Url, "application/json", bytes.NewBuffer(body))
	default:
		return nil, errors.New("invalid method")
	}

	if err != nil {
		return nil, errors.New("failed to send request")
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	return body, nil
}
