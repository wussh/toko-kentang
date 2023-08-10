package handler

import "ecommerce/features/product"

type ProductResponse struct {
	Title       string `validate:"required" json:"title" form:"title"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	UserID      uint   `json:"user_id" form:"user_id"`
}

type GetAllRespon struct {
	ID    uint   `json:"id" form:"id"`
	Title string `validate:"required" json:"title" form:"title"`
	Price uint   `validate:"required" json:"price" form:"price"`
	Image string `json:"image" form:"image"`
}

func ToResponse(data product.CoreProduct) GetAllRespon {
	return GetAllRespon{
		ID:    data.ID,
		Title: data.Title,
		Price: data.Price,
		// Description: data.Description,
		Image: data.Image,
	}
}

func ToResponseArr(data []product.CoreProduct) []GetAllRespon {
	res := []GetAllRespon{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
