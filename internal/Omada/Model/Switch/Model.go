package Switch

import (
	"omada_exporter_go/internal/Omada/Model/Devices"
)

const PATH_SWITCH = "/openapi/v1/{omadaID}/sites/{siteID}/switches/{switchMac}"

type PoeMode int8

const (
	PoeMode_off PoeMode = 0
	PoeMode_on  PoeMode = 1
)

type PortStatus int8

const (
	PortStatus_disabled PortStatus = 0
	PortStatus_enabled  PortStatus = 1
)

type SwitchPort struct {
	Port                   int        `json:"port"`
	PortName               string     `json:"name"`
	ProfileID              string     `json:"profileId"`
	ProfileName            string     `json:"profileName"`
	ProfileOverrideEnabled bool       `json:"profileOverrideEnabled"`
	PoeMode                PoeMode    `json:"poeMode"`
	LagPort                bool       `json:"lagPort"`
	Status                 PortStatus `json:"status"`
}

type Switch struct {
	DeviceType      Devices.DeviceType `json:"deviceType"`
	Name            string             `json:"name"`
	MacAddress      string             `json:"mac"`
	IP              string             `json:"ip"`
	IPv6List        []string           `json:"ipv6List"`
	Model           string             `json:"model"`
	FirmwareVersion string             `json:"firmwareVersion"`
	Version         string             `json:"version"`
	HwVersion       string             `json:"hwVersion"`
	CpuUsage        int                `json:"cpuUtil"`
	RamUsage        int                `json:"memUtil"`
	Uptime          string             `json:"uptime"`
	PortList        []SwitchPort       `json:"portList"`
}
