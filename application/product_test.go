package application_test

import (
	"testing"

	"github.com/go-hexa/application"
	"github.com/stretchr/testify/require"
)

func TestProductEnabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enabled()

	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 0
	err = product.Enabled()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProductDisabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disabled()

	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 10
	err = product.Disabled()

	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.ENABLED
	product.Price = 10

	isValid, _ := product.IsValid()

	require.Equal(t, true, isValid)

	product.Price = 0
	isValid, _ = product.IsValid()

	require.Equal(t, false, isValid)
}

func TestProductGetID(t *testing.T) {
	product := application.Product{}
	product.ID = "123"

	require.Equal(t, "123", product.GetID())
}

func TestProductGetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"

	require.Equal(t, "Product Test", product.GetName())
}

func TestProductGetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(t, application.ENABLED, product.GetStatus())

	product.Status = application.DISABLED
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductGetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(t, 10.0, product.GetPrice())
}
