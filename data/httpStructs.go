package data

//Costumers

//Binnacles

//GetBinnacleReq is the request for getting a binnacle
type GetBinnacleReq struct {
	ClientID string `json:"clientID"`
}

//GetBinnacleRes returns the specified binnacle
type GetBinnacleRes struct {
	Binnacle Binnacle `json:"binnacle"`
}

type PostBinnacleReq struct {
	CostumerID string `json:"binnacle"`
}
type PostBinnacleCellReq struct {
	BinnacleID string       `json:"binnacleID"`
	Cell       BinnacleCell `json:"cell"`
}
