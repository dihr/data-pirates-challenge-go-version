package correiosapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type (
	CorreiosApiSvc interface {
		GetHomePage() (*http.Response, error)
		GetData(uf string, startPage interface{}, endPage interface{}) (*http.Response, error)
	}

	correiosApiAimp struct {
		baseURl string
	}
)

func NewCorreiosApiService() CorreiosApiSvc {
	return &correiosApiAimp{
		baseURl: "https://www2.correios.com.br/sistemas/buscacep",
	}
}

func (c *correiosApiAimp) GetHomePage() (*http.Response, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buscaFaixaCep.cfm", c.baseURl), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *correiosApiAimp) GetData(uf string, startPage interface{}, endPage interface{}) (*http.Response, error) {
	var payload *strings.Reader

	if startPage != nil && endPage != nil {
		payload = strings.NewReader(fmt.Sprintf("UF=%s&Localidade=**&Bairro=&qtdrow=50&pagini=%d&pagfim=%d", uf,
			startPage.(int), endPage.(int)))
	} else {
		payload = strings.NewReader(fmt.Sprintf("UF=%s&Localidade=**&Bairro=&qtdrow=50&", uf))
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ResultadoBuscaFaixaCEP.cfm", c.baseURl), payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://www2.correios.com.br")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
