package controller

import (
	"BLOG-APP/server/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var collection *mongo.Collection
func SetCollection(c *mongo.Collection){
	collection = c
}

func GetAllPosts(c *fiber.Ctx) error {
	blogs := []models.Blog{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		blog := models.Blog{}
		if err := result.Decode(&blog); err != nil {
			return c.Status(500).JSON(fiber.Map{"msg": err})
		}
		blogs = append(blogs, blog)

	}
	return c.Status(200).JSON(blogs)
}

func GetPost(c *fiber.Ctx) error {
	blog := models.Blog{}
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
func CreatePost(c *fiber.Ctx) error {
	blog := new(models.Blog)
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

func UpdatePost(c *fiber.Ctx) error {
	blog := new(models.Blog)
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
func DeletePost(c *fiber.Ctx) error {
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