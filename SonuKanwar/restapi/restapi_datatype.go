package restapi

type Client struct {
	ArticleID string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}
