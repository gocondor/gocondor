package input

type User struct {
	Name string `json:"name" binding:"exists,alphanum"`
	Age  int    `json:"age" binding:"exists,alphanum,min=18"`
}
