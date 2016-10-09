package main

import (
        "fmt"
        "os"
	"net/http"
	"io/ioutil"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Println(os.Stderr, "Usage: ip URL\n")
                os.Exit(1)
        }

	res, err := http.Get(os.Args[1])

	if err != nil {
        	panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

}
