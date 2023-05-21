package httpserver

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) handleUnlikePost() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		postID, _ := strconv.Atoi(chi.URLParam(req, "id"))

		err := s.db.UnLikePost(postID)
		if err != nil {
			log.Error().Msgf("db query Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("disliked")
		w.WriteHeader(http.StatusOK)
	}
}
