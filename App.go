package main

import (
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

var Module = fx.Options(
	fx.Provide(
		NewMux,
	),
)
