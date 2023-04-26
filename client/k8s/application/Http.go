package application

import (
	"net/http"
)

type HeaderAdder struct {
	rt http.RoundTripper
}

func (h *HeaderAdder) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json;as=Table;g=meta.k8s.io;v=v1")
	return h.rt.RoundTrip(req)
	
	//
}
