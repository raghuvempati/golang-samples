package internal

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	bookCollection = "books"
)

type MongoBookRepository struct {
	Id         BookId //increment Book Id
	mtx        sync.Mutex
	collection *mongo.Collection
}

func NewMongoBookRepository(db *mongo.Database) *MongoBookRepository {
	return &MongoBookRepository{
		collection: db.Collection(bookCollection),
	}
}

func (r *MongoBookRepository) CreateBook(ctx context.Context, book *Book) (BookId, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.Id += 1

	book.Id = r.Id

	_, err := r.collection.InsertOne(ctx, book)
	if err != nil {
		return 0, err
	}

	return r.Id, err
}

func (r *MongoBookRepository) RetrieveBook(ctx context.Context, id BookId) (*Book, error) {
	var book Book

	err := r.collection.FindOne(ctx, map[string]BookId{"id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *MongoBookRepository) UpdateBook(ctx context.Context, book *Book) error {
	_, err := r.collection.UpdateOne(ctx, map[string]BookId{"id": BookId(book.Id)}, map[string]interface{}{"$set": book})
	return err
}

func (r *MongoBookRepository) DeleteBook(ctx context.Context, id BookId) error {
	_, err := r.collection.DeleteOne(ctx, map[string]BookId{"id": id})
	return err
}

func (r *MongoBookRepository) ListBook(ctx context.Context, offset int64, limit int64) ([]*Book, error) {
	var books []*Book
	cursor, err := r.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
