package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/yuhueijin/MyFirstGoAPI/service"
	"github.com/yuhueijin/MyFirstGoAPI/endpoint"

)

func MakeHandler(s service.Service) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	addHandler := httptransport.NewServer(
		endpoint.MakeAddEndpoint(s),
		decodeAddRequest,
		encodeAddResponse,
		options...,
	)

	removeHandler := httptransport.NewServer(
		endpoint.MakeRemoveEndpoint(s),
		decodeRemoveRequest,
		encodeRemoveResponse,
		options...,
	)

	getAllHandler := httptransport.NewServer(
		endpoint.MakeGetAllEndpoint(s),
		decodeGetAllRequest,
		encodeGetAllResponse,
		options...,
	)

	r := chi.NewRouter()
	r.Route("/items", func(r chi.Router) {
		r.Get("/", getAllHandler.ServeHTTP)
		r.Post("/add", addHandler.ServeHTTP)
		r.Get("/remove/{ID}", removeHandler.ServeHTTP)
	})

	return r
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrBadRequest
	}
	return request, nil
}

func decodeRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return nil, ErrInvalidId
	}
	return endpoint.RemoveRequest{
		ID: id,
	}, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

var ErrBadRequest = errors.New("bad request")
var ErrInvalidId = errors.New("invalid id")

func encodeAddResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoint.AddResponse)
	return json.NewEncoder(w).Encode(res.Err)
}

func encodeRemoveResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoint.RemoveResponse)
	return json.NewEncoder(w).Encode(res.Err)
}

func encodeGetAllResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoint.GetAllResponse)
	if res.Err != nil {
		return json.NewEncoder(w).Encode(res.Err)
	}
	return json.NewEncoder(w).Encode(res.Payload)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidId:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
