package services

import (
	"ecommerce/config/helper"
	"ecommerce/features/product"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator"
)

type productUseCase struct {
	qry product.ProductData
	vld *validator.Validate
}

func New(pd product.ProductData) product.ProductService {
	return &productUseCase{
		qry: pd,
		vld: validator.New(),
	}
}
func (puu *productUseCase) GetAll() ([]product.CoreProduct, error) {
	res, err := puu.qry.GetAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []product.CoreProduct{}, errors.New(msg)
	}
	return res, nil
}

func (puu *productUseCase) Add(newProduct product.CoreProduct, token interface{}, file *multipart.FileHeader) (product.CoreProduct, error) {
	id := helper.ExtractToken(token)
	fmt.Println("======service=====")
	if file != nil {
		if file.Size > 5000000 {
			return product.CoreProduct{}, errors.New("file size is too big")
		}

		formFile, err := file.Open()
		if err != nil {
			return product.CoreProduct{}, errors.New("open file error")
		}

		if !helper.TypeFile(formFile) {
			return product.CoreProduct{}, errors.New("use jpg or png type file")
		}
		defer formFile.Close()
		formFile, _ = file.Open()
		uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

		if err != nil {
			return product.CoreProduct{}, errors.New("server error")
		}

		newProduct.Image = uploadUrl
	}

	// err := cuu.vld.Struct(newProduct)
	// if err != nil {
	// 	log.Println(err)
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		log.Println(err)
	// 	}
	// 	return product.CoreProduct{}, errors.New("field required wajib diisi")
	// }
	res, err := puu.qry.Add(newProduct, uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "id tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return product.CoreProduct{}, errors.New(msg)
	}

	return res, nil

}

func (puu *productUseCase) GetById(token interface{}, idProduct uint) ([]product.CoreProduct, error) {
	idUser := helper.ExtractToken(token)
	res, err := puu.qry.GetById(uint(idUser), idProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []product.CoreProduct{}, errors.New(msg)
	}

	return res, nil
}
