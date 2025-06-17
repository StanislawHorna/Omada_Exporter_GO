package Enum

type DeviceType string

const (
	DeviceType_Switch      DeviceType = "switch"
	DeviceType_AccessPoint DeviceType = "ap"
	DeviceType_Gateway     DeviceType = "gateway"
)
