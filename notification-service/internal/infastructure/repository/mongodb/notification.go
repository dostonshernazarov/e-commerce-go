package mongodb

import (
	"context"
	"ekzamen_5/notification-service/internal/config"
	"ekzamen_5/notification-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

type MongoDB struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	collection  *mongo.Collection
	logger      *slog.Logger
}

func NewMongoDB(cfg *config.Config, logger *slog.Logger) (*MongoDB, error) {
	uri := "mongodb://" + cfg.DB.Host + cfg.DB.Port
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		mongoClient: client,
		db:          client.Database(cfg.DB.Name),
		collection:  client.Database(cfg.DB.Name).Collection(cfg.DB.CollectionName),
		logger:      logger,
	}, nil
}

func (m *MongoDB) SaveNotification(ctx context.Context, notification *entity.Notification) error {
	const op = "MongoDB.Save Notification"
	log := m.logger.With(
		slog.String("method", op))

	log.Info("saving notification")
	defer log.Info("saved notification")
	_, err := m.collection.InsertOne(ctx, notification)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetNotification(ctx context.Context, req *entity.GetNotificationReq) ([]entity.Message, error) {
	const op = "MongoDB.GetNotification"
	log := m.logger.With(slog.String("method", op))
	log.Info("getting notification")
	defer log.Info("finished getting notification")

	var messages []entity.Message

	filter := bson.M{"user_id": req.UserId}
	projection := bson.M{"messages": 1, "_id": 0}

	findOptions := options.Find().SetProjection(projection)
	cursor, err := m.collection.Find(ctx, filter, findOptions)

	if err != nil {
		log.Error("error finding documents", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var notification struct {
			Messages []entity.Message `bson:"messages"`
		}
		err := cursor.Decode(&notification)
		if err != nil {
			log.Error("error decoding document", err.Error())
			return nil, err
		}
		log.Info("decoded notification", slog.Any("notification", notification))
		messages = append(messages, notification.Messages...)
	}

	if err := cursor.Err(); err != nil {
		log.Error("error during cursor iteration", err.Error())
		return nil, err
	}

	return messages, nil
}

func (m *MongoDB) AddNotification(ctx context.Context, userID string, message entity.Message) error {
	const op = "MongoDB.AddNotification"
	log := m.logger.With(slog.String("method", op))
	log.Info("adding notification")

	filter := bson.M{"user_id": userID, "messages": bson.M{"$eq": nil}}

	filter = bson.M{"user_id": userID}
	update := bson.M{
		"$push": bson.M{"messages": message},
	}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("failed to add notification", err)
		return err
	}

	log.Info("notification added successfully")
	return nil
}

func (m *MongoDB) GetOffsetNotification(ctx context.Context, userId string) (int64, error) {
	const op = "MongoDB.GetOffsetNotification"
	log := m.logger.With(slog.String("method", op))
	log.Info("getting offset notification")

	filter := bson.M{"user_id": userId}
	projection := bson.M{"offset": 1}

	var result struct {
		Offset int64 `bson:"offset"`
	}

	err := m.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		log.Error("error getting offset notification", err)
		return 0, err
	}

	return result.Offset, nil
}

func (m *MongoDB) UpdateOffsetNotification(ctx context.Context, userID string, offset int64) error {
	const op = "MongoDB.UpdateOffsetNotification"
	log := m.logger.With(slog.String("method", op))
	log.Info("updating offset notification")
	defer log.Info("updating offset notification")

	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": bson.M{"offset": offset}}
	_, err := m.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
