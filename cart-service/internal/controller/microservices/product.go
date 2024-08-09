package microservices

import (
	"cart-service/internal/constants"
	"cart-service/internal/response"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func CallGetProduct(url, token string, productResponse *response.ProductResponse) (*response.ProductResponse, error) {
	responseObject := &response.ResponseObject{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println(resp)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}
	responseObject.Data = &response.ProductResponse{}

	err = json.NewDecoder(resp.Body).Decode(responseObject)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(constants.MICROSERVICE_FAIL)
	}

	return responseObject.Data.(*response.ProductResponse), nil
}
