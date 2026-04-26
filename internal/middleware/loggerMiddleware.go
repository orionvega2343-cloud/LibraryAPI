package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //Closure
		start := time.Now()                                                                                                                    //Mark a point start
		next.ServeHTTP(w, r)                                                                                                                   //Hand over control next
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr, "user", r.UserAgent(), "duration", time.Since(start)) //Create log
	})
}
