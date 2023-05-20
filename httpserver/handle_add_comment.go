package httpserver

import (
	"net/http"
)

func (s *Server) handleAddComment() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		//postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
		//
		//var comment Comment
		//_ = json.NewDecoder(r.Body).Decode(&comment)
		//
		//stmt, err := db.Prepare("INSERT INTO comments (post_id, content, user_name) VALUES ($1, $2, $3) RETURNING id")
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//
		//var commentID int
		//err = stmt.QueryRow(postID, comment.Content, comment.User).Scan(&commentID)
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//
		//comment.ID = commentID
		//
		//w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(comment)
	}
}
