package bidder

import (
	"github.com/lkiversonlk/aflodit/model"
	"github.com/lkiversonlk/aflodit/context"
)

type IBidder interface {
	Bid(*model.BidRequest) *model.BidResponse
}

type Bidder struct {
	context *context.BidderContext
	IBidder
}


