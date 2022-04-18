package application_test

import (
	"testing"

	"go-hexagonal/application"
	mock_application "go-hexagonal/application/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	peristence := mock_application.NewMockProductPersistenceInterface(ctrl)
	peristence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: peristence,
	}

	result, err := service.Get("1")

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	peristence := mock_application.NewMockProductPersistenceInterface(ctrl)
	peristence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: peristence,
	}

	result, err := service.Create("Product 1", 10.0)

	require.Nil(t, err)
	require.Equal(t, product, result)
}