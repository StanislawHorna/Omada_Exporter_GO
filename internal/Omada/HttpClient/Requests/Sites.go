package Requests

const PATH_SITES = "/openapi/v1/{omadaID}/sites"

type Sites struct {
	SiteID    string `json:"siteId"`
	Name      string `json:"name"`
	Region    string `json:"region"`
	TimeZone  string `json:"timeZone"`
	Scenario  string `json:"scenario"`
	Type      int    `json:"type"`
	SupportES bool   `json:"supportES"`
	SupportL2 bool   `json:"supportL2"`
}
