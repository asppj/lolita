package mongo

import (
	"context"

	"github.com/opentracing/opentracing-go/ext"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dao 数据库操作结构
type Dao struct {
	ClientName ClientName
	DBName     string
	CollName   string
	CompanyID  string
	collection *mongo.Collection
}

// MgoConn 分库连接类型
type MgoConn = func(string, string) *Dao

// NewDao 新建数据库操作对象
func NewDao(clientName ClientName, dbName string, collName string, companyID string) (d *Dao) {
	d = &Dao{}
	d.ClientName = clientName
	d.DBName = dbName
	d.CollName = collName
	d.CompanyID = companyID
	d.collection = collection(clientName, dbName, GetCompanyCollName(companyID, collName))
	return
}

// Collection 获取collection
// 注意：这里没有返回error。是因为不可能出错，理由同Client
func collection(clientName ClientName, dbName string, collName string) *mongo.Collection {
	client := Client(clientName)
	return client.Database(dbName).Collection(collName)
}

// DefaultCtx 默认上下文环境生成
func (d *Dao) DefaultCtx() (context.Context, context.CancelFunc) {
	return DefaultCtxWithOut()
}

// DecodeOne 处理DecodeOne的错误
func (d *Dao) DecodeOne(err error) error {
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err == nil {
		return nil
	}
	return err
}

// HasDuplicatedError 是否是重复id写入错误
func (d *Dao) HasDuplicatedError(err error) bool {
	if err, ok := err.(mongo.WriteException); ok {
		for _, e := range err.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}

// Name https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Name
func (d *Dao) Name() string {
	return d.collection.Name()
}

// Find https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Find
func (d *Dao) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*Cursor, error) {
	sp := d.cmdSpan(ctx, "Find", d.CompanyID)
	defer sp.Finish()
	cur, err := d.collection.Find(ctx, filter, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return cur, err
}

// FindOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.FindOne
func (d *Dao) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *SingleResult {
	sp := d.cmdSpan(ctx, "FindOne", d.CompanyID)
	defer sp.Finish()
	return d.collection.FindOne(ctx, filter, opts...)
}

// InsertMany https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
func (d *Dao) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*InsertManyResult, error) {
	sp := d.cmdSpan(ctx, "InsertMany", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.InsertMany(ctx, documents, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Docs", documents)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// InsertOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne
func (d *Dao) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*InsertOneResult, error) {
	sp := d.cmdSpan(ctx, "InsertOne", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.InsertOne(ctx, document, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Doc", document)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// CountDocuments https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.CountDocuments
func (d *Dao) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	sp := d.cmdSpan(ctx, "CountDocuments", d.CompanyID)
	defer sp.Finish()
	cnt, err := d.collection.CountDocuments(ctx, filter, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return cnt, err
}

// DeleteMany https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteMany
func (d *Dao) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*DeleteResult, error) {
	sp := d.cmdSpan(ctx, "DeleteMany", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// DeleteOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteOne
func (d *Dao) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*DeleteResult, error) {
	sp := d.cmdSpan(ctx, "DeleteOne", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// UpdateMany https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go#L548
func (d *Dao) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*UpdateResult, error) {
	sp := d.cmdSpan(ctx, "UpdateMany", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Update", update)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// UpdateOne https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go#L532
func (d *Dao) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*UpdateResult, error) {
	sp := d.cmdSpan(ctx, "UpdateOne", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Update", update)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// Distinct  https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go
func (d *Dao) Distinct(ctx context.Context, fieldName string, filter interface{},
	opts ...*options.DistinctOptions) ([]interface{}, error) {
	sp := d.cmdSpan(ctx, "Distinct", d.CompanyID)
	defer sp.Finish()
	ret, err := d.collection.Distinct(ctx, fieldName, filter, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.FieldName", fieldName)
		sp.LogKV("db.Cmd.Filter", filter)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return ret, err
}

// FindOneAndUpdate 更新一个
func (d *Dao) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	sp := d.cmdSpan(ctx, "FindOneAndUpdate", d.CompanyID)
	defer sp.Finish()
	return d.collection.FindOneAndUpdate(ctx, filter, update, opts...)
}

// FindOneAndDelete 删除一个
func (d *Dao) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	sp := d.cmdSpan(ctx, "FindOneAndDelete", d.CompanyID)
	defer sp.Finish()
	return d.collection.FindOneAndDelete(ctx, filter, opts...)
}

// UseSession 开启默认事务
func (d *Dao) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	sp := d.cmdSpan(ctx, "UseSession", d.CompanyID)
	defer sp.Finish()
	err := d.collection.Database().Client().UseSession(ctx, fn)
	if err != nil {
		ext.Error.Set(sp, true)
	}
	return err
}

// Aggregate 聚合查询
func (d *Dao) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*Cursor, error) {
	sp := d.cmdSpan(ctx, "UseSession", d.CompanyID)
	defer sp.Finish()
	cur, err := d.collection.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		sp.LogKV("db.Cmd.Pipeline", pipeline)
		sp.LogKV("db.Cmd.Options", opts)
		ext.Error.Set(sp, true)
	}
	return cur, err
}
