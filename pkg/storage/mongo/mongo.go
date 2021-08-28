package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
}

const (
	databaseName   = "news"  // имя БД
	collectionName = "posts" // имя коллекции в БД
)

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрывать ресурсы
	//defer client.Disconnect(context.Background())
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	s := Storage{
		Client: client,
		DB:     client.Database(databaseName),
	}
	return &s, nil
}

// Posts возвращает список всех статей из БД.
func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	//defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var t storage.Post
		err := cur.Decode(&t)
		if err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, cur.Err()
}

// AddPost создает статью.
func (s *Storage) AddPost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePost обновляет статью.
func (s *Storage) UpdatePost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	update := bson.M{
		"$set": p,
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": p.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost удаляет статью по id.
func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": p.ID})
	if err != nil {
		return err
	}
	return nil
}
