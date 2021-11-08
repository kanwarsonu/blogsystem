package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	data "github.com/SonuKanwar/data"
	"github.com/julienschmidt/httprouter"
)

// AddClient :
func AddClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var objClient Client

	var jsonStr = []byte(`
	{
		"article_id": 1234,
		"title": "Hello World",
		"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		"author": "John",
	}`)
	err := json.Unmarshal(jsonStr, &objClient)

	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err.Error())
	}

	if r.Body == nil {
		return
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.Client{}
	client.Timeout = time.Duration(120 * time.Second)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	var jsonMessage data.JSONMessageContent
	if resp.StatusCode == 200 {
		jsonMessage.Message = "Success"
	} else {
		jsonMessage.Message = err.Error()
	}

	err = InsertArticles("articles", objClient)
	if err != nil {
		fmt.Println(err.Error())
	}

	var dataInfo data.Data
	dataInfo.ID, _ = strconv.Atoi(objClient.ArticleID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(data.JSONMessage(resp.StatusCode, jsonMessage.Message, dataInfo))

	return
}

// GetClient :
func GetClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	articleID := strings.Split(r.RequestURI, "/")[2]
	if articleID != "all" {
		_, err := strconv.Atoi(articleID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// -------- get article details --------
	clientDetails, err := GetArticleDetails("articles", articleID)
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusCode int
	var message string
	if len(clientDetails) > 0 {
		statusCode = http.StatusOK
		message = "Success"
	} else {
		statusCode = http.StatusNotFound
		message = err.Error()
	}

	// -------------- rendering the output --------------
	final := data.JSONMessageWrappedObj(statusCode, message, clientDetails)
	data.WebResponseJSONObjectNoCache(w, statusCode, final)
	return
}
