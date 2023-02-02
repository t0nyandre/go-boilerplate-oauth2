package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				elapsed := time.Since(start)
				logger.Infow(req.Method, "Path", req.URL.String(), "Duration", fmt.Sprintf("%.3f ms", float64(elapsed)/float64(time.Millisecond)))
			}()

			next.ServeHTTP(w, req)
		})
	}
}
