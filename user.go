package function

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		// depending on X-ACTION
		switch r.Header.Get("X-ACTION") {
		case "CREATE":
			createUser(r, w)
		case "SIGN_IN":
			signInEmailPassword(r, w)
		case "VERIFICATION_LINK":
			getVerificationLink(r, w)
		case "PASSWORD_RESET_LINK":
			getPasswordResetLink(r, w)
		default:
			respond(http.StatusBadRequest, map[string]interface{}{"error": "unsupported X-ACTION"}, w)
		}
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
	//ctx := context.Background()
	//// Authorization: Bearer [token]  (RFC 6750)
	//token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	//t, err := authClient.VerifyIDToken(ctx, token)
	//if err != nil {
	//	log.Fatalf("verify token: %v", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//userId := t.Claims["user_id"]
	//log.Printf("user_id = %s", userId)
	//
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil && err != io.EOF {
	//	log.Fatalf("read body: %v", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//var requestJson Json
	//err = json.Unmarshal(body, &requestJson)
	//if err != nil {
	//	log.Fatalf("json unmarshal: %v", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//fmt.Println(requestJson.Data)
	//
	//responseJson, err := json.Marshal(map[string]string{
	//	"data": "pong",
	//})
	//log.Printf("data: %s", responseJson)
	//if err != nil {
	//	log.Fatalf("marshal json: %v", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//fmt.Fprint(w, string(responseJson))
	respondStatus(http.StatusOK, w)
}

// createUser a user
func createUser(r *http.Request, w http.ResponseWriter) {
	for key, val := range r.Header {
		log.Printf("%v - %v", key, val)
	}
	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		Disabled(false)
	u, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	log.Printf("Successfully created user: %v\n", u)
	respondStatus(http.StatusCreated, w)
}

// updateUser a user
func updateUser(r *http.Request, w http.ResponseWriter) {
	params := (&auth.UserToUpdate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		Disabled(false)
	u, err := authClient.UpdateUser(context.Background(), "UID HERE>..", params)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	log.Printf("Successfully updated user: %v\n", u)
	respondStatus(http.StatusOK, w)
}

// deleteUser a user
func deleteUser(r *http.Request, w http.ResponseWriter) {
	err := authClient.DeleteUser(context.Background(), "UID HERE>..")
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	log.Printf("Successfully deleted user: %v\n", "UID...")
	respondStatus(http.StatusOK, w)
}

// get verification email link
func getVerificationLink(r *http.Request, w http.ResponseWriter) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	type payload struct {
		Email string `json:"email"`
	}
	var p payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	link, err := authClient.EmailVerificationLink(context.Background(), p.Email)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
	}
	respond(http.StatusBadRequest, map[string]interface{}{"verification_link": link}, w)
}

// get password reset link
func getPasswordResetLink(r *http.Request, w http.ResponseWriter) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	type payload struct {
		Email string `json:"email"`
	}
	var p payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	link, err := authClient.PasswordResetLink(context.Background(), p.Email)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
	}
	respond(http.StatusBadRequest, map[string]interface{}{"verification_link": link}, w)
}

// signInEmailPassword returns the auth token
func signInEmailPassword(r *http.Request, w http.ResponseWriter) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var p payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
		return
	}
	link, err := authClient.CustomTokenWithClaims (context.Background(), p.Email)
	if err != nil {
		respond(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}, w)
	}
	respond(http.StatusBadRequest, map[string]interface{}{"verification_link": link}, w)
}
