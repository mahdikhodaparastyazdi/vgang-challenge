package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// vgangRepositoryImpl is the http implementation of vgangRepository
type vgangRepositoryImpl struct {
	config     *pkg.Config
	httpClient http.Client
}

// NewVgangRepository creates a new vgangRepository
func NewVgangRepository(config *pkg.Config) (model.VgangRepository, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	return &vgangRepositoryImpl{
		config: config,
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil

	//with  10 times Retry
	// client := *retryClient.StandardClient()
	// client.Timeout = 10 * time.Second
	// return &vangRepositoryImpl{
	// 	config:     config,
	// 	httpClient: client,
	// }, nil
}

func (o *vgangRepositoryImpl) GetToken() (*model.GetTokenResponse, error) {

	url := "https://vgang.io/api/vgang-core/v1/auth/login/retailer/vgang"
	method := "POST"

	payload := strings.NewReader(`{
    "DeviceId": "ed3c9cc1-4e2c-4d10-9623-8c4967e1bc94",
    "email": "mahdi.kho59@gmail.com",
    "password": "m59595959"
}`)

	client := o.httpClient
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var getTokenDataResoponse model.GetTokenResponse
	if err := json.Unmarshal([]byte(body), &getTokenDataResoponse); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &getTokenDataResoponse, nil
}

func (o *vgangRepositoryImpl) GetCategoriesIDs(token *model.GetTokenResponse) (*[]int, error) {
	url := "https://vgang.io/api/vgang-core/v1/common/categories"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.Data.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var getCategories []model.CategoriesResponse
	if err := json.Unmarshal([]byte(body), &getCategories); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var categoryIDs []int

	for _, category := range getCategories {
		categoryIDs = append(categoryIDs, category.ID)
	}
	return &categoryIDs, nil
}
func (o *vgangRepositoryImpl) GetProductIDs(token *model.GetTokenResponse, categoryID int, count int, offset int) (*model.GetProductIDsResponse, error) {
	// fmt.Println("offsest:", offset, "count", count)
	url := "https://vgang.io/api/vgang-core/v1/retailers/products?count=" + strconv.Itoa(count) +
		"&offset=" + strconv.Itoa(offset) + "&category=" + strconv.Itoa(categoryID) + "&dont_show_out_of_stock=1"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("eeerrr")
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.Data.AccessToken)
	// fmt.Println("&GetProductIDs3")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	// fmt.Println("&GetProductIDs2")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var response model.ProductResponse
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	// fmt.Println("&GetProductIDs1")
	var GetProductIDs model.GetProductIDsResponse
	GetProductIDs.TotalCount = response.TotalCount
	for _, product := range response.Products {
		GetProductIDs.ProductId = append(GetProductIDs.ProductId, product.ID)
	}
	// fmt.Println("&GetProductIDsssssssss")
	return &GetProductIDs, nil
}
