package dhtxx

import (

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/d2r2/go-dht"
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

	Pin := context.GetInput(ivPinNumber).(int)
	//output := context.GetInput(ivoutput).(string)

	// do eval
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, Pin, true, 10)
	if err != nil {
		log.Debug(err)
	}
	// Print temperature and humidity
    context.SetOutput(ovResult, temperature)
	context.SetOutput(ovResult, humidity)
	context.SetOutput(ovResult, retried)

	// fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		// temperature, humidity, retried)

	return true, nil
}
