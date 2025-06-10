package Requests

import (
	utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

const PATH_SITES = "/openapi/v1/{omadaID}/sites"

type SitesResponse struct {
	SiteID    string `json:"siteId"`
	Name      string `json:"name"`
	Region    string `json:"region"`
	TimeZone  string `json:"timeZone"`
	Scenario  string `json:"scenario"`
	Type      int    `json:"type"`
	SupportES bool   `json:"supportES"`
	SupportL2 bool   `json:"supportL2"`
}

func (SitesResponse) Path(placeholders map[string]string) string {
	return utils.FillInEndpointPlaceholders(PATH_SITES, placeholders)
}
