package socket

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/configs"
	"fmt"
	"log"
	"net/http"
	"strings"

	io "github.com/ambelovsky/gosf-socketio"
)

func StartConnectionHandler() {
	server := configs.IoServer

	server.On(io.OnConnection, func(c *io.Channel) {
		userID := getUserID(c)
		if userID == "" {
			return
		}

		c.Join("user:" + userID)
		fmt.Println("Client Connected")
	})

	server.On(io.OnDisconnection, func(c *io.Channel) {
		fmt.Println("Client Disconnected")
	})

	server.On("join:room", func(c *io.Channel, roomID string) {
		c.Join("room:" + roomID)
		fmt.Println("Success to join room:" + roomID)
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", withCORS(server))

	log.Println("Socket.IO running on :8001")
	go http.ListenAndServe(":8001", mux)
}

func withCORS(next http.Handler) http.Handler {
	allowedOrigins := map[string]struct{}{}
	for _, origin := range []string{
		configs.AppConfig.App.Frontend_URL,
		"http://localhost:3000",
		"http://localhost:5173",
	} {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowedOrigins[origin] = struct{}{}
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getUserID(c *io.Channel) string {
	params := c.Request().URL.Query()
	token := params.Get("token")

	fmt.Println("Token received")

	claims, err := helpers.ValidateToken(token)
	if err != nil {
		fmt.Println("Token invalid:", err.Error())
		return ""
	}

	return claims.UserID
}
