package models

type AccelerationSensor struct {
	AccelerationSensorTotal int
	AccelerationSensorX     int
	AccelerationSensorY     int
	AccelerationSensorZ     int
}

type Accelerator struct {
	AbsoluteThrottlePositionB int
	AcceleratorPedalPositionD int
	AcceleratorPedalPositionE int
	AcceleratorPedalPositionF int
}
