package application_test

import (
	"testing"

	"github.com/go-hexa/application"
	mock_application "github.com/go-hexa/application/mocks"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	gmock "github.com/golang/mock/gomock"
)

func setupProductServiceTest(t *testing.T) (*gmock.Controller, *mock_application.MockProductInterface, *mock_application.MockProductPersistenceInterface, application.ProductService) {
	ctrl := gmock.NewController(t)

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	return ctrl, product, persistence, service
}

func TestProductService_Get(t *testing.T) {
	ctrl, product, _, service := setupProductServiceTest(t)
	defer ctrl.Finish()

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl, product, _, service := setupProductServiceTest(t)
	defer ctrl.Finish()

	result, err := service.Create("Product 1", 10.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductEnabled_Enable(t *testing.T) {
	ctrl, product, _, service := setupProductServiceTest(t)
	defer ctrl.Finish()

	product.EXPECT().Enabled().Return(nil).AnyTimes()

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductEnabled_Disable(t *testing.T) {
	ctrl, product, _, service := setupProductServiceTest(t)
	defer ctrl.Finish()

	product.EXPECT().Disabled().Return(nil).AnyTimes()

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
