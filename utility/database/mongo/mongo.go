package mongo

import (
	goContext "context"

	"github.com/cjexp/base/utility/configuration"
	"github.com/cjtoolkit/ctx/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMain(context ctx.Context) *mongo.Client {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		mongoDial := configuration.GetConfig(context).Database.MongoDial

		c, err := mongo.NewClient(options.Client().ApplyURI(mongoDial))
		if err != nil {
			return c, err
		}
		err = c.Connect(goContext.Background())
		return c, err
	}).(*mongo.Client)
}

func GetMainDatabase(context ctx.Context) *mongo.Database {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		mongoDb := configuration.GetConfig(context).Database.MongoDb

		return GetMain(context).Database(mongoDb), nil
	}).(*mongo.Database)
}

func GetMainGridFs(context ctx.Context) *gridfs.Bucket {
	type mainMongoGridFs struct{}
	return context.Persist(mainMongoGridFs{}, func() (interface{}, error) {
		return gridfs.NewBucket(GetMainDatabase(context))
	}).(*gridfs.Bucket)
}

func MustObjectIdFromHex(s string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return id
}
