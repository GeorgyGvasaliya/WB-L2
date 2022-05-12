package httpserver

import (
	"log"
	"net"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		elapsed := time.Since(start)
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			log.Printf("error parsing ip: " + err.Error())
			return
		}

		log.Printf("%s [%s] %s %s %d %s", ip,
			start.Format("02/Jan/2006:03:04:05 Z0700"),
			req.Method,
			req.RequestURI,
			elapsed.Microseconds(),
			req.UserAgent(),
		)
	})
}
