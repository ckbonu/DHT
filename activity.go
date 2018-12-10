package dhtxx

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/MichaelS11/go-dht"
)

var log = logger.GetLogger("activity-dhtxx")

const (

	ivPinNumber  = "PinNumber"
	ovResult = "output"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	//sensorType := context.GetInput(ivsensorType).(string)

	PinNumber := context.GetInput(ivPinNumber).(string)
	//output := context.GetInput(ivoutput).(string)

	dht, err := dht.NewDHT(PinNumber, dht.Fahrenheit, "")
	if err != nil {
		context.SetOutput(ovResult, err)
		return false, err
	}

	humidity, temperature, err := dht.Read()
	if err != nil {
		context.SetOutput(ovResult, err)
		return false, err
	}
	// do eval

	// Print temperature and humidity
    context.SetOutput(ovResult, temperature)
	context.SetOutput(ovResult, humidity)

	return true, nil
}
