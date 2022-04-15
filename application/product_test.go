package application_test

import (
	"go-hexagonal/application"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.NotNil(t, err)
}

func TesProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.NotNil(t, err)
}

func TestIsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Status = "INVALID"

	isValid, err = product.IsValid()
	require.NotNil(t, err)
	require.False(t, isValid)

	product.Status = application.ENABLED

	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Price = -10

	isValid, err = product.IsValid()
	require.NotNil(t, err)
	require.False(t, isValid)
}