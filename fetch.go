package derpi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func apiGet(url string, params url.Values) ([]byte, error) {
	paramsString := params.Encode()

	if len(paramsString) > 0 {
		url += fmt.Sprintf("?%s", paramsString)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
