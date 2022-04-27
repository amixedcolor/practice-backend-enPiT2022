package main

import "fmt"

type User struct{
	Name string
	Email string
	Id string
}

var user = User{}
var user2 = User{}
var user3 = User{}

func createUser(name string, email string) {
	tmpData := User{Name: name, Email: email, Id: "example-user-id"}
	user = tmpData
}

func createUser2(name string, email string) {
	tmpData := User{Name: name, Email: email, Id: "example-user-id2"}
	user2 = tmpData
}

func readUser(id string) (users []User) {
	if user.Id == "example-user-id" {
		users = append(users, user)
		return
	} else if id == "example-user-id2" {
		users = append(users, user2)
		return
	} else if len(id) == 0 {
		users = append(users, user)
		users = append(users, user2)
		return
	} else {
		return nil
	}
}

func updateUser(updateUserData User) {
	if updateUserData.Id == "example-user-id" {
		user = updateUserData
	} else {
		// IDがなかった場合は新規ユーザー登録とみなす
		user3 = updateUserData
		// IDが存在しないのでIDを付与する
		user3.Id = "example-user-id3"
	}
	return
}

func deleteUser() {
	// ハードデリート
	// user = User{}
	// ソフトデリート
	user.Id = "foo"
}

func main() {
	// Create
	createUser("hogehoge", "hoge@example.com")
	createUser2("piyopiyo", "piyo@example.com")
	// Read
	fmt.Println(readUser("example-user-id"))
	// Update
	updateUserData := User{Name: "hogehoge", Email: "huga@example.com", Id: "example-user-id"}
	updateUser(updateUserData)
	fmt.Println(readUser("example-user-id"))
	// Delete
	deleteUser()
	fmt.Println(readUser("example-user-id"))
}