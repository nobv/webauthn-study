package models

type RelyingParty struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetDefaultRelyingParty() (RelyingParty, error) {
	rp := RelyingParty{}

	return rp, nil
}
