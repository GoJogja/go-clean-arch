package httphandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRequestReader struct{}

func NewMuxRequestReader() RequestReader {
	return &muxRequestReader{}
}

func (rr *muxRequestReader) GetRouteParam(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}
