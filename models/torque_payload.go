package models

type TorquePayload struct {
	Version   int
	Session   string
	Id        string
	Timestamp int64

	AirStatus                       float64
	AmbientAirTemperature           float32
	AndroidDeviceBatteryLevel       int
	MovingOnlyAverageTripSpeed      float64
	StoppedOrMovingAverageTripSpeed float64
	AndroidDeviceBarometer          float64
}
