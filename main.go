package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Blog struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
	Date    time.Time          `json:"date"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading environment variables : ", err)
		}
	}

	MONGO_URI := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	fmt.Println("MONGODB connection successfully")
	collection = client.Database("golang_blog_db").Collection("Blogs")
	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:5173/",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))
	app.Get("/api/post/", getAllPosts)
	app.Get("/api/post/:id", getPost)
	app.Post("/api/post/", createPost)
	app.Patch("/api/post/:id", updatePost)
	app.Delete("/api/post/:id", deletePost)
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4411"
	}
	log.Fatal(app.Listen(":" + port))
}

func getAllPosts(c *fiber.Ctx) error {
	blogs := []Blog{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		blog := Blog{}
		if err := result.Decode(&blog); err != nil {
			return c.Status(500).JSON(fiber.Map{"msg": err})
		}
		blogs = append(blogs, blog)

	}
	return c.Status(200).JSON(blogs)
}

func getPost(c *fiber.Ctx) error {
	blog := Blog{}
	id := c.Params("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "Invalid Id, please try again?"})
	}
	filter := bson.M{"_id": object_id}
	err = collection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(blog)
}
func createPost(c *fiber.Ctx) error {
	blog := new(Blog)
	if err := c.BodyParser(blog); err != nil {
		return err
	}
	if blog.Title == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Please fill post title.."})
	}
	if blog.Content == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Please fill post content.."})
	}
	result, err := collection.InsertOne(context.Background(), blog)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": err})
	}
	blog.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(fiber.Map{"msg": "Inserted successfully"})
}

func updatePost(c *fiber.Ctx) error {
	blog := new(Blog)
	id := c.Params("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "Invalid Id, please try again?"})
	}
	if err := c.BodyParser(&blog); err != nil {
		return err
	}
	update_fields := bson.M{}
	if blog.Title != "" {
		update_fields["title"] = blog.Title
	}
	if blog.Content != "" {
		update_fields["content"] = blog.Content
	}
	filter := bson.M{"_id": object_id}
	update := bson.M{"$set": update_fields}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(fiber.Map{"msg": "Updated successfully"})
}
func deletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "Invalid Id, please try again?"})
	}
	filter := bson.M{"_id": object_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(fiber.Map{"msg": "Deleted successfully"})
}
