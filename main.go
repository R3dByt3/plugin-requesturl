package plugin_requesturl

import (
	"context"
	"log"
	"net/http"
)

const defaultHeader = "X-Request-Url"

type Config struct {
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		requestURL := r.URL.String()

		r.Header.Add(config.HeaderName, requestURL)

		next.ServeHTTP(rw, r)
	}), nil
}
