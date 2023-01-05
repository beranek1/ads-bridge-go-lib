package adsbridgegolib

import (
	"testing"

	"github.com/beranek1/goadsinterface"
)

func TestInterface(t *testing.T) {
	var _ goadsinterface.AdsLibrary = ADSBridge{}
	var _ goadsinterface.AdsLibrary = (*ADSBridge)(nil)
}
