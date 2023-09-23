package renderer

import (
	"html/template"
	"log"
	"net/http"
)

func RenderJoke(w http.ResponseWriter, joke interface{}) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Failed to load template:", err)
		return
	}

	tmpl.Execute(w, joke)
}
