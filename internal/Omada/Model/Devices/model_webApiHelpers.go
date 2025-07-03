package Devices

import "omada_exporter_go/internal/Omada/Enum"

const path_WebApiDevicesList = "{omadaID}/api/v2/grid/devices/adopted"

type webApiDevice struct {
	MacAddress    string             `json:"mac"`
	Type          Enum.DeviceType    `json:"type"`
	Version       string             `json:"version"`
	LatestVersion string             `json:"latestVersion"`
	UpgradeNeeded Enum.UpgradeNeeded `json:"needUpgrade"`
	ClientNum     int64              `json:"clientNum"`
	SiteName      string             `json:"siteName"`
}
