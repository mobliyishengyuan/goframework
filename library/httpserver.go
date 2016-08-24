package library

import (
	"net/http"
	"log"
)

func InitHttpServer(port string) {
	var router = GetRouter()

	err := http.ListenAndServe(":" + port, router)
	
	if err != nil {
		log.Fatal("library InitHttpServer: ", err)
	}
}