package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homeworks/lesson16/comments"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/comments"
	//create special body for comments
	comments1 := &comments.Comments{
		PostId: 1,
		Name:   "lary",
		Email:  "lary@kazmail.kaz",
		Body:   "bullding",
	}
	//from struct to json
	jsonData, err := json.Marshal(comments1)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(jsonData)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
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
	//from json to struct
	item := &comments.Comments{}
	err = json.Unmarshal(dataJson, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item.Id)
	fmt.Println(item.Name)
	fmt.Println(item.Email)
	fmt.Println(item.Body)
}
