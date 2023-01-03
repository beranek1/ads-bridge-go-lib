package adsbridgegolib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct used for managing ADSBridge with its address addr
type ADSBridge struct {
	addr string
}

// Creates instance of type ADSBridge with given address and checks connection, returns ADSBridge and error if connection fails.
func Connect(addr string) (ADSBridge, error) {
	var b ADSBridge
	b.addr = addr
	_, err := b.GetVersion()
	if err != nil {
		return b, err
	}
	return b, nil
}

// Reads and converts JSON response of ADSBridge
func processResponse(r io.Reader) (map[string]interface{}, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return nil, err
	}
	return dat, nil
}

// Performs GET request with ADSBridge, converts and returns result
func (b ADSBridge) Get(path string) (map[string]interface{}, error) {
	var url = b.addr + path
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return processResponse(resp.Body)
}

// Performs POST request with ADSBridge, converts and returns result
func (b ADSBridge) Post(path string, jsonStr string) (map[string]interface{}, error) {
	var url = b.addr + path
	resp, err := http.Post(url, "text/json", bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return processResponse(resp.Body)
}

func (b ADSBridge) GetVersion() (map[string]interface{}, error) {
	return b.Get("/version")
}

func (b ADSBridge) GetState() (map[string]interface{}, error) {
	return b.Get("/state")
}

func (b ADSBridge) GetDeviceInfo() (map[string]interface{}, error) {
	return b.Get("/deviceInfo")
}

func (b ADSBridge) GetSymbolInfo(name string) (map[string]interface{}, error) {
	return b.Get("/getSymbolInfo/" + name)
}

func (b ADSBridge) GetSymbolValue(name string) (map[string]interface{}, error) {
	return b.Get("/getSymbolValue/" + name)
}

func (b ADSBridge) ListSymbols() (map[string]interface{}, error) {
	return b.Get("/symbolList")
}

func (b ADSBridge) SetSymbolValue(name string, value string) (map[string]interface{}, error) {
	return b.Post("/setSymbolValue/"+name, "{\"data\":"+value+"}")
}

func (b ADSBridge) WriteControl(adsState uint16, deviceState uint16) (map[string]interface{}, error) {
	if adsState != 0 {
		if deviceState != 0 {
			return b.Post("/writeControl", "{\"adsState\":"+fmt.Sprint(adsState)+","+"\"deviceState\":"+fmt.Sprint(deviceState)+"}")
		} else {
			return b.Post("/writeControl", "{\"adsState\":"+fmt.Sprint(adsState)+"}")
		}
	} else if deviceState != 0 {
		return b.Post("/writeControl", "{\"deviceState\":"+fmt.Sprint(deviceState)+"}")
	}
	return b.Post("/writeControl", "{}")
}
