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

	service := application.ProductService{Persistence: application.IProductPersistent{persistence, nil}}

	result, err := service.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
