package google_books

import (
	"net/http"
)

type GoogleBookApi struct {
	Client      http.Client
	BaseUrl     string
	ApiKey      string
	Title       string
	Price       string
	Author      string
	Publisher   string
	Category    string
	Description string
}

func NewGoogleBookApi(api GoogleBookApi) *GoogleBookApi {
	return &GoogleBookApi{
		Client:      http.Client{},
		BaseUrl:     api.BaseUrl,
		ApiKey:      api.ApiKey,
		Title:       api.Title,
		Price:       api.Price,
		Author:      api.Author,
		Publisher:   api.Publisher,
		Category:    api.Category,
		Description: api.Description,
	}
}

// func (api *GoogleBookApi) GetTitle(ctx context.Context, title string) (books.Domain, error){
// 	uri:=api.BaseUrl + api.Title + title
// 	req,_:=http.NewRequest("GET", uri,nil)
// 	req.Header.Set(,api.ApiKey)
// }