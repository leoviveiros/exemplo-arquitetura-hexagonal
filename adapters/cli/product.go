package cli

import (
	"fmt"
	"go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
		case "create": {
			product, err := service.Create(productName, productPrice)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf("Product %s created with ID %s, status %s and price %f", product.GetName(), product.GetID(), product.GetStatus(), product.GetPrice())
		}

		case "enable": {
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}
			
			res, err := service.Enable(product)

			if err != nil {
				return result, err
			}
 
			result = fmt.Sprintf("Product %s has been enabled", res.GetName())
		}

		case "disable": {
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}
			
			res, err := service.Disable(product)

			if err != nil {
				return result, err
			}
 
			result = fmt.Sprintf("Product %s has been disabled", res.GetName())
		}

		default: {
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf("Product %s with ID %s has status %s and price %f", product.GetName(), product.GetID(), product.GetStatus(), product.GetPrice())
		}
	}


	return result, nil
}
