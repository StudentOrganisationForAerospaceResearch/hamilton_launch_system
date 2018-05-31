package main

type Vec3 struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
	Z int32 `json:"z"`
}

type AccelGyroMagnetismMsg struct {
	Type    string `json:"type"`
	Accel   Vec3   `json:"accel"`
	Gyro    Vec3   `json:"gyro"`
	Magneto Vec3   `json:"magneto"`
}

type BarometerMsg struct {
	Type        string `json:"type"`
	Pressure    int32  `json:"pressure"`
	Temperature int32  `json:"temperature"`
}

type GpsMsg struct {
	Type          string `json:"type"`
	Altitude      int32  `json:"altitude"`
	EpochTimeMsec int32  `json:"epochTimeMsec"`
	Latitude      int32  `json:"latitude"`
	Longitude     int32  `json:"longitude"`
}

type OxidizerTankPressureMsg struct {
	Type     string `json:"type"`
	Pressure int32  `json:"pressure"`
}

type CombustionChamberPressureMsg struct {
	Type     string `json:"type"`
	Pressure int32  `json:"pressure"`
}

type FlightPhaseMsg struct {
	Type        string `json:"type"`
	FlightPhase int8   `json:"flightPhase"`
}
