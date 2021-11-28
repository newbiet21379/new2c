package dao

import (
	"context"
	"github.com/newbiet21379/new2c/connectionhelper"
	"github.com/newbiet21379/new2c/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateVideo - Insert a new document in the collection.
func CreateVideo(video entity.Video) error {
	// Get MongoDB connection using connectionhelper
	client,err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	// Create a handle to the respective collection in the database
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)
	video.ID = primitive.NewObjectID()
	// Perform InsertOne operation & validate against the error.
	_ , err = collection.InsertOne(context.TODO(),video)
	if err != nil{
		return err
	}
	// Return success without any error
	return nil
}

// CreateMany - Insert Multiple documents at once in the collection
func CreateMany( list []entity.Video ) error{
	// Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i,v := range list {
		insertableList[i] = v
	}
	// Get MongoDB connection using the connectionhelper
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	// Create a handle to the respective collection in the database
	collection := client.Database(connectionhelper.CONNECTIONSTRING).Collection(connectionhelper.VIDEOS)
	// Perform InsertMany operation & validate against the error
	_ , err = collection.InsertMany(context.TODO(),insertableList)
	if err != nil {
		return err
	}
	//Return success without any error
	return nil
}

// GetOneVideo - Get One videos By email and video title
func GetOneVideo(author string,title string) (entity.Video, error) {
	result := entity.Video{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "author.email", Value: author},primitive.E{Key:"title",Value: title}}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

//GetAllVideos - Get All videos for collection
func GetAllVideos() ([]entity.Video, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	var videos []entity.Video
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return videos, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return videos, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := entity.Video{}
		err := cur.Decode(&t)
		if err != nil {
			return videos, err
		}
		videos = append(videos, t)
	}
	// once exhausted, close the cursor
	err = cur.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return videos, mongo.ErrNoDocuments
	}
	return videos, nil
}

// UpdateURL - Update URL by id
func UpdateURL(id string, newUrl string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "url", Value: newUrl},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)

	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//DeleteOne - Get All videos for collection
func DeleteOne(id string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)
	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//DeleteAll - Delete All videos for collection
func DeleteAll() error {
	//Define filter query for fetching specific document from collection
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.VIDEOS)
	//Perform DeleteMany operation & validate against the error.
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}