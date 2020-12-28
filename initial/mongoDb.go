package initial

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	IsddcMongoDb *mongo.Database
)

func InitMongo() {
	// 设置客户端连接配置
	opts := &options.ClientOptions{}
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	opts.ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "testPersister",
		Username:      "sunbo",
		Password:      "123456"})

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	IsddcMongoDb = client.Database("testPersister")
}
