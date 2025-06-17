package Enum

type OnlineDetection int8

const (
	OnlineDetection_PortDisabled OnlineDetection = -2
	OnlineDetection_LAN_Port     OnlineDetection = -1
	OnlineDetection_No           OnlineDetection = 0
	OnlineDetection_Yes          OnlineDetection = 1
)
