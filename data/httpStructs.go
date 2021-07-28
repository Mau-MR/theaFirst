package data

import "github.com/Mau-MR/theaFirst/data/types"

//Costumers

//Binnacles

//GetBinnacleReq is the request for getting a binnacle
type GetBinnacleReq struct {
	ClientID string `json:"clientID"`
}

//GetBinnacleRes returns the specified binnacle
type GetBinnacleRes struct {
	Binnacle types.Binnacle `json:"binnacle"`
}

type PostBinnacleReq struct {
	CostumerID string `json:"binnacle"`
}
type PostBinnacleCellReq struct {
	BinnacleID string             `json:"binnacleID"`
	Cell       types.BinnacleCell `json:"cell"`
}
