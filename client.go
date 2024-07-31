package apieasy

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
)

// HTTPClientConfig defines the configuration for the HTTP client
type HTTPClientConfig struct {
	Timeout       time.Duration
	SkipTLSVerify bool
}

// NewHTTPClient creates a new HTTP client with the given configuration
func NewHTTPClient(config HTTPClientConfig) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipTLSVerify},
	}
	return &http.Client{
		Transport: tr,
		Timeout:   config.Timeout,
	}
}

type APIResponse struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

// SendRequest sends an HTTP request with optional headers and body, and returns the complete response or an error
func SendRequest(client *http.Client, method, rawURL string, headers map[string]string, body io.Reader) (*APIResponse, error) {
	req, err := http.NewRequest(method, rawURL, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(responseBody),
	}, nil
}

// AddURLParams adds URL parameters to the given URL
func AddURLParams(rawURL string, params map[string]string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
