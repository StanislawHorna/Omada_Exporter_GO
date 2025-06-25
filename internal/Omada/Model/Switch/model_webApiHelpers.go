package Switch

import "omada_exporter_go/internal/Omada/Enum"

const path_WebApiSwitchPort = "{omadaID}/api/v2/sites/{siteID}/switches/{switchMac}/ports"

type webApiSwitchPortStatus struct {
	Port          int             `json:"port"`
	LinkStatus    Enum.LinkStatus `json:"linkStatus"`
	LinkSpeed     Enum.LinkSpeed  `json:"linkSpeed"`
	Duplex        Enum.DuplexMode `json:"duplex"`
	Poe           bool            `json:"poe"`
	Transmit      int64           `json:"tx"`
	Receive       int64           `json:"rx"`
	StpDiscarding bool            `json:"stpDiscarding"`
}

type webApiSwitchPort struct {
	Port         int                    `json:"port"`
	ProfileName  string                 `json:"profileName"`
	Disabled     bool                   `json:"disabled"`
	MaxLinkSpeed Enum.LinkSpeed         `json:"maxSpeed"`
	PortStatus   webApiSwitchPortStatus `json:"portStatus"`
}
