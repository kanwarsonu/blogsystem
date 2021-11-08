package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynDb is the pointer to the DynamoDB database.
var DynDb *dynamodb.DynamoDB

// InitDynamoDB - Inits a DynamoDB session and returns it
func InitDynamoDB() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Failed to create session,", err)
		return nil, errors.New("dynamoDb error : " + err.Error())
	}
	fmt.Printf("Connecting to DynamoDB (localhost)....")
	creds := credentials.NewStaticCredentials("admin", "admin", "")
	sess.Config.Region = aws.String("us-east-1")
	sess.Config.Endpoint = aws.String("http://127.0.0.1:8000")
	sess.Config.Credentials = creds

	DynDb = dynamodb.New(sess)
	fmt.Println("Connected to region:", *sess.Config.Region)
	return DynDb, nil
}

// JSONMessageContent :
type JSONMessageContent struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       Data   `json:"data"`
}

// Data
type Data struct {
	ID int `json:"id"`
}

// JSONWrappedContent :
type JSONWrappedContent struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Content    interface{} `json:"content"`
}

// JSONMessage returns a preformatted ISG JSON with the response code and message.
func JSONMessage(code int, msg string, data Data) []byte {
	jsonString := JSONMessageContent{
		StatusCode: code,
		Message:    msg,
		Data:       data,
	}

	//result, err := json.Marshal(jsonString)
	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return result
}

// JSONMessageWrappedObj returns an encoded JSON of the object provided.
func JSONMessageWrappedObj(code int, message string, obj interface{}) []byte {
	jsonString := JSONWrappedContent{
		StatusCode: code,
		Message:    message,
		Content:    obj,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// WebResponseJSONObjectNoCache is a wrapper function that returns an already prepared JSON object as web response with caching disabled.
func WebResponseJSONObjectNoCache(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Vary", "Accept-Encoding")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	finalobj := obj.([]byte)
	w.Header().Set("Content-Length", strconv.Itoa(len(finalobj)))
	w.WriteHeader(code)
	w.Write(finalobj)
}
