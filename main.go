package main

import (
	"log"
	"net/http"
	"time"

	"github.com/arseto/go-clean-arch/context/greeting"
	"github.com/arseto/go-clean-arch/transport/httphandler"
	greetinghttp "github.com/arseto/go-clean-arch/transport/httphandler/greeting"
	"github.com/gorilla/mux"
)

func main() {

	muxrr := httphandler.NewMuxRequestReader()

	guc := greeting.NewGreetingUseCase()
	gh := greetinghttp.NewGreetingHandler(guc, muxrr)

	router := mux.NewRouter()
	router.HandleFunc("/greet/{name}", gh.Greet)

	log.Fatal(http.ListenAndServe(":8080", logger(router)))
}

func logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
