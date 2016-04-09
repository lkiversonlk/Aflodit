package model

type BidResponse struct{
	Result int
	ImageId string
}

const (
	BID = iota
	NO_BID = iota
)

var (
	NOT_BID_RESPONSE = &BidResponse{Result: NO_BID, ImageId:""}
)