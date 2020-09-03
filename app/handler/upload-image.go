package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/manh737/go-mux/app/helper"
	"github.com/manh737/go-mux/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post form request
type Post struct {
	Image string
}

// UploadImage will handle the upload image post request
func UploadImage(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	image := new(model.Image)
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
	timeNow := time.Now()
	image.CreatedAt = timeNow
	image.UpdatedAt = timeNow
	image.ID = primitive.NewObjectID()
	image.ImageLink = "uploads/" + image.ID.Hex() + ".jpg"
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			ResponseWriter(w, http.StatusNotAcceptable, "image already exists in database.", nil)
		default:
			ResponseWriter(w, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	error := helper.Base64toJpg(post.Image, image.ID.Hex())
	if error != "" {
		ResponseWriter(w, http.StatusBadRequest, error, nil)
		return
	}
	_, err = db.Collection("images").InsertOne(nil, image)
	ResponseWriter(w, http.StatusCreated, "", image)
}