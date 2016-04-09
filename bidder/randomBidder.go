package bidder

import (
	"github.com/lkiversonlk/aflodit/context"
	"github.com/lkiversonlk/aflodit/model"
	"time"
	"fmt"
	"math/rand"
	"github.com/lkiversonlk/aflodit/dao"
	//"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type RandomBidder struct {
	*Bidder
	IBidder

	imgCount int
	counter chan struct{}
}

func NewRandomBidder(c *context.BidderContext) *RandomBidder {
	ret := &RandomBidder{Bidder : &Bidder{context:c}}
	ret.counter = make(chan struct{})

	count, error := ret.Bidder.context.DB.Count("images", nil)
	if error != nil {
		fmt.Println("fail to read img count")
	} else {
		fmt.Println("count is ", count)
		ret.imgCount = count
	}

	//refresh img count every 20 minutes
	ticker := time.NewTicker(20 * time.Minute)
	go func() {
		for{
			select {
			case <- ticker.C:
				count, error := ret.Bidder.context.DB.Count("images", nil)
				if error != nil {
					fmt.Println("fail to read img count")
				} else {
					ret.imgCount = count
				}
			case <- ret.counter:
				ticker.Stop()
			        fmt.Println("stop timer")
				return
			}
		}
	}()

	return ret
}


func (bidder *RandomBidder) Bid(req *model.BidRequest) (res *model.BidResponse) {
	i := rand.Intn(bidder.imgCount)
	result, _, error := bidder.context.DB.Query(dao.ImageCollection, nil, bson.M{"file_id" : 1, "_id" : 0}, i, 1)

	if error != nil {
		return nil
	}

	if len(result) > 0 {
		if file_id, ok := result[0]["file_id"].(string); ok {
			fmt.Println("use file id" + file_id)
			return &model.BidResponse{ImageId: file_id, Result:model.BID}
		} else {
			fmt.Println("fail to retrieve file_id", error)
			return model.NOT_BID_RESPONSE
		}

	}else {
		return model.NOT_BID_RESPONSE
	}

}

func (bidder *RandomBidder) Close() {
	close(bidder.counter)
}