package httpserver

import (
	"net/http"
)

func (s *Server) handleGetPost() http.HandlerFunc {

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

		//strconv.Atoi(chi.URLParam(r, "id"))

		//page, _ := strconv.Atoi(req.URL.Query().Get("page"))
		//limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
		//
		//start := (page - 1) * limit
		//
		//rows, err := db.Query("SELECT * FROM posts LIMIT $1 OFFSET $2", limit, start)
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//defer rows.Close()
		//
		//posts := make([]Post, 0)
		//for rows.Next() {
		//	var post Post
		//	err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.MediaURL, &post.Likes)
		//	if err != nil {
		//		log.Println(err)
		//		continue
		//	}
		//	post.Comments, err = getComments(post.ID)
		//	if err != nil {
		//		log.Println(err)
		//		continue
		//	}
		//	posts = append(posts, post)
		//}

		//w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(posts)
	}
}
