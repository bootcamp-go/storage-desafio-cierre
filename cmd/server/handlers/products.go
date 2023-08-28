package handlers

import (
	"app/internal/products/storage"
	"app/pkg/web/request"
	"app/pkg/web/response"
	"net/http"
)

// NewControllerProduct is a constructor for the product controller
func NewControllerProduct(st storage.StorageProduct) *ControllerProduct {
	return &ControllerProduct{st: st}
}

// ControllerProduct is a product controller that return handlers
type ControllerProduct struct {
	st storage.StorageProduct
}

// GetAll returns a handler for getting all products
type ProductResponseGetAll struct {
	Id			int		`json:"id"`
	Description	string	`json:"description"`
	Price		float64	`json:"price"`
}
type ResponseBodyGetAllProducts struct {
	Message string					 `json:"message"`
	Data    []*ProductResponseGetAll `json:"data"`
	Error	bool					 `json:"error"`
}
func (ct *ControllerProduct) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		ps, err := ct.st.ReadAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyGetAllProducts{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetAllProducts{Message: "Success", Data: make([]*ProductResponseGetAll, 0), Error: false}
		for _, p := range ps {
			body.Data = append(body.Data, &ProductResponseGetAll{
				Id: p.Id,
				Description: p.Description,
				Price: p.Price,
			})
		}

		response.JSON(w, code, body)
	}
}

// Create returns a handler for creating a product
type RequestCreateProducts struct {
	Description	string	`json:"description"`
	Price		float64	`json:"price"`
}
type ProductResponseCreate struct {
	Id			int		`json:"id"`
	Description	string	`json:"description"`
	Price		float64	`json:"price"`
}
type ResponseBodyCreateProducts struct {
	Message string				   `json:"message"`
	Data    *ProductResponseCreate `json:"data"`
	Error	bool				   `json:"error"`
}
func (ct *ControllerProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestCreateProducts
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := &ResponseBodyCreateProducts{Message: "Invalid request body", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// process
		// -> deserialization
		p := &storage.Product{
			Description: reqBody.Description,
			Price: reqBody.Price,
		}
		if err := ct.st.Create(p); err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyCreateProducts{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyCreateProducts{Message: "Success", Data: &ProductResponseCreate{
			Id: p.Id,
			Description: p.Description,
			Price: p.Price,
		}, Error: false}

		response.JSON(w, code, body)
	}
}

// TopSelled returns a handler for getting the top selled products
type ProductResponseTopSelled struct {
	Description	string	`json:"description"`
	Quantity	int		`json:"quantity"`
}
type ResponseBodyTopSelledProducts struct {
	Message string					 `json:"message"`
	Data    []*ProductResponseTopSelled `json:"data"`
	Error	bool					 `json:"error"`
}
func (ct *ControllerProduct) TopSelled() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		ps, err := ct.st.TopSelled()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyTopSelledProducts{Message: "Internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyTopSelledProducts{Message: "Success", Data: make([]*ProductResponseTopSelled, 0), Error: false}
		for _, p := range ps {
			body.Data = append(body.Data, &ProductResponseTopSelled{
				Description: p.Description,
				Quantity: p.Quantity,
			})
		}

		response.JSON(w, code, body)
	}
}