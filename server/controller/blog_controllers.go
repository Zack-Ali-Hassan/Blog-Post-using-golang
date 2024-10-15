package controller

import (
	"BLOG-APP/server/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}

func GetAllPosts(c *gin.Context) {
	blogs := []models.Blog{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching posts"})
		return
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		blog := models.Blog{}
		if err := result.Decode(&blog); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		blogs = append(blogs, blog)

	}
	c.JSON(http.StatusOK, blogs)
}

func GetPost(c *gin.Context) {
	blog := models.Blog{}
	id := c.Params.ByName("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": object_id}
	err = collection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching posts"})
		return
	}
	c.JSON(http.StatusOK, blog)
}
func CreatePost(c *gin.Context) {
	blog := new(models.Blog)
	if err := c.ShouldBindJSON(blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing request body"})
		return
	}
	if blog.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please fill post title.."})
		return
	}
	if blog.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please fill post content.."})
		return
	}
	result, err := collection.InsertOne(context.Background(), blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching post.."})
		return
	}
	blog.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusOK, gin.H{"msg": "Inserted successfully"})
}

func UpdatePost(c *gin.Context) {
	blog := new(models.Blog)
	id := c.Params.ByName("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing request body"})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Updated successfully"})

}
func DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": object_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully"})
}
