package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetPokeData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d and \nbody: %s\n", res.StatusCode, body)
		return nil, errors.New("invalid response received")
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
