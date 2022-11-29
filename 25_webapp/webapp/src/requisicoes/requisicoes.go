package requisicoes

import (
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
)

func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	uri := config.API(url)

	request, err := http.NewRequest(metodo, uri, dados)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie["token"]))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
