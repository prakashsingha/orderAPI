package services

import (
	"context"

	"github.com/prakashsingha/orderAPI/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client Client

func SetClient(c *mongo.Client) {
	client = &mongoClient{c}
}

type Client interface {
	Database(string) Database
	Connect(ctx context.Context) error
	StartSession() (mongo.Session, error)
}

type Database interface {
	Collection(name string) Collection
	Client() Client
}

type Collection interface {
	Find(context.Context, interface{}) (Cursor, error)
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
}

type Cursor interface {
	Next(ctx context.Context) bool
	Decode(v interface{}) error
	Err() error
	Close(context.Context) error
}

type SingleResult interface {
	Decode(v interface{}) error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoCursor struct {
	cur *mongo.Cursor
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoInsertOneResult struct {
	ir *mongo.InsertOneResult
}

type mongoSession struct {
	mongo.Session
}

func NewClient(cnf *config.Config) (*mongo.Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(cnf.URL))

	return c, err
}

func NewDatabase(cnf *config.Config, client Client) Database {
	return client.Database(cnf.DatabaseName)
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.cl.Connect(ctx)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCursor) Next(ctx context.Context) bool {
	return mc.cur.Next(ctx)
}

func (mc *mongoCursor) Decode(v interface{}) error {
	return mc.cur.Decode(v)
}

func (mc *mongoCursor) Err() error {
	return mc.cur.Err()
}

func (mc *mongoCursor) Close(ctx context.Context) error {
	return mc.cur.Close(ctx)
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}) (Cursor, error) {
	cursor, err := mc.coll.Find(ctx, filter)
	return &mongoCursor{cur: cursor}, err
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := mc.coll.InsertOne(ctx, document)
	return result, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := mc.coll.UpdateOne(ctx, filter, update)
	return result, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
