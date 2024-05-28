package handlers

import (
	"gothstarter/views/home"
	"net/http"
)

var items = []string{"item1", "item2", "item3"}

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index(items))
}
