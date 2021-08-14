package main

import (
	"encoding/json"
	"fmt"
	"homeworks/lesson16/comments"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/comments"
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json; charset=UTF-8'")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	items := []comments.Comments{}
	err = json.Unmarshal(dataJson, &items)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i].Id, items[i].Name)
	}
}