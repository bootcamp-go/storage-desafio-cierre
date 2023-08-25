package handlers

import (
	"app/internal/sales/storage"
	"app/pkg/web/request"
	"app/pkg/web/response"
	"net/http"
)

// NewControllerSale is a constructor for the sale controller
func NewControllerSale(st storage.StorageSale) *ControllerSale {
	return &ControllerSale{st: st}
}

// ControllerSale is a sale controller that returns handlers
type ControllerSale struct {
	st storage.StorageSale
}

// GetAll returns a handler for getting all sales
type SaleResponseGetAll struct {
	Id         int `json:"id"`
	Quantity   int `json:"quantity"`
	ProductId  int `json:"product_id"`
	InvoiceId  int `json:"invoice_id"`
}
type ResponseBodyGetAllSales struct {
	Message string				  `json:"message"`
	Data    []*SaleResponseGetAll `json:"data"`
	Error   bool				  `json:"error"`
}
func (ct *ControllerSale) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		sales, err := ct.st.ReadAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyGetAllSales{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetAllSales{Message: "Success", Data: make([]*SaleResponseGetAll, 0), Error: false}
		for _, sale := range sales {
			body.Data = append(body.Data, &SaleResponseGetAll{
				Id:         sale.Id,
				Quantity:   sale.Quantity,
				ProductId:  sale.ProductId,
				InvoiceId:  sale.InvoiceId,
			})
		}

		response.JSON(w, code, body)
	}
}

// Create returns a handler for creating a sale
type RequestCreateSale struct {
	Quantity   int `json:"quantity"`
	ProductId  int `json:"product_id"`
	InvoiceId  int `json:"invoice_id"`
}
type SaleResponseCreate struct {
	Id         int `json:"id"`
	Quantity   int `json:"quantity"`
	ProductId  int `json:"product_id"`
	InvoiceId  int `json:"invoice_id"`
}
type ResponseBodyCreateSale struct {
	Message string              `json:"message"`
	Data    *SaleResponseCreate `json:"data"`
	Error   bool                `json:"error"`
}
func (ct *ControllerSale) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestCreateSale
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := &ResponseBodyCreateSale{Message: "Invalid request body", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// process
		// -> deserialization
		sale := &storage.Sale{
			Quantity:   reqBody.Quantity,
			ProductId:  reqBody.ProductId,
			InvoiceId:  reqBody.InvoiceId,
		}
		if err := ct.st.Create(sale); err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyCreateSale{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyCreateSale{Message: "Success", Data: &SaleResponseCreate{
			Id:         sale.Id,
			Quantity:   sale.Quantity,
			ProductId:  sale.ProductId,
			InvoiceId:  sale.InvoiceId,
		}, Error: false}

		response.JSON(w, code, body)
	}
}