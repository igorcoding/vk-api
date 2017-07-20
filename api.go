package vkapi

import (
	"errors"
	"log"
	"net/url"
	"os"
)

const (
	defaultVersion = "5.67"
	defaultScheme  = "https"
	defaultHost    = "api.vk.com"
	defaultPath    = "method"
	defaultMethod  = "GET"

	defaultHTTPS    = "1"
	defaultLanguage = LangEN

	paramVersion  = "v"
	paramLanguage = "lang"
	paramHTTPS    = "https"
	paramToken    = "access_token"
)

const (
	ErrApiClientNotFound = "APIClient not found."
)

const (
	LangRU = "ru" //Russian
	LangUA = "ua" //Ukrainian
	LangBE = "be" //Belarusian
	LangEN = "en" //English
	LangES = "es" //Spanish
	LangFI = "fi" //Finnish
	LangDE = "de" //German
	LangIT = "it" //Italian
)

// APIClient allows you to send requests to API server.
type APIClient struct {
	httpClient  HTTPClient
	APIVersion  string
	AccessToken *AccessToken
	secureToken string

	// If Log is true, APIClient will write logs.
	Log    bool
	Logger *log.Logger

	// HTTPS defines if use https instead of http. 1 - use https. 0 - use http.
	HTTPS string

	// Language define the language in which different data will be returned, for example, names of countries and cities.
	Language string
}

// SetAccessToken sets access token to APIClient.
func (api *APIClient) SetAccessToken(token string) {
	api.AccessToken = &AccessToken{
		AccessToken: token,
	}
}

// Values returns values from this APIClient.
func (api *APIClient) Values() (values url.Values) {
	values = url.Values{}
	values.Add(paramVersion, api.APIVersion)
	values.Add(paramLanguage, api.Language)
	values.Add(paramHTTPS, api.HTTPS)
	return
}

// Authenticate run authentication this APIClient from Application.
func (api *APIClient) Authenticate(application Application) (err error) {
	api.AccessToken, err = Authenticate(api, application)
	if err != nil {
		return err
	}

	if api.AccessToken.Error != "" {
		return errors.New(api.AccessToken.Error + " : " + api.AccessToken.ErrorDescription)
	}

	return nil
}

// NewApiClient creates a new *APIClient instance.
func NewApiClient() *APIClient {
	client := &APIClient{
		httpClient: defaultHTTPClient(),
		APIVersion: defaultVersion,
		Logger:     log.New(os.Stdout, "", log.LstdFlags),
		HTTPS:      defaultHTTPS,
		Language:   defaultLanguage,
	}

	return client
}

// ApiURL return standard url for interacting with server API.
func ApiURL() (url url.URL) {
	url.Host = defaultHost
	url.Path = defaultPath
	url.Scheme = defaultScheme
	return url
}

func (api *APIClient) logPrintf(format string, v ...interface{}) {
	if api.Log {
		api.Logger.Printf(format, v...)
	}
}

func (api *APIClient) SetHttpClient(httpClient HTTPClient) {
	api.httpClient = httpClient
}
