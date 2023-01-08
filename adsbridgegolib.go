package adsbridgegolib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/beranek1/goadsinterface"
)

// Struct used for managing ADSBridge with its address addr
type ADSBridge struct {
	addr string
}

// Creates instance of type ADSBridge with given address and checks connection, returns ADSBridge and error if connection fails.
func Connect(addr string) (*ADSBridge, error) {
	b := &ADSBridge{addr: addr}
	_, err := b.GetVersion()
	if err != nil {
		return b, err
	}
	return b, nil
}

// Reads and converts JSON response of ADSBridge
func processResponse(r io.Reader, data any) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

// Performs GET request with ADSBridge, converts and returns result
func (b ADSBridge) Get(path string, data any) error {
	var url = b.addr + path
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return processResponse(resp.Body, data)
}

// Performs POST request with ADSBridge, converts and returns result
func (b ADSBridge) Post(path string, jsonStr string, data any) error {
	var url = b.addr + path
	resp, err := http.Post(url, "text/json", bytes.NewBufferString(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return processResponse(resp.Body, data)
}

func (b ADSBridge) GetVersion() (goadsinterface.AdsVersion, error) {
	var data goadsinterface.AdsVersion
	err := b.Get("/version", &data)
	return data, err
}

func (b ADSBridge) GetState() (goadsinterface.AdsState, error) {
	var data goadsinterface.AdsState
	err := b.Get("/state", &data)
	return data, err
}

func (b ADSBridge) GetDeviceInfo() (goadsinterface.AdsDeviceInfo, error) {
	var data goadsinterface.AdsDeviceInfo
	err := b.Get("/device/info", &data)
	return data, err
}

func (b ADSBridge) GetSymbol(name string) (goadsinterface.AdsSymbol, error) {
	var data goadsinterface.AdsSymbol
	err := b.Get("/symbol/"+name, &data)
	return data, err
}

func (b ADSBridge) GetSymbolInfo() (goadsinterface.AdsSymbolInfo, error) {
	var data goadsinterface.AdsSymbolInfo
	err := b.Get("/symbol", &data)
	return data, err
}

func (b ADSBridge) GetSymbolValue(name string) (goadsinterface.AdsData, error) {
	var data goadsinterface.AdsData
	err := b.Get("/symbol/"+name+"/value", &data)
	return data, err
}

func (b ADSBridge) GetSymbolList() (goadsinterface.AdsSymbolList, error) {
	var data goadsinterface.AdsSymbolInfo
	err := b.Get("/symbol", &data)
	if err != nil {
		return nil, err
	}
	symbols := make([]string, len(data))
	i := 0
	for k := range data {
		symbols[i] = k
		i++
	}
	return symbols, err
}

func (b ADSBridge) SetState(state goadsinterface.AdsState) (goadsinterface.AdsState, error) {
	jsonStr, err := json.Marshal(state)
	if err != nil {
		return state, err
	}
	var data goadsinterface.AdsState
	err = b.Post("/state", string(jsonStr), &data)
	return data, err
}

func (b ADSBridge) SetSymbolValue(name string, value goadsinterface.AdsData) (goadsinterface.AdsData, error) {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		return value, err
	}
	var data goadsinterface.AdsData
	err = b.Post("/symbol/"+name+"/value", string(jsonStr), &data)
	return data, err
}
