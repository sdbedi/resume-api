package education

import "net/http"

type School struct{}
type API []School

func (a *API) List(w http.ResponseWriter, r *http.Request) {}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {}
