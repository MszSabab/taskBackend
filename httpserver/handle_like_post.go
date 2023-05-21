package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"

	"net/http"
	"strconv"
)

func (s *Server) handleLikePost() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		postID, _ := strconv.Atoi(chi.URLParam(req, "id"))

		fmt.Print(postID)
		err := s.db.LikePost(postID)
		if err != nil {
			log.Error().Msgf("db query Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("liked")
		w.WriteHeader(http.StatusOK)
	}
}
