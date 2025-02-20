package main

import (
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		Module,
		fx.Invoke(func(mux *http.ServeMux) {
			fmt.Println("Server started on :8080")
			http.ListenAndServe(":8080", mux)
		}),
	).Run()
}
