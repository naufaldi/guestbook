package app

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/naufaldi/guestbook/internal/handler"
)

func (a *App) loadRoutes(tmpl *template.Template) {
	guestbook := handler.New(a.logger, a.db, tmpl)

	files := http.FileServer(http.Dir("./static"))

	a.router.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		http.StripPrefix("/static/", files).ServeHTTP(w, r)
	}))

	a.router.Handle("GET /{$}", http.HandlerFunc(guestbook.Home))

	a.router.Handle("POST /{$}", http.HandlerFunc(guestbook.Create))
}
