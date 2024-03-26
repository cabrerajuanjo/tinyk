package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// TODO: create fuctions to generate error messages appending the error detail in a "detail" field
var badRequest = "{\"message\":\"Bad request\"}"
var success = "{\"message\":\"Success\"}"

type Body struct {
	Url      string `json:"url"`
	SmallUrl string `json:"smallUrl"`
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Hello World"))
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	var body Body

	rbody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		w.Write([]byte(badRequest))
		return
	}

	fmt.Printf("res body: %s\n", string(rbody))

	if err := json.Unmarshal(rbody, &body); err != nil {
		println("Error parsing form\n")
		w.Write([]byte(badRequest))
		return
	}

	if body.Url == "" || body.SmallUrl == "" {
		fmt.Printf("Error: missing fields\n")
		w.Write([]byte(badRequest))
		return
	}

	fmt.Printf("Unmarshaled body %s\n", body)

	w.Write([]byte(success))
	spew.Dump(body)
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Hello World"))
}
