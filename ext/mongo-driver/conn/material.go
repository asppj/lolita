package conn

import (
	"github.com/asppj/lolita/config"
	mongo2 "github.com/asppj/lolita/ext/mongo-driver/mongo"
)

// Service Service
type Service struct {
	Dao       *mongo2.Dao
	CompanyID string
	CompanyDB string
}

// MgoConn 链接类型
type MgoConn = func(string, string) *mongo2.Dao

// ArticleMgoConn mgo connect
func ArticleMgoConn(companyID, companyDB string) *mongo2.Dao {
	collName := config.Get().Mongo.MongoMK.MKDB.Coll.MaterialArticle
	d := mongo2.NewDao(mongo2.MKBiz, companyDB, collName, companyID)
	return d
}
