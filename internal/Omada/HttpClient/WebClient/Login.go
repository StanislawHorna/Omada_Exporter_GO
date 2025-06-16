package WebClient

type Login struct {
	OmadaID string `json:"omadacId"`
	Token   string `json:"token"`
}

type IsLoggedIn struct {
	Login bool `json:"login"`
}
