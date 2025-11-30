package main

import (
	"fmt"
	"net/http"

	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post store.Post
	var postPayload CreatePostPayload

	userId := 1

	if err := readJSON(w, r, postPayload); err != nil {
		writeJSONError(w, http.StatusBadRequest, fmt.Sprintf("it was not possible to parse the request body: %v", error.Error(err)))
		return
	}

	post.UserID = int64(userId)
	post.Title = postPayload.Title
	post.Content = postPayload.Content

	app.store.Posts.Create(ctx, &post)
}
