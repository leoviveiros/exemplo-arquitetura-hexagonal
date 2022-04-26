package cli_test

import (
	"fmt"
	"go-hexagonal/adapters/cli"
	mock_application "go-hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := float64(25.99)
	productStatus := "enabled"
	productId := "12345"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product %s created with ID %s, status %s and price %f", productName, productId, productStatus, productPrice)

	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)

	result, err = cli.Run(serviceMock, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)

	result, err = cli.Run(serviceMock, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s with ID %s has status %s and price %f", productName, productId, productStatus, productPrice)

	result, err = cli.Run(serviceMock, "get", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}