package testing

import (
	"testing"

	"github.com/SonuKanwar/restapi"
)

func TestAddClient(t *testing.T) {
	objClient := []restapi.Client{
		{"1", "Hello World", "Loreum Ipsum", "John"},
		{"2", "Hello", "Loreum", "Loreum"},
		{"1", "World", "Ipsum", "Ipsum"},
	}

	for _, val := range objClient {
		err := restapi.InsertArticles("articles", val)
		if err != nil {
			t.Errorf("Add articles = %v", val)
		}
	}

}

func TestGetClient(t *testing.T) {

	var objClient = []restapi.Client{
		{"1", "", "", ""},
		{"2", "", "", ""},
		{"3", "", "", ""},
		{"all", "", "", ""},
	}

	for _, val := range objClient {
		clientsInfo, err := restapi.GetArticleDetails("articles", val.ArticleID)
		if len(clientsInfo) == 0 || err != nil {
			t.Errorf("Get articles = %v", err.Error())
		}
	}

}
