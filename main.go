package main

import (
	"log"
	"net/http"

	"github.com/nishisuke/go-i18n-api/i18n"
)

func main() {
	http.HandleFunc("GET /greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		r = i18n.AcceptLanguageRequest(r) // 実際はmiddlewareで行う
		ctx := r.Context()

		name := r.PathValue("name")
		msg := i18n.Translatef(ctx, "greet", name)
		w.Write([]byte(msg))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
