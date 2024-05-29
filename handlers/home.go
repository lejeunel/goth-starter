package handlers

import (
	"fmt"
	"gothstarter/views/home"
	"net/http"
)

var items = []string{"item1", "item2", "item3"}

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	var items = []string{}
	for i := range 80 {
		items = append(items, fmt.Sprintf("item %v", i))
	}
	return Render(w, r, home.Index(items))
}
