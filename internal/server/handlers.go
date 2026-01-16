package server

import (
	"net/http"

	"github.com/ilarisorvali/sorvali-systems/internal/models"
)

var path string = "./markdown/"

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	posts, _ := models.LoadMarkdownPosts(path)

	//fmt.Println(posts)

	app.render(w, r, http.StatusOK, "home.html", posts)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	posts, _ := models.LoadMarkdownPosts(path)

	//fmt.Println(posts)

	app.render(w, r, http.StatusOK, "posts.html", posts)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	posts, _ := models.LoadMarkdownPosts(path)

	//fmt.Println(posts)

	app.render(w, r, http.StatusOK, "post.html", posts)
}
