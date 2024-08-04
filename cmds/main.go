package main

import (
	"fmt"
	"net/http"

	"github.com/ElijahB09/personal-website/views/home"
	"github.com/a-h/templ"
)

func main() {
	nav_items := []string{"Home", "Blog", "Project", "About"}

	component := home.Home(nav_items)

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
