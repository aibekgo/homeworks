package main

import (
	"encoding/json"
	"fmt"
	"homeworks/lesson16/comments"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/comments/1"
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//from json to struct
	item := &comments.Comments{}
	err = json.Unmarshal(dataJson, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(item.PostId)
	//fmt.Println(item.Id)
	//fmt.Println(item.Body)
	//fmt.Println(item.Email)

	for i := 0; i < len(item); i++ {
		fmt.Println(item[i].PostId, item[i].Name)
	}
}

