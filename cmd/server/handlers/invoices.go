package handlers

import (
	"app/internal/invoices/storage"
	"app/pkg/web/request"
	"app/pkg/web/response"
	"net/http"
	"time"
)

// NewControllerInvoice is a constructor for the invoice controller
func NewControllerInvoice(st storage.StorageInvoice) *ControllerInvoice {
	return &ControllerInvoice{st: st}
}

// ControllerInvoice is an invoice controller that returns handlers
type ControllerInvoice struct {
	st storage.StorageInvoice
}

// GetAll returns a handler for getting all invoices
type InvoiceResponseGetAll struct {
	Id         int       `json:"id"`
	Datetime   time.Time `json:"datetime"`
	Total      float64   `json:"total"`
	CustomerId int       `json:"customer_id"`
}
type ResponseBodyGetAllInvoices struct {
	Message string					 `json:"message"`
	Data    []*InvoiceResponseGetAll `json:"data"`
	Error   bool					 `json:"error"`
}
func (ct *ControllerInvoice) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		invoices, err := ct.st.ReadAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyGetAllInvoices{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetAllInvoices{Message: "Success", Data: make([]*InvoiceResponseGetAll, 0), Error: false}
		for _, inv := range invoices {
			body.Data = append(body.Data, &InvoiceResponseGetAll{
				Id:         inv.Id,
				Datetime:   inv.Datetime,
				Total:      inv.Total,
				CustomerId: inv.CustomerId,
			})
		}

		response.JSON(w, code, body)
	}
}

// Create returns a handler for creating an invoice
type RequestCreateInvoice struct {
	Datetime   time.Time `json:"datetime"`
	Total      float64   `json:"total"`
	CustomerId int       `json:"customer_id"`
}
type InvoiceResponseCreate struct {
	Id         int       `json:"id"`
	Datetime   time.Time `json:"datetime"`
	Total      float64   `json:"total"`
	CustomerId int       `json:"customer_id"`
}
type ResponseBodyCreateInvoice struct {
	Message string				   `json:"message"`
	Data    *InvoiceResponseCreate `json:"data"`
	Error   bool				   `json:"error"`
}

func (ct *ControllerInvoice) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestCreateInvoice
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := &ResponseBodyCreateInvoice{Message: "Invalid request body", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// process
		// -> deserialization
		inv := &storage.Invoice{
			Datetime:   reqBody.Datetime,
			Total:      reqBody.Total,
			CustomerId: reqBody.CustomerId,
		}
		if err := ct.st.Create(inv); err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyCreateInvoice{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyCreateInvoice{Message: "Success", Data: &InvoiceResponseCreate{
			Id:         inv.Id,
			Datetime:   inv.Datetime,
			Total:      inv.Total,
			CustomerId: inv.CustomerId,
		}, Error: false}

		response.JSON(w, code, body)
	}
}


// UpdateTotals returns a handler for updating the totals of all invoices
type ResponseBodyUpdateTotals struct {
	Message string `json:"message"`
	Data	any	   `json:"data"`
	Error   bool   `json:"error"`
}
func (ct *ControllerInvoice) UpdateTotals() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		if err := ct.st.UpdateTotals(); err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyUpdateTotals{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyUpdateTotals{Message: "Success", Data: nil, Error: false}

		response.JSON(w, code, body)
	}
}