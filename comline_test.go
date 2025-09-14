package comline

import (
	"fmt"
	"testing"
)

func ExampleSendRequest() {
	action := ProcessReadAction{}
	response, err := SendRequest(LocalComline, Request{
		Action: action,
		Authentication: Extension{
			Description: "example extension",
			URL:         "https://docs.ray.pyrret.com",
		},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Data.Payload)
	}
}

func TestSendRequest(t *testing.T) {
	action := ProcessReadAction{}
	response, err := SendRequest(LocalComline, Request{
		Action: action,
		Authentication: Extension{
			Description: "example extension",
			URL:         "https://docs.ray.pyrret.com",
		},
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Log(response.Data.Payload)
	}
}