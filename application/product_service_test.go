package application_test

import (
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	mockapp "github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapp.NewMockIProduct(ctrl)
	persistence := mockapp.NewMockIProductReader(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: application.IProductPersistent{IProductReader: persistence}}

	result, err := service.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapp.NewMockIProduct(ctrl)
	persistence := mockapp.NewMockIProductWriter(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: application.IProductPersistent{IProductWriter: persistence}}

	result, err := service.Create("Product 1", 10)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapp.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()

	persistence := mockapp.NewMockIProductWriter(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: application.IProductPersistent{IProductWriter: persistence}}

	result, err := service.Enable(product)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapp.NewMockIProduct(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()

	persistence := mockapp.NewMockIProductWriter(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: application.IProductPersistent{IProductWriter: persistence}}

	result, err := service.Disable(product)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
