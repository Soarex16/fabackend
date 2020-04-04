package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/middlewares"
	"github.com/soarex16/fabackend/routes"
	"net/http"
)

// Handler - generic type for all handlers structs in this package
type Handler struct{}

// RouteParams - abstraction over mux libs
// At current moment - wrapper over httprouter
// TODO: по хорошему, надо запросы кешировать (в качестве ключа - requestID)
// 	если использовать map и мьютекс, то мы замедлим работу сервера из-за блокировок
//	 хотя если не ставить блокировки при чтении, то все должно быть ок
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

// ValidationErrors - type for describing errors in model
type ValidationErrors map[string]string

// ModelValidationError - writes validation errors as json
func (h *Handler) ModelValidationError(w http.ResponseWriter, r *http.Request, errs *ValidationErrors) {
	h.UnprocessableEntity(w, r)
	err := h.WriteJsonBody(w, r, errs)

	// error already has been logged in WriteJsonBody
	if err != nil {
		return
	}
}

// Ok - writes resp as JSON into body and send success
func (h *Handler) WriteJsonBody(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	bytes, err := json.Marshal(obj)

	// NOTE: can cause "superfluous response.WriteHeader call" because we must
	// write response code and all headers before writing body
	if err != nil {
		h.InternalServerError(w, r, err, "Error while serializing response")
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)

	return nil
}

// Ok - 200 (OK)
func (h *Handler) Ok(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusOK, "")
}

// Created - 201 (Created)
func (h *Handler) Created(w http.ResponseWriter, r *http.Request, resp interface{}) {
	w.WriteHeader(http.StatusCreated)
	err := h.WriteJsonBody(w, r, resp)

	if err != nil {
		return
	}

	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusCreated, "")
}

// BadRequest - 400 (Bad Request)
func (h *Handler) BadRequest(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusBadRequest, fmt.Sprintf("Bad request at route: %v", r.URL.Path))

	w.WriteHeader(http.StatusBadRequest)
}

// Unauthorized - 401 (Unauthorized)
func (h *Handler) Unauthorized(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusUnauthorized, fmt.Sprintf("Unauthorized user at route: %v", r.URL.Path))

	w.WriteHeader(http.StatusUnauthorized)
}

// Forbidden - 403 (Forbidden)
func (h *Handler) Forbidden(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusForbidden, fmt.Sprintf("Resource forbidden for user: %v", r.URL.Path))

	w.WriteHeader(http.StatusForbidden)
}

// NotFound - 404 (Not Found)
func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusNotFound, fmt.Sprintf("Cannot find resource at route: %v", r.URL.Path))

	w.WriteHeader(http.StatusNotFound)
}

// UnprocessableEntity - 422 (Unprocessable Entity)
func (h *Handler) UnprocessableEntity(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusUnprocessableEntity, fmt.Sprintf("Invalid request at route^ %v", r.URL.Path))

	w.WriteHeader(http.StatusUnprocessableEntity)
}

// InternalServerError - 500 (Internal Server Error)
func (h *Handler) InternalServerError(w http.ResponseWriter, r *http.Request, err error, errMsg string) {
	if errMsg == "" {
		errMsg = fmt.Sprintf("Error while processing request: %v", err)
	}

	reqId := h.RequestID(r)
	logErrorResponse(reqId, http.StatusInternalServerError, err, "")

	w.WriteHeader(http.StatusInternalServerError)
}

// NotImplemented - 501 (Not Implemented)
func (h *Handler) NotImplemented(w http.ResponseWriter, r *http.Request) {
	reqId := h.RequestID(r)
	logSuccessResponse(reqId, http.StatusNotImplemented, fmt.Sprintf("Not implemented at route: %v", r.URL.Path))

	w.WriteHeader(http.StatusNotImplemented)
}

func logSuccessResponse(reqId uuid.UUID, statusCode int, desc string) {
	if desc == "" {
		desc = "Successfully processed request"
	}

	logrus.
		WithField("requestId", reqId).
		WithField("statusCode", statusCode).
		Infof(desc)
}

func logErrorResponse(reqId uuid.UUID, statusCode int, err error, errMsg string) {
	if errMsg == "" {
		errMsg = fmt.Sprintf("Error while processing request: %v", err)
	}

	logrus.
		WithField("requestId", reqId).
		WithField("statusCode", statusCode).
		WithField("err", err).
		Error(errMsg)
}
