package service

import (
	"fmt"
	db "myFirstApiWithGo/util/db"

	structUser "myFirstApiWithGo/user/struct"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *structUser.User) interface{} {
	mongo := db.Create("User", user)
	fmt.Println(mongo)

	return mongo

}

func GetUser(name string) interface{} {
	mongo := db.FindOne("User", bson.D{{"name", name}})

	return mongo
}

func UpdateUser(user *structUser.User) interface{} {
	find := db.FindOne("User", bson.D{{"name", user.Name}})
	fmt.Println(find)
	if find.Data == nil {

		return find
	}
	mongo := db.Update("User", user)

	return mongo
}
func DeleteUser(user string) interface{} {
	find := db.FindOne("User", bson.D{{"name", user}})
	fmt.Println(find)
	if find.Data == nil {

		return find
	}
	mongo := db.Delete("User", user)

	return mongo
}

// func UpdateUser(user *userSchema.User) string {
// 	return "succesfully"
// }

// func DeleteUser(user *userSchema.User) string {
// 	return "succesfully"
// }
