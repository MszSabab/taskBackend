package httpserver

import (
	"encoding/json"
	"github.com/rs/zerolog/log"

	"net/http"
)

func (s *Server) handleCreatePost() http.HandlerFunc {
	type Comment struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		UserName string `json:"user_name"`
	}
	type Post struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		MediaURL string `json:"mediaURL"`
		Likes    int    `json:"likes"`
	}
	type PostResponse struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		MediaURL string `json:"mediaURL"`
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		post := Post{}
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			log.Error().Msgf("json.NewDecoder Error", err)
			return
		}

		err = s.db.CreatePost(0, post.Title, post.Content, post.MediaURL)
		if err != nil {
			log.Error().Msgf("db query Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		PostRes := PostResponse{
			Title:    post.Title,
			Content:  post.Content,
			MediaURL: post.MediaURL,
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(PostRes)
		if err != nil {
			// TODO add log
			return
		}
	}
}
