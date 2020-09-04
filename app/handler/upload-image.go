package handler

import (
	"encoding/json"
	"net/http"
)

// Post form request
type Post struct {
	Image string
}

// UploadImage will handle the upload image post request
func UploadImage(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	if post.Image == "" {
		ResponseWriter(w, http.StatusBadRequest, "Image not null!!!", nil)
		return
	}
	ResponseWriter(w, http.StatusCreated, "", post.Image)
}
