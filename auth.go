package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
	projectID = os.Getenv("PROJECT_ID")
)

type Json struct {
	Data string `json:"data"`
}

var (
	authClient *auth.Client
)

func init() {
	var err error

	config := &firebase.Config{
		ProjectID: projectID,
	}
	firebaseApp, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err = firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("firebaseApp.Auth: %v", err)
	}
}

// User is our main function.
func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(r, w)
	case http.MethodPost:
		createUser(r, w)
	case http.MethodPut:
		updateUser(r, w)
	case http.MethodDelete:
		deleteUser(r, w)
	default:
		respond(http.StatusBadRequest, map[string]interface{}{"error": "unsupported http verb"}, w)
	}
}

// getUser a user
func getUser(r *http.Request, w http.ResponseWriter) {
	ctx := context.Background()
	// Authorization: Bearer [token]  (RFC 6750)
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	t, err := authClient.VerifyIDToken(ctx, token)
	if err != nil {
		log.Fatalf("verify token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userId := t.Claims["user_id"]
	log.Printf("user_id = %s", userId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		log.Fatalf("read body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestJson Json
	err = json.Unmarshal(body, &requestJson)
	if err != nil {
		log.Fatalf("json unmarshal: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(requestJson.Data)

	responseJson, err := json.Marshal(map[string]string{
		"data": "pong",
	})
	log.Printf("data: %s", responseJson)
	if err != nil {
		log.Fatalf("marshal json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(responseJson))
}

// createUser a user
func createUser(r *http.Request, w http.ResponseWriter) {
	respondStatus(http.StatusOK, w)
}

// updateUser a user
func updateUser(r *http.Request, w http.ResponseWriter) {
	respondStatus(http.StatusOK, w)
}

// deleteUser a user
func deleteUser(r *http.Request, w http.ResponseWriter) {
	respondStatus(http.StatusOK, w)
}
