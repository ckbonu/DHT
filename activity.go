package dhtxx

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/d2r2/go-dht"

)

const (

	ivPinNumber  = "PinNumber"
	//ivSensorType = "SensorType"
	ovResult = "output"																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																	
)

var log = logger.GetLogger("activity-dhtxx")

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

	PinNumber := context.GetInput(ivPinNumber).(int)
	SensorType := dht.DHT11



	//sensorType := dht.DHT11

	temperature, humidity, retried, err :=
    	dht.ReadDHTxxWithRetry(SensorType, PinNumber, false, 10)

	if err != nil {
		log.Error(err)

		//context.SetOutput(ovResult, err.Error())
		return false, err
	}


	log.Debug(err)

    context.SetOutput(ovResult, temperature)
	context.SetOutput(ovResult, humidity)
    context.SetOutput(ovResult, retried)



	return true, nil
}
