package adsbridgegolib

import (
	"testing"

	"github.com/beranek1/goadsinterface"
)

func TestInterface(t *testing.T) {
	var _ goadsinterface.AdsLibrary = ADSBridge{}
	var _ goadsinterface.AdsLibrary = (*ADSBridge)(nil)
}

func TestFunctions(t *testing.T) {
	l, err := Connect("http://invalid:1234")
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/version\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetState()
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/state\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetDeviceInfo()
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/device/info\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetSymbol("test")
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/symbol/test\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetSymbolInfo()
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/symbol\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetSymbolValue("test")
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/symbol/test/value\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.GetSymbolList()
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Get \"http://invalid:1234/symbol\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.SetState(goadsinterface.AdsState{})
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Post \"http://invalid:1234/state\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
	_, err = l.SetSymbolValue("test", goadsinterface.AdsData{})
	if err == nil {
		t.Error("No error returned for invalid address")
	}
	if err.Error() != "Post \"http://invalid:1234/symbol/test/value\": dial tcp: lookup invalid: no such host" {
		t.Error("Wrong ADS request.")
	}
}
