package httpserver

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleCreatePost() http.HandlerFunc {

	type Comment struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
		User    string `json:"user"`
	}

	type Post struct {
		ID       int       `json:"id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		MediaURL string    `json:"mediaURL"`
		Likes    int       `json:"likes"`
		Comments []Comment `json:"comments"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		post := Post{}
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			// TODO add log
			return
		}

		s.db.Create(post.Title, post.Content, post.MediaURL, post.ID)
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		post.Likes = 0

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(post)
		if err != nil {
			// TODO add log
			return
		}
	}
}
