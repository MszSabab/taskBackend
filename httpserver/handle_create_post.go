package httpserver

import (
	"encoding/json"
	"github.com/taskBackend/repository"
	"log"
	"net/http"
)

func (s *Server) handleCreatePost() http.HandlerFunc {

	//type Comment struct {
	//	ID      int    `json:"id"`
	//	Content string `json:"content"`
	//	User    string `json:"user"`
	//}

	type Post struct {
		ID       int                  `json:"id"`
		Title    string               `json:"title"`
		Content  string               `json:"content"`
		MediaURL string               `json:"mediaURL"`
		Likes    int                  `json:"likes"`
		Comments []repository.Comment `json:"comments"` // sub struct of comments declared in repository directory
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
			// TODO add log
			return
		}

		err = s.db.CreatePost(post.Title, post.Content, post.MediaURL)
		if err != nil {
			log.Println(err)
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
