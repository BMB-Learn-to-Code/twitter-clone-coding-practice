package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"twitter-clone-coding-practice/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var postPayload CreatePostPayload

	if err := readJSON(w, r, &postPayload); err != nil {
		writeJSONError(w, http.StatusBadRequest, fmt.Sprintf("it was not possible to parse the request body: %v", error.Error(err)))
		return
	}

	post := &store.Post{
		Title:   postPayload.Title,
		Content: postPayload.Content,
		// TODO: Implmement Auth to get the correct User id
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	postId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, fmt.Sprintf("invalid post id: %v", error.Error(err)))
		return
	}

	post, err := app.store.Posts.GetById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrorNotFound):
			writeJSONError(w, http.StatusNotFound, "post not found")
			return
		default:
			app.internalServerError(w, r, err)
		}
	}

	writeJSON(w, http.StatusOK, post)
}
