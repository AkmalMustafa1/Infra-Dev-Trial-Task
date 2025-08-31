package auth

import (
    "context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func init() {
    uri := os.Getenv("MONGODB_URI")
    db := os.Getenv("MONGO_DB")
    coll := os.Getenv("MONGO_COLLECTION")

    if uri == "" || db == "" || coll == "" {
        log.Fatal("MongoDB environment variables are not set")
    }

    var err error
    client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database(db).Collection(coll)
}

// ValidateAPIKey checks if the API key exists in MongoDB
func ValidateAPIKey(key string) bool {
    if key == "" {
        return false
    }

    var result bson.M
    err := collection.FindOne(context.TODO(), bson.M{"key": key}).Decode(&result)
    if err != nil {
        return false
    }
    return true
}
