package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) handleLikePost() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		postID, _ := strconv.Atoi(chi.URLParam(req, "id"))

		fmt.Print(postID, ">>>>>>>>...")
		err := s.db.LikePost(postID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
