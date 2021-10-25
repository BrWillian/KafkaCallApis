package routers

import (
	"encoding/json"
	"net/http"

	config "github.com/brwillian/kafka-consumer-api/config"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Version": "1.0.0"})
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	oracleConn := config.DbConn()
	w.Header().Set("Content-Type", "application/json")

	if !oracleConn {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return
}
func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
