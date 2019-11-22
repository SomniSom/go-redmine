package redmine

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	endpoint string
	apikey   string
	*http.Client
	Limit  int
	Offset int
	fhttp  *fasthttp.Client
}

var DefaultLimit int = -1  // "-1" means "No setting"
var DefaultOffset int = -1 //"-1" means "No setting"

func NewClient(endpoint, apikey string) *Client {
	return &Client{endpoint, apikey, http.DefaultClient, DefaultLimit, DefaultOffset, &fasthttp.Client{}}
}

// URLWithFilter return string url by concat endpoint, path and filter
// err != nil when endpoin can not parse
func (c *Client) URLWithFilter(path string, f Filter) (string, error) {
	var fullURL *url.URL
	fullURL, err := url.Parse(c.endpoint)
	if err != nil {
		return "", err
	}
	fullURL.Path += path
	if c.Limit > -1 {
		f.AddPair("limit", strconv.Itoa(c.Limit))
	}
	if c.Offset > -1 {
		f.AddPair("offset", strconv.Itoa(c.Offset))
	}

	if _, ok := f.filters["key"]; !ok {
		f.AddPair("key", c.apikey)
	}

	fullURL.RawQuery = f.ToURLParams()
	return fullURL.String(), nil
}

func (c *Client) getPaginationClause() string {
	clause := ""
	if c.Limit > -1 {
		clause = clause + fmt.Sprintf("&limit=%v", c.Limit)
	}
	if c.Offset > -1 {
		clause = clause + fmt.Sprintf("&offset=%v", c.Offset)
	}
	return clause
}

//easyjson:json
type errorsResult struct {
	Errors []string `json:"errors"`
}

//easyjson:json
type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//easyjson:json
type Id struct {
	Id int `json:"id"`
}
