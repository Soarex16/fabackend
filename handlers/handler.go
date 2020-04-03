package handlers

import (
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/soarex16/fabackend/middlewares"
	"github.com/soarex16/fabackend/routes"
	"net/http"
)

// Handler - generic type for all handlers structs in this package
type Handler struct {
	/**
	TODO:
		200 ok
		201 created
		400 bad request + model validation
		401 unauthorized
		403 forbidden
		404 not found
		500 server error
		501 not implemented
	*/
}

// RouteParams - abstraction over mux libs
// At current moment - wrapper over httprouter
// TODO: по хорошему, надо запросы кешировать (в качестве ключа - requestID)
func (h *Handler) RouteParams(r *http.Request) routes.RouteParams {
	paramsSlice, ok := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	if !ok {
		return nil
	}

	params := make(routes.RouteParams)
	for _, entry := range paramsSlice {
		params[entry.Key] = entry.Value
	}

	return params
}

// RequestID - extracts the request id from context
func (h *Handler) RequestID(r *http.Request) uuid.UUID {
	logCtx, ok := r.Context().Value(middlewares.RequestIDContextKey).(middlewares.RequestIDContext)

	if !ok {
		return uuid.Nil
	}

	return logCtx.ID
}
