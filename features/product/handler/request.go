package handler

import "ecommerce/features/product"

type AddReq struct {
	Title       string `validate:"required" json:"title" form:"title"`
	Category    string `validate:"required" json:"category" form:"category"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

func ToCore(data interface{}) *product.CoreProduct {
	res := product.CoreProduct{}

	switch cnv := data.(type) {
	case AddReq:
		res.Title = cnv.Title
		res.Category = cnv.Category
		res.Price = cnv.Price
		res.Description = cnv.Description
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
