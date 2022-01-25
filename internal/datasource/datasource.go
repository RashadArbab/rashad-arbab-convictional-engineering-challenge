package datasource

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetData() ([]byte, error) {

	resp, err := http.Get("https://my-json-server.typicode.com/convictional/engineering-interview-api/products")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil

}
