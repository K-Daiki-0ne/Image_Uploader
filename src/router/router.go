package router

import (
	"0ne/src/controller"
	"net/http"
)

func Router() {
	http.HandleFunc("/upload", controller.Upload)
	http.ListenAndServe(":8080", nil)
}
