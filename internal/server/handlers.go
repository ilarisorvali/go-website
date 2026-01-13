package server

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//Serve static files

	//Route handler for main/home page
	RenderTemplate(w, "home.html")
}
