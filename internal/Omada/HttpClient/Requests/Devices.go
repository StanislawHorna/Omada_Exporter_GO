package Requests

const PATH_DEVICES_LIST = "/openapi/v1/{omadaID}/sites/{siteID}/devices"

type Device struct {
	MacAddress string `json:"mac"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Model      string `json:"model"`
	IP         string `json:"ip"`
	Uptime     string `json:"uptime"`
	Status     int    `json:"status"`
	CpuUsage   int    `json:"cpuUtil"`
	RamUsage   int    `json:"memUtil"`
	TagName    string `json:"tagName"`
}
