package Requests

const PATH_API_INFO = "/api/info"

type ApiInfoResponse struct {
	ControllerVersion string `json:"controllerVer"`
	ApiVersion        string `json:"apiVer"`
	Configured        bool   `json:"configured"`
	Type              int    `json:"type"`
	SupportApp        bool   `json:"supportApp"`
	OmadaID           string `json:"omadacId"`
	RegisteredRoot    bool   `json:"registeredRoot"`
	OmadaCategory     string `json:"omadacCategory"`
	MspMode           bool   `json:"mspMode"`
	OmadaCloudUrl     string `json:"omadaCloudUrl"`
}

func (ApiInfoResponse) Path(placeholders map[string]string) string {
	return PATH_API_INFO
}
