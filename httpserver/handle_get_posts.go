package httpserver

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) handleGetPosts() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(req.URL.Query().Get("offset"))

		posts, err := s.db.GetAllPosts(limit, offset)
		if err != nil {
			log.Error().Msgf("db query Error", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}
