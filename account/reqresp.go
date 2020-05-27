package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		ID    	int `json:"id"`
		Name	string `json:"name"`
		City	string `json:"city"`
		Phone	int `json:"phone"`
		Sal		int `json:"sal"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		Name string `json:"name"`
	}
)


func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetReq(r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		ID: vars["id"],
	}
	return req, nil
}