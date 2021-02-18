package requests

type TorqueProRequest struct {
	Version   int     `json:"v" bson:"version,omitempty"`
	Session   int64   `json:"session" bson:"session,omitempty" `
	ID        *string `json:"id" bson:"_id,omitempty" `
	Timestamp int64   `json:"time" bson:"timestamp,omitempty" `

	MassAirFlowRate          *float32 `json:"k10" bson:"massAirFlowRate,omitempty"`
	ManifoldThrottlePosition *float32 `json:"k11" bson:"manifoldThrottlePosition,omitempty"`
	AirStatus                *int     `json:"k12" bson:"airStatus,omitempty"`
	FuelTrimBank1Sensor1     *float32 `json:"k14" bson:"fuelTrimBank1Sensor1,omitempty" `
	FuelTrimBank1Sensor2     *float32 `json:"k15" bson:"fuelTrimBank1Sensor2,omitempty" `
	FuelTrimBank1Sensor3     *float32 `json:"k16" bson:"fuelTrimBank1Sensor3,omitempty" `
	FuelTrimBank1Sensor4     *float32 `json:"k17" bson:"fuelTrimBank1Sensor4,omitempty" `
	FuelTrimBank2Sensor1     *float32 `json:"k18" bson:"fuelTrimBank2Sensor1,omitempty" `
	FuelTrimBank2Sensor2     *float32 `json:"k19" bson:"fuelTrimBank2Sensor2,omitempty" `
	FuelTrimBank2Sensor3     *float32 `json:"k1a" bson:"fuelTrimBank2Sensor3,omitempty" `
	FuelTrimBank2Sensor4     *float32 `json:"k1b" bson:"fuelTrimBank2Sensor4,omitempty" `
	RunTimeSinceEngineStart  *int64   `json:"k1f" bson:"runTimeSinceEngineStart,omitempty"`

	DistanceTravelledWithMILOrCelLit        *float32 `json:"k21" bson:"distanceTravelledWithMILOrCelLit,omitempty"`
	FuelRailPressureRelatedToManifoldVacuum *float32 `json:"k22" bson:"fuelRailPressureRelatedToManifoldVacuum,omitempty"`
	FuelRailPressure                        *float32 `json:"k23" bson:"fuelRailPressure,omitempty"`
	O2Sensor1EquivalenceRatio               *float32 `json:"k24" bson:"o2Sensor1EquivalenceRatio,omitempty"`
	O2Sensor2EquivalenceRatio               *float32 `json:"k25" bson:"o2Sensor2EquivalenceRatio,omitempty"`
	O2Sensor3EquivalenceRatio               *float32 `json:"k26" bson:"o2Sensor3EquivalenceRatio,omitempty"`
	O2Sensor4EquivalenceRatio               *float32 `json:"k27" bson:"o2Sensor4EquivalenceRatio,omitempty"`
	O2Sensor5EquivalenceRatio               *float32 `json:"k28" bson:"o2Sensor5EquivalenceRatio,omitempty"`
	O2Sensor6EquivalenceRatio               *float32 `json:"k29" bson:"o2Sensor6EquivalenceRatio,omitempty"`
	O2Sensor7EquivalenceRatio               *float32 `json:"k2a" bson:"o2Sensor7EquivalenceRatio,omitempty"`
	O2Sensor8EquivalenceRatio               *float32 `json:"k2b" bson:"o2Sensor8EquivalenceRatio,omitempty"`
	EGRCommanded                            *float32 `json:"k2c" bson:"eGRCommanded,omitempty"`
	EGRError                                *float32 `json:"k2d" bson:"eGRError,omitempty"`
	EngineECUFuelLevel                      *float32 `json:"k2f" bson:"engineECUFuelLevel,omitempty"`

	FuelStatus                         *float32 `json:"k3" bson:"fuelStatus,omitempty"`
	DistanceTravelledSinceCodesCleared *float32 `json:"k31" bson:"distanceTravelledSinceCodesCleared,omitempty"`
	EvapSystemVaporPressure            *float32 `json:"k32" bson:"evapSystemVaporPressure,omitempty"`
	VehicleBarometricPressure          *float64 `json:"k33" bson:"vehicleBarometricPressure,omitempty"`
	AlternateO2Sensor1EquivalenceRatio *float32 `json:"k34" bson:"alternateO2Sensor1EquivalenceRatio,omitempty"`
	CatalystTemperatureBank1Sensor1    *float32 `json:"k3c" bson:"catalystTemperatureBank1Sensor1,omitempty"`
	CatalystTemperatureBank1Sensor2    *float32 `json:"k3e" bson:"catalystTemperatureBank1Sensor2,omitempty"`
	CatalystTemperatureBank2Sensor1    *float32 `json:"k3d" bson:"catalystTemperatureBank2Sensor1,omitempty"`
	CatalystTemperatureBank2Sensor2    *float32 `json:"k3f" bson:"catalystTemperatureBank2Sensor2,omitempty"`

	EngineLoad                *float32 `json:"k4" bson:"engineLoad,omitempty"`
	VoltageControlModule      *float32 `json:"k42" bson:"voltageControlModule,omitempty"`
	EngineLoadAbsolute        *float32 `json:"k43" bson:"engineLoadAbsolute,omitempty"`
	CommandedEquivalenceRatio *float32 `json:"k44" bson:"commandedEquivalenceRatio,omitempty"`
	RelativeThrottlePosition  *float32 `json:"k45" bson:"relativeThrottlePosition,omitempty"`
	AmbientAirTemperature     *int     `json:"k46" bson:"ambientAirTemperature,omitempty"`
	AbsoluteThrottlePositionB *int     `json:"k47" bson:"absoluteThrottlePositionB,omitempty"`
	AcceleratorPedalPositionD *int     `json:"k49" bson:"acceleratorPedalPositionD,omitempty"`
	AcceleratorPedalPositionE *int     `json:"k4a" bson:"acceleratorPedalPositionE,omitempty"`
	AcceleratorPedalPositionF *int     `json:"k4b" bson:"acceleratorPedalPositionF,omitempty"`

	EngineCoolantTemperature         *float32 `json:"k5" bson:"engineCoolantTemperature,omitempty"`
	EthanolFuelPercentage            *float32 `json:"k52" bson:"ethanolFuelPercentage,omitempty"`
	RelativeAcceleratorPedalPosition *float32 `json:"k5a" bson:"relativeAcceleratorPedalPosition,omitempty"`
	EngineOilTemperature             *float32 `json:"k5c" bson:"engineOilTemperature,omitempty"`

	FuelTrimBank1ShortTerm *float32 `json:"k6" bson:"fuelTrimBank1ShortTerm,omitempty"`

	FuelTrimBank1LongTerm             *float32 `json:"k7" bson:"fuelTrimBank1LongTerm,omitempty"`
	ExhaustGasTemperatureBank1Sensor1 *float32 `json:"k78" bson:"exhaustGasTemperatureBank1Sensor1,omitempty"`
	ExhaustGasTemperatureBank2Sensor1 *float32 `json:"k79" bson:"exhaustGasTemperatureBank2Sensor1,omitempty"`

	FuelTrimBank2ShortTerm *float32 `json:"k8" bson:"fuelTrimBank2ShortTerm,omitempty"`
	FuelTrimBank2LongTerm  *float32 `json:"k9" bson:"fuelTrimBank2LongTerm,omitempty"`

	FuelPressure *float32 `json:"ka" bson:"fuelPressure,omitempty"`

	IntakeManifoldPressure         *float32 `json:"kb" bson:"intakeManifoldPressure,omitempty"`
	TransmissionTemperatureMethod2 *float32 `json:"kb4" bson:"transmissionTemperatureMethod2,omitempty"`

	EngineRPM            *float32 `json:"kc" bson:"engineRPM,omitempty"`
	SpeedOBD             *float32 `json:"kd" bson:"speedOBD,omitempty"`
	TimingAdvance        *float32 `json:"ke" bson:"timingAdvance,omitempty"`
	IntakeAirTemperature *float32 `json:"kf" bson:"intakeAirTemperature,omitempty"`

	TransmissionTemperatureMethod1 *float32 `json:"kfe1805" bson:"transmissionTemperatureMethod1,omitempty"`

	SpeedGPS     *float32 `json:"kff1001" bson:"speedGps,omitempty"`
	GPSLongitude *float32 `json:"kff1005" bson:"gpsLongitude,omitempty"`
	GPSLatitude  *float32 `json:"kff1006" bson:"gpsLatitude,omitempty"`

	GPSAltitude *float32 `json:"kff1010" bson:"gpsAltitude,omitempty"`

	InstantMilesPerGallon                 *float32 `json:"kff1201" bson:"instantMilesPerGallon,omitempty"`
	TurboBoostAndVacuumGauge              *float32 `json:"kff1202" bson:"turboBoostAndVacuumGauge,omitempty"`
	InstantKilometersPerLitre             *float32 `json:"kff1203" bson:"instantKilometersPerLitre,omitempty"`
	TripDistance                          *float32 `json:"kff1204" bson:"tripDistance,omitempty"`
	TripAverageMPG                        *float32 `json:"kff1205" bson:"tripAverageMPG,omitempty"`
	TripAverageKPL                        *float32 `json:"kff1206" bson:"tripAverageKPL,omitempty"`
	InstantLitresPerHundredKilometers     *float32 `json:"kff1207" bson:"instantLitresPerHundredKilometers,omitempty"`
	TripAverageLitresPerHundredKilometers *float32 `json:"kff1208" bson:"tripAverageLitresPerHundredKilometers,omitempty"`
	TripDistanceStoredInProfileVehicle    *float32 `json:"kff120c" bson:"tripDistanceStoredInProfileVehicle,omitempty"`

	O2VoltsBank1Sensor1 *float32 `json:"kff1214" bson:"o2VoltsBank1Sensor1,omitempty"`
	O2VoltsBank1Sensor2 *float32 `json:"kff1215" bson:"o2VoltsBank1Sensor2,omitempty"`
	O2VoltsBank1Sensor3 *float32 `json:"kff1216" bson:"o2VoltsBank1Sensor3,omitempty"`
	O2VoltsBank1Sensor4 *float32 `json:"kff1217" bson:"o2VoltsBank1Sensor4,omitempty"`
	O2VoltsBank2Sensor1 *float32 `json:"kff1218" bson:"o2VoltsBank2Sensor1,omitempty"`
	O2VoltsBank2Sensor2 *float32 `json:"kff1219" bson:"o2VoltsBank2Sensor2,omitempty"`
	O2VoltsBank2Sensor3 *float32 `json:"kff121a" bson:"o2VoltsBank2Sensor3,omitempty"`
	O2VoltsBank2Sensor4 *float32 `json:"kff121b" bson:"o2VoltsBank2Sensor4,omitempty"`

	AccelerationSensorX     *int     `json:"kff1220" bson:"accelerationSensorX,omitempty"`
	AccelerationSensorY     *int     `json:"kff1221" bson:"accelerationSensorY,omitempty"`
	AccelerationSensorZ     *int     `json:"kff1222" bson:"accelerationSensorZ,omitempty"`
	AccelerationSensorTotal *int     `json:"kff1223" bson:"accelerationSensorTotal,omitempty"`
	Torque                  *float32 `json:"kff1225" bson:"torque,omitempty"`
	HorsePowerAtWheels      *float32 `json:"kff1226" bson:"horsePowerAtWheels,omitempty"`
	ZeroToSixtyMphTime      *int64   `json:"kff122d" bson:"zeroToSixtyMphTime,omitempty"`
	ZeroToHundredKphTime    *int64   `json:"kff122e" bson:"zeroToHundredKphTime,omitempty"`
	AQuarterMileTime        *int64   `json:"kff122f" bson:"aQuarterMileTime,omitempty"`

	AnEighthMileTime        *int64   `json:"kff1230" bson:"anEighthMileTime,omitempty"`
	GPSByOBDSpeedDifference *float32 `json:"kff1237" bson:"gpsByOBDSpeedDifference,omitempty"`
	VoltageOBDAdapter       *float32 `json:"kff1238" bson:"voltageOBDAdapter,omitempty"`
	GPSAccuracy             *float32 `json:"kff1239" bson:"gpsAccuracy,omitempty"`
	GPSSatellites           *float32 `json:"kff123a" bson:"gpsSatellites,omitempty"`
	GPSBearing              *float32 `json:"kff123b" bson:"gpsBearing,omitempty"`

	O2Sensor1WideRangeVoltage *float32 `json:"kff1240" bson:"o2Sensor1WideRangeVoltage,omitempty"`
	O2Sensor2WideRangeVoltage *float32 `json:"kff1241" bson:"o2Sensor2WideRangeVoltage,omitempty"`
	O2Sensor3WideRangeVoltage *float32 `json:"kff1242" bson:"o2Sensor3WideRangeVoltage,omitempty"`
	O2Sensor4WideRangeVoltage *float32 `json:"kff1243" bson:"o2Sensor4WideRangeVoltage,omitempty"`
	O2Sensor5WideRangeVoltage *float32 `json:"kff1244" bson:"o2Sensor5WideRangeVoltage,omitempty"`
	O2Sensor6WideRangeVoltage *float32 `json:"kff1245" bson:"o2Sensor6WideRangeVoltage,omitempty"`
	O2Sensor7WideRangeVoltage *float32 `json:"kff1246" bson:"o2Sensor7WideRangeVoltage,omitempty"`
	O2Sensor8WideRangeVoltage *float32 `json:"kff1247" bson:"o2Sensor8WideRangeVoltage,omitempty"`
	MeasuredAirFuelRatio      *float64 `json:"kff1249" bson:"measuredAirFuelRatio,omitempty"`
	TiltX                     *float32 `json:"kff124a" bson:"tiltX,omitempty"`
	TiltY                     *float32 `json:"kff124b" bson:"tiltY,omitempty"`
	TiltZ                     *float32 `json:"kff124c" bson:"tiltZ,omitempty"`
	CommandedAirFuelRatio     *float64 `json:"kff124d" bson:"commandedAirFuelRatio,omitempty"`
	ZeroToTwoHundredKphTime   *int64   `json:"kff124f" bson:"zeroToTwoHundredKphTime,omitempty"`

	InstantaneousCO2            *float32 `json:"kff1257" bson:"instantaneousCO2,omitempty"`
	AverageCO2                  *float32 `json:"kff1258" bson:"averageCO2,omitempty"`
	FuelFlowRateByMinute        *float32 `json:"kff125a" bson:"fuelFlowRateByMinute,omitempty"`
	FuelCostTrip                *float32 `json:"kff125c" bson:"fuelCostTrip,omitempty"`
	FuelFlowRateByHour          *float32 `json:"kff125d" bson:"fuelFlowRateByHour,omitempty"`
	SixtyToHundredTwentyMphTime *int64   `json:"kff125e" bson:"sixtyToHundredTwentyMphTime,omitempty"`

	FortyToSixtyMphTime            *int64   `json:"kff1260" bson:"fortyToSixtyMphTime,omitempty"`
	AverageTripSpeedMovingOnly     *float64 `json:"kff1263" bson:"averageTripSpeedMovingOnly,omitempty"`
	HundredToZeroKphTime           *int64   `json:"kff1264" bson:"hundredToZeroKphTime,omitempty"`
	SixtyToZeroMphTime             *int64   `json:"kff1265" bson:"sixtyToZeroMphTime,omitempty"`
	TripTimeSinceJourneyStart      *uint64  `json:"kff1266" bson:"tripTimeSinceJourneyStart,omitempty"`
	TripTimeWhilstStationary       *uint64  `json:"kff1267" bson:"tripTimeWhilstStationary,omitempty"`
	TripTimeWhilstMoving           *uint64  `json:"kff1268" bson:"tripTimeWhilstMoving,omitempty"`
	CalculatedVolumetricEfficiency *float32 `json:"kff1269" bson:"calculatedVolumetricEfficiency,omitempty"`
	EstimatedDistanceToEmpty       *float32 `json:"kff126a" bson:"estimatedDistanceToEmpty,omitempty"`
	FuelRemaining                  *float32 `json:"kff126b" bson:"fuelRemaining,omitempty"`
	InstantCostPerMileByKm         *float32 `json:"kff126d" bson:"instantCostPerMileByKm,omitempty"`
	TripCostPerMileByKm            *float32 `json:"kff126e" bson:"tripCostPerMileByKm,omitempty"`

	AndroidDeviceBarometer          *float64 `json:"kff1270" bson:"androidDeviceBarometer,omitempty"`
	FuelUsedTrip                    *float32 `json:"kff1271" bson:"fuelUsedTrip,omitempty"`
	AverageTripSpeedStoppedOrMoving *float64 `json:"kff1272" bson:"averageTripSpeedStoppedOrMoving,omitempty"`
	EngineKWAtWheels                *float32 `json:"kff1273" bson:"engineKWAtWheels,omitempty"`
	SixtyToHundredThirtyMphTime     *int64   `json:"kff1276" bson:"sixtyToHundredThirtyMphTime,omitempty"`
	ZerotoThirtyMphTime             *int64   `json:"kff1277" bson:"zerotoThirtyMphTime,omitempty"`
	ZeroToHundredMphTime            *int64   `json:"kff1278" bson:"zeroToHundredMphTime,omitempty"`

	HundredToTwoHundredKphTime *int64 `json:"kff1280" bson:"hundredToTwoHundredKphTime,omitempty"`

	PercentageOfCityDriving    *float32 `json:"kff1296" bson:"percentageOfCityDriving,omitempty"`
	PercentageOfHighwayDriving *float32 `json:"kff1297" bson:"percentageOfHighwayDriving,omitempty"`
	PercentageOfIdleDriving    *float32 `json:"kff1298" bson:"percentageOfIdleDriving,omitempty"`
	AndroidDeviceBatteryLevel  *int     `json:"kff129a" bson:"androidDeviceBatteryLevel,omitempty"`

	LongTermAverageMilesPerGallon             *float32 `json:"kff5201" bson:"longTermAverageMilesPerGallon,omitempty"`
	LongTermAverageKilometersPerLitre         *float32 `json:"kff5202" bson:"longTermAverageKilometersPerLitre,omitempty"`
	LongTermAverageLitresPerHundredKilometers *float32 `json:"kff5203" bson:"longTermAverageLitresPerHundredKilometers,omitempty"`
}
