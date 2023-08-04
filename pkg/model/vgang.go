package model

// Repository is the interface for the Vgang Api Repository
type VgangRepository interface {
	GetToken() (*GetTokenResponse, error)
	GetCategoriesIDs(token *GetTokenResponse) (IDs *[]int, err error)
	GetProductIDs(token *GetTokenResponse, category int, count int, offset int) (*GetProductIDsResponse, error)
}

type GetTokenResponse struct {
	Message string `json:"message"`
	Data    struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
}

type CategoriesResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	IsPopular bool   `json:"is_popular"`
	Image     string `json:"image"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Iso2 string `json:"iso2"`
	Iso3 string `json:"iso3"`
}

type Shipping struct {
	Country         string  `json:"country"`
	CountryISO2     string  `json:"countryISO2"`
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	Currency        string  `json:"currency"`
	MinDeliveryDate int     `json:"min_delivery_date"`
	MaxDeliveryDate int     `json:"max_delivery_date"`
}

type ProductImage struct {
	Image          string `json:"image"`
	ThumbnailImage string `json:"thumbnailImage"`
}

type Product struct {
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	SellerName       string         `json:"sellerName"`
	SellerCurrency   string         `json:"sellerCurrency"`
	SellerID         int            `json:"sellerID"`
	MinPrice         float64        `json:"minPrice"`
	MaxPrice         float64        `json:"maxPrice"`
	MinRetailPrice   float64        `json:"minRetailPrice"`
	MaxRetailPrice   float64        `json:"maxRetailPrice"`
	Stock            int            `json:"stock"`
	Country          Country        `json:"country"`
	ProductImages    []ProductImage `json:"productImages"`
	ImportListStatus string         `json:"importListStatus"`
	ImportListID     interface{}    `json:"importListID"`
	Shippings        []Shipping     `json:"shippings"`
	Tags             []string       `json:"tags"`
}

type ProductResponse struct {
	FreeShipping bool      `json:"freeShipping"`
	Products     []Product `json:"products"`
	Sellers      []string  `json:"sellers"`
	TotalCount   int       `json:"totalCount"`
}

type GetProductIDsResponse struct {
	ProductId  []int
	TotalCount int
}

type ProductIDs struct {
	ProductId []int
}
