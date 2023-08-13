package router

import (
	ctnR "github/angolarti/pluralcloud/cmd/api/container/router"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WelcomeResponse struct {
	Message string `json:"message"`
}

func Home(app chi.Router) {

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcome to Plural Cloud!"))
	})

}

func SetupRouter(app chi.Router) {
	// Crie um novo grupo de rotas com um prefixo
	api := chi.NewRouter()
	api.Route("/", func(api chi.Router) {
		Home(api)
		ctnR.ContainerRouter(api)
	})

	app.Mount("/", api)
}
