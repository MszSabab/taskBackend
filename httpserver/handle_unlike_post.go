package httpserver

import (
	"net/http"
)

func (s *Server) handleUnlikePost() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		//postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
		//
		//_, err := db.Exec("UPDATE posts SET likes = likes - 1 WHERE id = $1", postID)
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//
		//w.WriteHeader(http.StatusOK)
	}
}
