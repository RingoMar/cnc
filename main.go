package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type msg struct {
	Url string `json:"secure_m3u8_url"`
}

func main() {
	fmt.Println("Connecting to API for link.")

	url := "https://player-api.new.livestream.com/accounts/27001737/events/8190593/stream_info"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	returnData := msg{}
	jsonErr := json.Unmarshal(body, &returnData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	c := exec.Command("vlc", returnData.Url)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
