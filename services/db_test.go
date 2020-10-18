package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type mockClient struct {
}
type mockDatabase struct {
}
type mockCollection struct {
}
type mockCursor struct {
}
type mockSingleResult struct {
}
type mockInsertOneResult struct {
}
type mockSession struct {
}

func (mc *mockClient) Database(dbName string) Database {
	return nil
}

func (mc *mockClient) StartSession() (mongo.Session, error) {
	return nil, nil
}

func (mc *mockClient) Connect(ctx context.Context) error {
	return nil
}

func (md *mockDatabase) Collection(colName string) Collection {
	return nil
}

func (md *mockDatabase) Client() Client {
	return nil
}

func (mc *mockCursor) Next(ctx context.Context) bool {
	return true
}

func (mc *mockCursor) Decode(v interface{}) error {
	return nil
}

func (mc *mockCursor) Err() error {
	return nil
}

func (mc *mockCursor) Close(ctx context.Context) error {
	return nil
}

func (mc *mockCollection) Find(ctx context.Context, filter interface{}) (Cursor, error) {
	return nil, nil
}

func (mc *mockCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	return nil
}

func (mc *mockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return nil, nil
}

func (mc *mockCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	return 0, nil
}

func (sr *mockSingleResult) Decode(v interface{}) error {
	return nil
}
