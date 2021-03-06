package goceps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type service struct{}

type Service interface {
	Search(zipcode string) (*Address, error)
}

func NewService() Service {
	return new(service)
}

func (s *service) Search(zipcode string) (*Address, error) {
	if strings.Contains(zipcode, "-") {
		zipcode = strings.Replace(zipcode, "-", "", len(zipcode))
	}

	if len(zipcode) != 8 {
		return nil, errors.New("Oops, zipcode must be 8 characters")
	}

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipcode)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Oops, error on search address")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var address *Address

	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return address, err
}
