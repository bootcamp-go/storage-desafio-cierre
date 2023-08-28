package handlers

import (
	"app/internal/customers/storage"
	"app/pkg/web/request"
	"app/pkg/web/response"
	"net/http"
)

// NewControllerCustomer is a constructor for the customer controller
func NewControllerCustomer(st storage.StorageCustomer) *ControllerCustomer {
	return &ControllerCustomer{
		storage: st,
	}
}

// ControllerCustomer is a customer controller that return handlers
type ControllerCustomer struct {
	storage storage.StorageCustomer
}

// GetAll returns a handler for getting all customers
type CustomerResponseGetAll struct {
	Id			int    `json:"id"`
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Condition	int    `json:"condition"`
}
type ResponseBodyGetAllCustomers struct {
	Message string					  `json:"message"`
	Data    []*CustomerResponseGetAll `json:"data"`
	Error	bool					  `json:"error"`
}
func (ct *ControllerCustomer) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		cs, err := ct.storage.ReadAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyGetAllCustomers{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetAllCustomers{Message: "Success", Data: make([]*CustomerResponseGetAll, 0), Error: false}
		for _, c := range cs {
			body.Data = append(body.Data, &CustomerResponseGetAll{
				Id: c.Id,
				FirstName: c.FirstName,
				LastName: c.LastName,
				Condition: c.Condition,
			})
		}

		response.JSON(w, code, body)
	}
}

// Create returns a handler for creating a customer
type RequestBodyCreateCustomers struct {
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Condition	int    `json:"condition"`
}
type CustomerResponseCreate struct {
	Id			int    `json:"id"`
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Condition	int    `json:"condition"`
}
type ResponseBodyCreateCustomers struct {
	Message string					`json:"message"`
	Data    *CustomerResponseCreate `json:"data"`
	Error	bool					`json:"error"`
}
func (ct *ControllerCustomer) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyCreateCustomers
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := &ResponseBodyCreateCustomers{Message: "Bad request", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// process
		// -> deserialization
		c := &storage.Customer{
			FirstName: reqBody.FirstName,
			LastName: reqBody.LastName,
			Condition: reqBody.Condition,
		}
		err := ct.storage.Create(c)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyCreateCustomers{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyCreateCustomers{Message: "Success", Data: &CustomerResponseCreate{
			Id: c.Id,
			FirstName: c.FirstName,
			LastName: c.LastName,
			Condition: c.Condition,
		}, Error: false}

		response.JSON(w, code, body)
	}
}

// ConditionInfo returns a handler for getting the condition info
type CustomerConditionInfoResponse struct {
	Condition int `json:"condition"`
	Total     int `json:"total"`
}
type ResponseBodyConditionInfo struct {
	Message string						 	 `json:"message"`
	Data    []*CustomerConditionInfoResponse `json:"data"`
	Error	bool							 `json:"error"`
}
func (ct *ControllerCustomer) ConditionInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		cs, err := ct.storage.ConditionInfo()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyConditionInfo{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyConditionInfo{Message: "Success", Data: make([]*CustomerConditionInfoResponse, 0), Error: false}
		for _, c := range cs {
			body.Data = append(body.Data, &CustomerConditionInfoResponse{
				Condition: c.Condition,
				Total: c.Total,
			})
		}

		response.JSON(w, code, body)
	}
}