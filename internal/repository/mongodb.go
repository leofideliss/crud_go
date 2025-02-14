package repository

import (
    "crud_go/internal/domain"
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    
)

type mongoRepository struct {
	collection *mongo.Collection
    context context.Context
}

func NewMongoRepository(collection *mongo.Collection) *mongoRepository {
	return &mongoRepository{collection,context.TODO()}
}

func Connect() *mongo.Collection {
    // Find .evn
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    // Get value from .env
    MONGO_URI := os.Getenv("MONGO_URI")

    // Connect to the database.
    clientOption := options.Client().ApplyURI(MONGO_URI)
    client, err := mongo.Connect(context.Background(), clientOption)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection.
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    // Create collection
    collection := client.Database("crud_go").Collection("categories")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to db")

    return collection
}

func (c *mongoRepository) Create(data *domain.Category) bool{
    _, err := c.collection.InsertOne(c.context , data)
	if err != nil{
        return false
    }
    return true
}

func (c *mongoRepository) Delete(id string) bool{
    objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("ID inválido:", err)
	}
    
	filter := bson.D{primitive.E{Key: "_id", Value: objId}}

	_, errDB := c.collection.DeleteOne(c.context, filter)
	if errDB != nil {
		return false
	}

	return true
}

func (c *mongoRepository) Read(id string) (*domain.Category , error){
    
    objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("ID inválido:", err)
	}

    filter := bson.D{primitive.E{Key:"_id" , Value:objId}}

    var category *domain.Category
    errDB := c.collection.FindOne(c.context,filter).Decode(&category)
	if errDB != nil {
		return nil, err
	}

    return category, nil
}


func (c *mongoRepository) Update(data *domain.Category , id string) bool{
    objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("ID inválido:", err)
	}
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", data}}
    result ,_ := c.collection.UpdateOne(c.context,filter,update)
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return true
	}

    return false
}

func (c *mongoRepository) List() interface{}{
   	filter := bson.M{}

	cursor, err := c.collection.Find(c.context, filter)
	if err != nil {
		log.Fatal(err)
	}

    var arrayCategories []interface{}
    for cursor.Next(c.context) {
        var categorie domain.Category
		if err := cursor.Decode(&categorie); err != nil {
			log.Fatal(err)
		}
		arrayCategories = append(arrayCategories, categorie)
	}

    return arrayCategories
}

