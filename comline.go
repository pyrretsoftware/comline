package comline

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const (
	LocalComline = ""
)

//LocalComlinePath sets the path used when interacting with the local ray server comline, for if ray server has been installed to a custom location. 
func LocalComlinePath(path string) {
	localComlineAddress = path
}

var localComlineAddress = ""
func getLocalComlineAddress() (string, error) {
	address := localComlineAddress
	if address == "" {
		switch runtime.GOOS {
		case "windows":
			dir, err := os.UserHomeDir()
			if err != nil {return "", err}

			address = filepath.Join(dir, "ray-env", "comsock.sock")
		case "linux":
			address = "/usr/bin/ray-env/comsock.sock"
		default:
			return "", errors.New("unsupported platform, linux and windows are supported")
		}
	}

	if _, err := os.Stat(address); err != nil {
		return "", errors.New("the local comsocket does not exist at the expected path")
	}
	return address, nil
}

//GetClient returns a http.Client that can be used to make requests to a comline based on if the comline is accessed over a unix socket or tcp, and the address to use with it. Use the constant LocalComline as the address for local comlines.
//
//This is only useful for low level interactions with comlines, for most cases you'd want to just use SendRequest.
func GetClient(address string) (*http.Client, string, error) {
	unix := false
	dummyAddress := address
	if address == LocalComline {
		unix = true
		_addr, err := getLocalComlineAddress()
		if err != nil {return nil, "", err}
		address = _addr
		dummyAddress = "http://comline-dummy-address"
	}
	transport := &http.Transport{
		DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
			if unix {
				network = "unix"
				addr = address
			}
			return net.Dial(network, addr)
		},
	}
	
	return &http.Client{
		Transport: transport,
	}, dummyAddress, nil
}

//SendRequestRaw sends a raw request to the comline at the provided address, returning a raw response and any errors the occured. Use the constant LocalComline as the address for local comlines.
//
//This is only useful for low level interactions with comlines, for most cases you'd want to just use SendRequest.
func SendRequestRaw(address string, req RawComRequest) (RawComResponse, error) {
	c, addr, err := GetClient(address)
	if err != nil {
		return RawComResponse{}, err
	}

	ba, err := json.Marshal(req)
	if err != nil {
		return RawComResponse{}, err
	}
	resp, err := c.Post(addr, "application/json", bytes.NewReader(ba))
	if err != nil {
		return RawComResponse{}, err
	}

	rba, err := io.ReadAll(resp.Body)
	if err != nil {
		return RawComResponse{}, err
	}

	var response RawComResponse
	jerr := json.Unmarshal(rba, &response)
	if jerr != nil {
		return RawComResponse{}, jerr
	}

	if resp.StatusCode != 200 {
		return RawComResponse{}, errors.New("comline reported error: " + response.Data.Error)
	}
	return response, nil
}

//SendRequestRaw sends a request to the comline at the provided address, returning a raw response and any errors the occured. Use the constant LocalComline as the address for local comlines.
func SendRequest(address string, req Request) (RawComResponse, error) {
	actionString, payload := req.Action.FormatAction()
	return SendRequestRaw(address, RawComRequest{
		Action: actionString,
		Payload: payload,
		Key: req.Authentication.GetKey(),
	})
}