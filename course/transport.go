package course

import (
	"fmt"
	"github.com/LuisRiveraBan/gocourse_domain/domain"
	c "github.com/ncostamagna/go_http_client/client"
	"net/http"
	"net/url"
	"time"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Code    int         `json:"code"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHTTP struct {
		client c.Transport
	}
)

func NewHttpClient(baseURL, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c *clientHTTP) Get(id string) (*domain.Course, error) {
	dataResponse := DataResponse{Data: &domain.Course{}}

	u := url.URL{}

	u.Path += fmt.Sprintf("/courses/%s", id)

	reps := c.client.Get(u.String())

	if err := reps.FillUp(&dataResponse); err != nil {
		return nil, fmt.Errorf("%s", reps)
	}

	if reps.Err != nil {
		return nil, reps.Err
	}
	if reps.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", dataResponse.Message)}
	}

	if reps.StatusCode > 299 {
		return nil, fmt.Errorf("%s", dataResponse.Message)
	}

	return dataResponse.Data.(*domain.Course), nil
}
