package routes

import (
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/handler"
)

func HandleHttpRequest(router *mux.Router){
	router.HandleFunc("/create", handler.GetHash)
	router.HandleFunc("/toko/{pattern}", handler.RedirectURL)
}
