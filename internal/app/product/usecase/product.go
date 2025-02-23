package usecase

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/repository"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseGetProduct, error)
	GetAllProducts() ([]dto.ResponseGetProduct, error)
	GetSpecificProduct(productID uuid.UUID) (dto.ResponseGetProduct, error)
	DeleteProduct(productID uuid.UUID) error
	UpdateProduct(request dto.RequestUpdateProduct) error
}

type ProductUsecase struct {
	ProductRepositoryI repository.ProductMySQLItf
}

func NewProductUsecase(productRepositoryI repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepositoryI: productRepositoryI,
	}
}

func (u *ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseGetProduct, error) {
	product := &entity.Product{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		PhotoUrl:    request.PhotoUrl,
	}

	err := u.ProductRepositoryI.Create(product)
	if err != nil {
		return dto.ResponseGetProduct{}, err
	}

	return product.ParseToDTO(), nil
}

func (u *ProductUsecase) GetAllProducts() ([]dto.ResponseGetProduct, error) {
	products := new([]entity.Product)

	err := u.ProductRepositoryI.GetAll(products)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseGetProduct, len(*products))
	for i, p := range *products {
		res[i] = p.ParseToDTO()
	}

	return res, nil
}

func (u *ProductUsecase) GetSpecificProduct(productID uuid.UUID) (dto.ResponseGetProduct, error) {
	product := &entity.Product{
		ID: productID,
	}

	err := u.ProductRepositoryI.GetSpecific(product)
	if err != nil {
		return dto.ResponseGetProduct{}, err
	}

	return product.ParseToDTO(), err
}

func (u *ProductUsecase) UpdateProduct(request dto.RequestUpdateProduct) error {
	product := &entity.Product{
		ID: request.ID,
	}

	err := u.ProductRepositoryI.GetSpecific(product)
	if err != nil {
		return err
	}

	product = &entity.Product{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		PhotoUrl:    request.PhotoUrl,
	}

	err = u.ProductRepositoryI.Updates(product)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUsecase) DeleteProduct(productID uuid.UUID) error {
	product := &entity.Product{
		ID: productID,
	}

	err := u.ProductRepositoryI.GetSpecific(product)
	if err != nil {
		return err
	}

	return u.ProductRepositoryI.Delete(product)
}
