package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Linker struct {
	Short string
	Link  string
	Token string 
}

func (l Linker) Create() (Linker, error) {

	if l.Link == "" {
		return l, errors.New("empty link")
	}

	values := map[string]string{"link": l.Link}
    json_data, err := json.Marshal(values)

    if err != nil {
        return l, err
    }

	resp, err := http.Post("https://ls.johannespour.de/create", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return l, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return l, err
	}

	res := make(map[string]string)
	json.Unmarshal(body, &res)

	l.Short = res["short"]
	l.Token = res["token"]

	return l, nil
}