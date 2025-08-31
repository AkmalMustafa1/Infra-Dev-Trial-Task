package auth

import (
    "context"
    "log"
    "os"
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
