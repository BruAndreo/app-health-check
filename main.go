package main

import (
	"apphealthmonitor/files"
	"fmt"
	"net/http"
	"time"
)

var FILE_NAME = "config.yaml"

func main() {
	configs := files.LoadConfig(FILE_NAME)

	watch(configs)

}

func watch(configs files.Configs) {
	for i := 1; i <= configs.Requests.Quantity; i++ {
		fmt.Println("-----------------------")
		fmt.Println("Start", i, "round")

		for _, site := range configs.Sites {
			statusCode, online := request(site)
			fmt.Println(site, statusCode, online)
		}

		if i < configs.Requests.Quantity {
			fmt.Println("Waiting next round")
			time.Sleep(time.Duration(configs.Requests.DelaySeconds) * time.Second)
		}
	}
}

func request(site string) (int, bool) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err.Error())
	}

	online := resp.StatusCode != http.StatusInternalServerError

	return resp.StatusCode, online
}
