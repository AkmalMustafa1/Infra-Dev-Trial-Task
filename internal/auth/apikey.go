package auth

import (
    "context"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

var coll *mongo.Collection

func init() {
    uri := os.Getenv("MONGODB_URI")
    dbName := os.Getenv("MONGO_DB")
    collName := os.Getenv("MONGO_COLLECTION")

    client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
    client.Connect(context.Background())
    coll = client.Database(dbName).Collection(collName)
}

func ValidateAPIKey(key string) bool {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var result bson.M
    err := coll.FindOne(ctx, bson.M{"key": key, "active": true}).Decode(&result)
    return err == nil
}
