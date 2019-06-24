package oauth

import (
	"fmt"
	"log"
	"net/http"
)

var oauthInterceptServer *http.Server

func StartServer() {
	http.HandleFunc("/hmm", handler)
	oauthInterceptServer = &http.Server{Addr: ":8080", Handler: nil}

	go func() {
		if err := oauthInterceptServer.ListenAndServe(); err != nil {
			log.Fatal("Failed to listen on port :8080")
		}
	}()
}

func StopServer() {
	if err := oauthInterceptServer.Shutdown(nil); err != nil {
		log.Fatal("Failed to shutdown oauthInterceptServer")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
