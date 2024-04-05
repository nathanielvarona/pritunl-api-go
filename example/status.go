package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {

	client := pritunl.PritunlClient()

	// You can also initialize an instance by manually providing the arguments.
	//
	// client := pritunl.PritunlClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL API URL>",
	// 	ApiToken:  "<PRITUNL API TOKEN>",
	// 	ApiSecret: "<PRITUNL API SECRET>",
	// })

	resp, err := client.StatusGet()

	if err != nil {
		log.Fatalf("Error performing status request: %v", err)
	}

	defer resp.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	fmt.Println(responseBody)

}
