package httphandler

import "net/http"

type RequestReader interface {
	GetRouteParam(r *http.Request, name string) string
	GetJsonData(r *http.Request, data interface{}) error
}
