package context

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/redis.v3"
	"github.com/lkiversonlk/aflodit/dao"
)
//Bidder Context including:
//mongo connection
//redis connection

type BidderContext struct {
	DB *dao.Database
	redisClient *redis.Client
}

func NewBidderContext() *BidderContext {
	return &BidderContext{}
}

func (c *BidderContext) ConnectMongo(url, db string) (err error){
	c.DB = dao.NewDatabase()
	return c.DB.Connect(url, db)
}

func (c *BidderContext) ConnectRedisWithDB(url, password string, DB int64){
	client := redis.NewClient(&redis.Options{
		Addr : url,
		Password : password,
		DB : DB,
	})
	c.redisClient = client
}