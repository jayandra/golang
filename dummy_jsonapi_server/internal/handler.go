package internal

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

type RequestHandler struct{}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var methodHandler http.HandlerFunc
	switch r.Method {
	case http.MethodPost:
		methodHandler = h.createPost
	case http.MethodGet:
		methodHandler = h.listBlogs
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	methodHandler(w, r)
}

func (h *RequestHandler) listBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != jsonapi.MediaType {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
	}

	blogs := seedData()

	w.Header().Set("Content-type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, blogs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *RequestHandler) createPost(w http.ResponseWriter, r *http.Request) {
	post := new(Post)
	fmt.Printf("-----------createPost called-------------\n")
	if err := jsonapi.UnmarshalPayload(r.Body, post); err != nil {
		fmt.Printf("%+v\n", post)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", jsonapi.MediaType)
}
