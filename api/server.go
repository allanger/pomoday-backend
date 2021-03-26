package api

import (
	"net/http"

	"github.com/spf13/viper"
)

// Serve crack only
func Serve() {
	port := viper.GetString("port")
	err := NewRouter()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("Starting to serve on port ", port)
	log.Fatal(http.ListenAndServe(port, r))
}