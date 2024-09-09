package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hinotora/sse-notification-service/internal/logger"
	"github.com/hinotora/sse-notification-service/internal/repository"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)

	response["status"] = "ok"

	json, _ := json.Marshal(response)

	w.Write(json)
}

func GetApplicationConnections(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*logger.Logger)

	applicationId := r.PathValue("application_id")

	data := repository.GetApplicationConnections(applicationId)

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	write, err := json.Marshal(data)

	if err != nil {
		logger.Error(fmt.Sprintf("get conn app list err: %s", err))
	}

	w.Write(write)
}

func GetUserConnections(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*logger.Logger)

	applicationId := r.PathValue("application_id")
	userId := r.PathValue("user_id")

	data := repository.GetUsersConnections(applicationId, userId)

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	write, err := json.Marshal(data)

	if err != nil {
		logger.Error(fmt.Sprintf("get conn app list err: %s", err))
	}

	w.Write(write)
}
