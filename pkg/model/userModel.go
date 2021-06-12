package model

type User struct {
	Id             uint   `json:"id"`
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Language       string `json:"language"`
	CreateTime     string `json:"create_time"`
	LastAccessTime string `json:"last_access_time"`
	Status         int8   `json:"status"`
}

func (b *User) TableName() string {
	return "user"
}
