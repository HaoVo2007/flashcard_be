package main

import (
	"context"
	"flashcard/config"
	"flashcard/internal/topics"
	"flashcard/internal/user"
	"flashcard/internal/words"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: Error loading .env file: %v", err)
		} else {
			log.Println("Successfully loaded .env file")
		}
	} else {
		log.Println("No .env file found, using environment variables")
	}

	cfg := config.LoadConfig()

	mongoClient, err := connectToMongoDB(cfg.MongoURI)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()

r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		log.Printf("Request from origin: '%s'", origin)
		log.Printf("Request method: %s", c.Request.Method)
		log.Printf("Request path: %s", c.Request.URL.Path)
		
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With, Cache-Control")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			log.Printf("Handling OPTIONS preflight request for path: %s", c.Request.URL.Path)
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})
	userCollections := mongoClient.Database("flashcard").Collection("users")
	userRepository := user.NewUserRepository(userCollections)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	wordsCollections := mongoClient.Database("flashcard").Collection("words")
	wordsRepository := words.NewWordRepository(wordsCollections)
	wordsService := words.NewWordService(wordsRepository)
	wordsHandler := words.NewWordHandler(wordsService)

	topicCollections := mongoClient.Database("flashcard").Collection("topics")
	topicRepository := topics.NewTopicRepository(topicCollections)
	topicService := topics.NewTopicService(topicRepository, wordsService)
	topicHandler := topics.NewTopicHandler(topicService)

	words.RegisterRoutes(r, wordsHandler)
	topics.RegisterRoutes(r, topicHandler)
	user.RegisterRoutes(r, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Failed to connect to MongoDB")
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Failed to ping to MongoDB")
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}