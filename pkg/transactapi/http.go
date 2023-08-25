package transactapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func GetRequest[T any](ctx *context.Context, url string) (T, error) {
	var m T
	r, err := http.NewRequestWithContext(*ctx, http.MethodGet, url, nil)
	if err != nil {
		return m, err
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}

func PostRequest[T any](ctx *context.Context, url string, data any) (T, error) {
	var m T
	b, err := toJSON(data)
	if err != nil {
		return m, err
	}
	byteReader := bytes.NewReader(b)
	r, err := http.NewRequestWithContext(*ctx, http.MethodPost, url, byteReader)
	if err != nil {
		return m, err
	}
	// Important to set
	r.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}

func parseJSON[T any](s []byte) (T, error) {
	var r T
	if err := json.Unmarshal(s, &r); err != nil {
		return r, err
	}
	return r, nil
}

func toJSON(T any) ([]byte, error) {
	return json.Marshal(T)
}

type BaseRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
}
