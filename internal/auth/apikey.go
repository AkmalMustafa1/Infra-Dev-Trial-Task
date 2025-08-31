
package auth

import (
    "context"
    "log"
    "os"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
    uri := os.Getenv("MONGODB_URI")
    dbName := os.Getenv("MONGO_DB")
    collName := os.Getenv("MONGO_COLLECTION")

    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database(dbName).Collection(collName)
}

func ValidateAPIKey(apiKey string) bool {
    filter := bson.M{"key": apiKey}
    err := collection.FindOne(context.Background(), filter).Err()
    return err == nil
}
