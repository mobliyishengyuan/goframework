package library

import (
	"net/http"
)

var router *http.ServeMux = http.NewServeMux()

func AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(pattern, handler)
}

func GetRouter() *http.ServeMux {
	return router
}