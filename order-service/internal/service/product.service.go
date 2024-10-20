package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"order-service/global"
	"order-service/internal/constants"
	"order-service/pkg/response"
)

type IProductService interface {
	GetProductById(productId string) (*response.ProductResponse, error)
}

type ProductService struct {
}

func NewProductService() IProductService {
	return &ProductService{}
}

func (p ProductService) GetProductById(productId string) (*response.ProductResponse, error) {
	url := fmt.Sprintf("http://product-service:9140/api/v1/products/%s", productId)
	responseObject := &response.ResponseData{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		global.Logger.Error("Http Request product:", zap.Error(err))
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Error("Http Response product:", zap.Error(err))
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		global.Logger.Error("Http Response product:", zap.Int("StatusCode", resp.StatusCode))
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}
	responseObject.Data = &response.ProductResponse{}

	err = json.NewDecoder(resp.Body).Decode(responseObject)
	if err != nil {
		global.Logger.Error("Decoder product réponse:", zap.Error(err))
		return nil, errors.New(constants.MICROSERVICE_READ_RESP)
	}

	global.Logger.Info("Get product successfully", zap.String("ProductId", productId))
	return responseObject.Data.(*response.ProductResponse), nil
}
