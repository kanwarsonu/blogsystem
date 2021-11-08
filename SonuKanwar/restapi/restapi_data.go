package restapi

import (
	data "github.com/SonuKanwar/Data"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// InsertArticles :
func InsertArticles(tableName string, info Client) error {

	attribs := make(map[string]*dynamodb.AttributeValue)
	attribs["id"] = &dynamodb.AttributeValue{S: aws.String(info.ArticleID)}
	attribs["title"] = &dynamodb.AttributeValue{S: aws.String(info.Title)}
	attribs["content"] = &dynamodb.AttributeValue{S: aws.String(info.Content)}
	attribs["author"] = &dynamodb.AttributeValue{S: aws.String(info.Author)}

	_, err := data.DynDb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      attribs,
	})

	if err != nil {
		return err
	}
	return nil
}

// GetArticleDetails :
func GetArticleDetails(tableName, articleID string) ([]Client, error) {

	objUserProfile := []Client{}

	if articleID == "all" {
		params := &dynamodb.ScanInput{
			TableName:      aws.String(tableName),
			Select:         aws.String("ALL_ATTRIBUTES"),
			ConsistentRead: aws.Bool(true),
		}
		user, err := data.DynDb.Scan(params)
		if err != nil {
			return objUserProfile, err

		}
		errmapping := dynamodbattribute.UnmarshalListOfMaps(user.Items, &objUserProfile)
		if errmapping != nil {
			return objUserProfile, errmapping
		}
	} else {
		params := &dynamodb.QueryInput{
			TableName: aws.String(tableName),
			KeyConditions: map[string]*dynamodb.Condition{
				"id": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(articleID),
						},
					},
				},
			},
			Select: aws.String("ALL_ATTRIBUTES"),
		}

		user, err := data.DynDb.Query(params)
		if err != nil {
			return objUserProfile, err
		}
		err = dynamodbattribute.UnmarshalListOfMaps(user.Items, &objUserProfile)
		if err != nil {
			return objUserProfile, err
		}
	}

	return objUserProfile, nil
}
