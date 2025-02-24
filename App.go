package main

import (
	"Auction/handlers/Category"
	"Auction/services/dbcontext"
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

var Module = fx.Options(
	fx.Provide(
		dbcontext.NewPgContext,
		Category.NewCreateCategoryHandler,
		NewMux,
	),
)
