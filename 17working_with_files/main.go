package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

const FileName string = "./user.csv"

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter full name: ")
	//fullName, _ := reader.ReadString('\n')
	//fmt.Print("Enter user name: ")
	//userName, _ := reader.ReadString('\n')
	//fmt.Print("Enter email: ")
	//email, _ := reader.ReadString('\n')
	//fmt.Print("Enter password: ")
	//password, _ := reader.ReadString('\n')
	//
	//// Create new user
	//newUser := User{
	//	FullName: strings.TrimSpace(fullName),
	//	UserName: strings.TrimSpace(userName),
	//	Email:    strings.TrimSpace(email),
	//	Password: []byte(password),
	//	Status:   "active",
	//}
	////newUser.createPassword(password)
	//// open a file for writing
	//file, err := os.Create(FileName)
	//if err != nil {
	//	panic(err)
	//}
	//// close file on exit in self invoking function
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(file)
	//
	//// append to file
	//_, err = file.WriteString(newUser.FullName + "," + newUser.UserName + "," + newUser.Email + "," + strings.TrimSpace(string(newUser.Password)) + "," + newUser.Status + "\n")
	//if err != nil {
	//	return
	//}
	//fmt.Println("User created successfully")
	readFile(FileName)
}

type User struct {
	FullName string
	UserName string
	Email    string
	Password []byte
	Status   string
}

func (user *User) createPassword(password string) User {
	user.Password = md5.New().Sum([]byte(password))
	return *user
}

func readFile(FileName string) {
	bytes, err := ioutil.ReadFile(FileName)
	if err != nil {
		return
	}
	fmt.Println(bytes)         // Raw bytes of the file
	fmt.Println(string(bytes)) // String version of the file

	//file, err := os.Open(FileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(file)
	//
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
}
