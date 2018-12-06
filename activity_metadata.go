package dhtxx

import "github.com/TIBCOSoftware/flogo-lib/core/activity"

var jsonMetadata = `{
	"name": "isteer-dhtxx",
	"version": "0.0.1",
	"type": "flogo:activity",
	"title": "dhtxx",
	"description": "activity description",
	"ref": "github.com/ckbonu/DHT",
	"author": "iSteer-ckbonu",
	"inputs": [
	{
		"name": "PinNumber",
		"type": "string"
	}
	],
	"outputs": [
	{
		"name": "output",
		"type": "string"
	}
	]
}`

func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}