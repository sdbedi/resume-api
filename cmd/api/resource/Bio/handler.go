package bio

import "net/http"

type API struct{
	Name string
	Email string
	Linkedin string
	
}

func (a *API) List(w http.ResponseWriter, r *http.Request) {}
