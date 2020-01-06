package dao

import (
	"context"
	"t-mk-opentrace/ext/log-driver/log"
	"t-mk-opentrace/ext/mongo-driver/conn"
	"t-mk-opentrace/ext/mongo-driver/mongo"
	"t-mk-opentrace/interval/material/model"

	"github.com/olivere/elastic"

	"t-mk-opentrace/ext/es-driver/es"

	"github.com/siddontang/go/bson"
)

// ArticleService client
type ArticleService struct {
	CompanyID string
	CompanyDB string
	Limit     int64
}

func (s *ArticleService) articleConn() *mongo.Dao {
	return conn.ArticleMgoConn(s.CompanyID, s.CompanyDB)
}

func (s *ArticleService) articleEs() *elastic.Client {
	return es.Client()
}

// List 显示列表
func (s *ArticleService) List(ctx context.Context, query bson.M) ([]*model.Article, error) {
	opt := &mongo.FindOptions{Limit: &s.Limit, Projection: bson.M{"title": 1, "tags": 1}}
	cursor, err := s.articleConn().Find(ctx, query, opt)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Warn(err)
		}
	}()
	resData := make([]*model.Article, 0, s.Limit)
	for cursor.Next(ctx) {
		v := &model.Article{}
		if err := cursor.Decode(v); err != nil {
			return nil, err
		}
		resData = append(resData, v)
	}
	return resData, nil
}

// Count cnt
func (s *ArticleService) Count(ctx context.Context, query bson.M) (int64, error) {
	return s.articleConn().CountDocuments(ctx, query)
}

// ListEs ListEs
func (s *ArticleService) ListEs(ctx context.Context, companyID string) (int64, error) {
	index := "mk_wx_all_clue_1_current"
	query := elastic.NewBoolQuery()
	resp, err := s.articleEs().Search().Index(index).Type("_doc").Query(query).Do(ctx)
	if err != nil {
		return 0, err
	}
	return resp.Hits.TotalHits, nil
}
