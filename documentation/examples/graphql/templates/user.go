package templates

type Post struct {
	ID          int32  `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type User struct {
	ID    int32  `json:"ID"`
	Name  string `json:"Name"`
	Posts []Post `json:"Posts"`
}
