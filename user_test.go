package function

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//
//import (
//	"bytes"
//	"github.com/buger/jsonparser"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)

func TestGet(t *testing.T) {
	t.Log("testing GET....")
}

func TestPost(t *testing.T) {
	t.Log("testing POST....")
	// create a user
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(User)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

//func TestGet(t *testing.T) {
//
//	// create a user
//	req, err := http.NewRequest("POST", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(TheFunction)
//	handler.ServeHTTP(rr, req)
//
//	body, err := ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	userID, _, _, err := jsonparser.Get(body, "id")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Printf("Created user with id %s", userID)
//
//
//	req, err = http.NewRequest("GET", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	//finally delete the user
//	req, err = http.NewRequest("DELETE", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	// testing without the id returns a bad request
//	req, err = http.NewRequest("GET", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusBadRequest {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusBadRequest)
//	}
//
//	// testing with a not found id results in not found
//	req, err = http.NewRequest("GET", "?id=XYZ", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusNotFound {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusNotFound)
//	}
//}
//
//
//func TestPost(t *testing.T) {
//	req, err := http.NewRequest("POST", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(TheFunction)
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusCreated {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusCreated)
//	}
//
//	body, err := ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	userID, _, _, err := jsonparser.Get(body, "id")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Printf("Created user with id %s", userID)
//	// cleanup the user
//	req, err = http.NewRequest("DELETE", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = http.HandlerFunc(TheFunction)
//	handler.ServeHTTP(rr, req)
//	log.Printf("User Deleted .... id %s", userID)
//}
//
//func TestPut(t *testing.T) {
//
//	// create a new user
//	req, err := http.NewRequest("POST", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(TheFunction)
//	handler.ServeHTTP(rr, req)
//
//	body, err := ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	userID, _, _, err := jsonparser.Get(body, "id")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Printf("Created user with id %s", userID)
//
//	var jsonStr = []byte(`{"id":"`+string(userID)+`", "name": "Jon Snow"}`)
//	req, err = http.NewRequest("PUT", "/", bytes.NewBuffer(jsonStr))
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// check to see if it updated ok
//	req, err = http.NewRequest("GET", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	body, err = ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	updatedName, _, _, err := jsonparser.Get(body, "name")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if string(updatedName) != "Jon Snow" {
//		t.Fatal("update failed!")
//	}
//
//	req, err = http.NewRequest("DELETE", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	log.Printf("Deleting user with id %s", userID)
//}
//
//func TestDelete(t *testing.T) {
//	// create a new user
//	req, err := http.NewRequest("POST", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(TheFunction)
//	handler.ServeHTTP(rr, req)
//
//	body, err := ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	userID, _, _, err := jsonparser.Get(body, "id")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Printf("Created user with id %s", userID)
//
//
//	req, err = http.NewRequest("DELETE", "?id="+string(userID), nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr = httptest.NewRecorder()
//	handler = TheFunction
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	log.Printf("Deleting user with id %s", userID)
//}