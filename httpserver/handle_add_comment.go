package httpserver

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) handleAddComment() http.HandlerFunc {
	type Comment struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		UserName string `json:"user_name"`
	}
	type CommentResponse struct {
		Content  string `json:"content"`
		UserName string `json:"user_name"`
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		comment := Comment{}
		err := json.NewDecoder(req.Body).Decode(&comment)
		if err != nil {
			log.Error().Msgf("json.NewDecoder Error", err)
			return
		}

		postID, _ := strconv.Atoi(chi.URLParam(req, "id"))

		err = s.db.AddComment(comment.Content, comment.UserName, postID)
		if err != nil {
			log.Error().Msgf("db query error", err)
			return
		}

		commentRes := CommentResponse{
			Content:  comment.Content,
			UserName: comment.UserName,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(commentRes)
	}
}
