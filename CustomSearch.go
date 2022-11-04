package go_googlesearchengine

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type CustomSearchConfig struct {
	Cx string
	Q  string
}

type CustomSearchResult struct {
	Items []CustomSearchResultItem `json:"items"`
}

type CustomSearchResultItem struct {
	Kind             string `json:"kind"`
	Title            string `json:"title"`
	HtmlTitle        string `json:"htmlTitle"`
	Link             string `json:"link"`
	DisplayLink      string `json:"displayLink"`
	Snippet          string `json:"snippet"`
	HtmlSnippet      string `json:"htmlSnippet"`
	FormattedUrl     string `json:"formattedUrl"`
	HtmlFormattedUrl string `json:"htmlFormattedUrl"`
}

func (service *Service) CustomSearch(cfg *CustomSearchConfig) (*CustomSearchResult, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("CheckEmailConfig must not be nil")
	}

	var values = url.Values{}

	values.Set("cx", cfg.Cx)
	values.Set("q", cfg.Q)

	var result CustomSearchResult

	var requestConfig = go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("customsearch/v1?%s", values.Encode())),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result, nil
}
