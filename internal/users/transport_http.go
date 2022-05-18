package users

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(svc Service, logger log.Logger) *mux.Router {
	opt := options(logger)

	createHandler := httptransport.NewServer(
		CreateEndpoint(svc),
		decodeCreate,
		encodeResponse,
		opt...,
	)

	updateHandler := httptransport.NewServer(
		UpdateEndpoint(svc),
		decodeUpdate,
		encodeResponse,
		opt...,
	)

	getHandler := httptransport.NewServer(
		GetEndpoint(svc),
		decodeGet,
		encodeResponse,
		opt...,
	)

	listHandler := httptransport.NewServer(
		ListEndpoint(svc),
		decodeList,
		encodeResponse,
		opt...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path(Create).Handler(createHandler)
	r.Methods("PUT").Path(Update).Handler(updateHandler)
	r.Methods("GET").Path(Get).Handler(getHandler)
	r.Methods("GET").Path(List).Handler(listHandler)

	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r

}

func decodeCreate(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UsersCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdate(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UsersUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	id, err := pathParametersID(r)
	if err != nil {
		return request, err
	}
	request.User.SetID(int64(id))

	return request, nil
}

func decodeGet(ctx context.Context, r *http.Request) (interface{}, error) {

	id, err := pathParametersID(r)
	if err != nil {
		return nil, err
	}

	return UsersGetRequest{
		User: &User{ID: int64(id)},
	}, nil
}

func decodeList(ctx context.Context, r *http.Request) (interface{}, error) {
	return &UsersListRequest{}, nil
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func options(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
}

func pathParametersID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	toInt, err := strconv.Atoi(id)

	return toInt, err
}
