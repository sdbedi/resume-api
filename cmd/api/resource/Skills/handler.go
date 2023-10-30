package skills

import "net/http"

type API struct{}

func (a *API) List(w http.ResponseWriter, r *http.Request) {}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {}
