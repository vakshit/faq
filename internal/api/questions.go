package api

import (
	// "fmt"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vakshit/faq/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAllQuestions(ctx *fiber.Ctx) error {
	Questions := database.MongoClient.Questions.Collection("Questions")
	ques := []database.Question{}
	filter := bson.D{{}}
	cursor, err := Questions.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Unable to get questions: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint((err)))
	}
	if err = cursor.All(context.TODO(), &ques); err != nil {
		log.Printf("Unable to get questions: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint((err)))
	}
	return ctx.Status(200).JSON(ques)
}

func getApprovedQuestions(ctx *fiber.Ctx) error {
	Questions := database.MongoClient.Questions.Collection("Questions")
	ques := []database.Question{}
	filter := bson.D{{Key: "approved", Value: true}}
	cursor, err := Questions.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Unable to get questions: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint((err)))
	}
	if err = cursor.All(context.TODO(), &ques); err != nil {
		log.Printf("Unable to get questions: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint((err)))
	}
	return ctx.Status(200).JSON(ques)
}

func postQuestion(ctx *fiber.Ctx) error {
	Questions := database.MongoClient.Questions.Collection("Questions")
	// parsing body into question struct
	q := new(database.Question)
	err := ctx.BodyParser(q)
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	if err != nil {
		log.Printf("Unable to parse question: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprint(err))
	}
	q.Question = strings.TrimSpace(q.Question)
	if len(q.Question) == 0 || q.Approved {
		log.Println("Question empty or already approved")
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprint(err))
	}
	// inserting to DB
	_, err = Questions.InsertOne(context.TODO(), q)
	if err != nil {
		log.Printf("Unable to add questions: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint(err))
	}
	return ctx.Status(200).JSON("true")
}

func answerQuestion(ctx *fiber.Ctx) error {
	Questions := database.MongoClient.Questions.Collection("Questions")
	// parsing into question
	q := new(database.Question)
	err := ctx.BodyParser(q)
	if err != nil {
		log.Printf("Unable to extract question: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprint(err))
	}
	// checking if question is empty
	q.Question = strings.TrimSpace(q.Question)
	// println(q)
	if len(q.Question) == 0 || !q.Approved {
		log.Println("Question empty or not approved")
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprint(err))
	}
	// Updating DB
	filter := bson.D{{Key: "_id", Value: q.ID}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "answer", Value: q.Answer}}}}
	update := bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "answer", Value: q.Answer}, bson.E{Key: "updatedAt", Value: time.Now()}, bson.E{Key: "approved", Value: q.Approved}}}}

	_, err = Questions.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Unable to answer question: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint(err))
	}
	return ctx.Status(200).JSON("true")
}

func deleteQuestion(ctx *fiber.Ctx) error {
	Questions := database.MongoClient.Questions.Collection("Questions")
	// parsing body
	payload := struct {
		ID primitive.ObjectID `bson:"_id,omitempty" json:"id" csv:"-"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprint(err))
	}

	// deleting question
	filter := bson.D{{Key: "_id", Value: payload.ID}}
	_, err := Questions.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Unable to delete question: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint(err))
	}
	return ctx.Status(200).JSON("true")
}
