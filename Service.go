package go_googlesearchengine

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

const (
	apiName string = "GoogleSearchEngine"
	apiURL  string = "https://customsearch.googleapis.com"
)

type Service struct {
	apiKey        string
	httpService   *go_http.Service
	errorResponse *ErrorResponse
}

type ServiceConfig struct {
	ApiKey string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("ApiKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiKey:      serviceConfig.ApiKey,
		httpService: httpService,
	}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add Api key
	_url, err := url.Parse(requestConfig.Url)
	if err != nil {
		return nil, nil, errortools.ErrorMessage(err)
	}
	query := _url.Query()
	query.Set("key", service.apiKey)

	// add error model
	service.errorResponse = &ErrorResponse{}
	requestConfig.ErrorModel = service.errorResponse

	(*requestConfig).Url = fmt.Sprintf("%s://%s%s?%s", _url.Scheme, _url.Host, _url.Path, query.Encode())

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		if service.errorResponse.Error.Message != "" {
			e.SetMessage(service.errorResponse.Error.Message)
		}
	}

	return request, response, e
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
