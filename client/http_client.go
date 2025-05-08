package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HTTPGet(ctx context.Context, url string, headers map[string]string, data any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed create new http request: %s", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read from response body: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("url (%s) status code was %d, body: %s", url, resp.StatusCode, string(body))
	}

	return json.Unmarshal(body, data)
}

func HTTPPost(ctx context.Context, url string, headers map[string]string, payload, data any) error {
	p, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %s", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(p))
	if err != nil {
		return fmt.Errorf("failed create new http request: %s", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read from response body: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("url (%s) status code was %d, body: %s", url, resp.StatusCode, string(body))
	}

	return json.Unmarshal(body, data)
}
