package comments

type Comments struct {
	PostId    int    `json:"postId"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string   `json:"body"`
}

