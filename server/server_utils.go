package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ashwinath/simple/framework"
)

type HTTPResponse struct {
	StatusCode int
	Message    any
}

type CustomHTTPHandler func(*http.Request) HTTPResponse

func convertHandler(fw framework.FW, customHandler CustomHTTPHandler) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		res := customHandler(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.StatusCode)

		defer func(st time.Time, response HTTPResponse) {
			fw.GetLogger().Infof("status=%d, method=%s, route=%s, duration=%s", response.StatusCode, r.Method, r.URL.String(), time.Since(st).String())
		}(startTime, res)

		if err := json.NewEncoder(w).Encode(res.Message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, `{"error": "internal service error"}`)
			return
		}
	})
}

func Unmarshal(r *http.Request, d any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(d)
}

type HTTPStandardError struct {
	Message string `json:"message"`
}

func UnmarshalError(err error) HTTPResponse {
	return badRequest(fmt.Sprintf("unable to unmarshal request, %s", err))
}

func badRequest(message string) HTTPResponse {
	return httpError(http.StatusInternalServerError, message)
}

func InternalError(message string) HTTPResponse {
	return httpError(http.StatusBadRequest, message)
}

func httpError(code int, message string) HTTPResponse {
	return HTTPResponse{
		StatusCode: code,
		Message: HTTPStandardError{
			Message: message,
		},
	}
}

func OK(m any) HTTPResponse {
	return HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    m,
	}
}

func Ping(_ *http.Request) HTTPResponse {
	return OK(nil)
}
