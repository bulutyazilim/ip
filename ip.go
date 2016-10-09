package main

import (
        "fmt"
	"net/http"
	"io/ioutil"
)

func main() {
        url := "http://whatismyip.akamai.com/"

	res, err := http.Get(url)

	if err != nil {
        	panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

}
