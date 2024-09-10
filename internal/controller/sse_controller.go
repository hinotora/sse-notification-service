package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hinotora/sse-notification-service/internal/connection"
	"github.com/hinotora/sse-notification-service/internal/logger"
	"github.com/hinotora/sse-notification-service/internal/worker"
)

func OpenSSE(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	connection := ctx.Value("connection").(*connection.Connection)
	logger := ctx.Value("logger").(*logger.Logger)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("X-Connection-Id", connection.GetId())

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher := w.(http.Flusher)

	flusher.Flush()

	go worker.Pinger(connection, ctx)
	go worker.Writer(connection, ctx)

	logger.Debug("Main cycle started")

loop:
	for {
		select {
		case msg := <-connection.BroadcastCh:
			data, _ := json.Marshal(msg.Data)

			fmt.Fprintf(w, "id: %s\n", msg.Id)
			fmt.Fprintf(w, "event: %s\n", msg.Mtype)
			fmt.Fprintf(w, "data: %s", string(data))
			fmt.Fprint(w, "\n\n")

			flusher.Flush()

			logger.Debug(fmt.Sprintf("Outgoing: TYPE=%s DATA=%s", msg.Mtype, data))

			continue
		case <-connection.CloseCh:
			logger.Debug("Client disconnected")
			break loop
		}
	}

	flusher.Flush()

	logger.Debug("Main cycle stopped. End of connection")
}
