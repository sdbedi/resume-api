package router

import (
	"github.com/go-chi/chi/v5"

	"myapp/api/resource/book"
	"myapp/api/resource/health"
	skills "resume/cmd/api/resource/Skills"
	"resume/cmd/api/resource/bio"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := &book.API{}
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)

		bioAPI := &bio.API{}
		r.Get("/bio", bioAPI.List)

		skillsAPI := &skills.API{}
		r.Get("/skills", skillsAPI.List)
		r.Get("/skills/{id}", bookAPI.Read)

		
	})

	return r
}
