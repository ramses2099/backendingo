package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ramses2099/backendingo/models"
)

//
var (
	posts []models.Post
)

//
func init() {
	posts = append(posts, models.Post{Id: 1, Title: "Title 1", Text: "Test text 1"})
}

//Note important only method with capital letter is import for other packages
//controller getallpost
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:"Error marshalling the posts array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

//w Response r Request
func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:"Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:"Error marshalling the post"}`))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result)

}
