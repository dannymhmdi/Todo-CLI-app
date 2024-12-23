package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	LoadData()
	command := flag.String("name", "no command", "what do you want to do")
	scanner := bufio.NewScanner(os.Stdin)
	flag.Parse()

	for {
		RunCommand(*command)
		fmt.Println("please enter your command:")
		scanner.Scan()
		*command = scanner.Text()
	}
}

type User struct {
	Name, Email, Password string
	Id                    int
}

type Task struct {
	Name, Dudate string
	CategoryId   int
	UserId       int
	Id           int
	Isdone       bool
}

type Category struct {
	Title, Color string
	Id, UserId   int
}

type Data struct {
	Users           []User
	Tasks           []Task
	CategoryStorage []Category
}

var (
	categoryStorage   = []Category{}
	tasks             = []Task{}
	users             = []User{}
	authenticatedUser *User
	data              = Data{}
)

// var categoryStorage = []Category{}

// var tasks = []Task{}

// var users = []User{}

// var authenticatedUser *User

// var data = Data{}

func RunCommand(cmd string) {
	if cmd != "registerUser" && cmd != "exit" && authenticatedUser == nil {
		LoginUserHandler()
		if authenticatedUser == nil {
			return
		}
	}
	if cmd == "login" {
		return
	}
	switch cmd {
	case "createTask":
		CreateTaskHandler()
	case "listTask":
		ListTask()
	case "createCategory":
		CreateCategoryHandler()
	case "registerUser":
		RegisterUserHandler()
	case "login":
		LoginUserHandler()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("No valid command!")

	}
}

func CreateTaskHandler() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, duedate string
	fmt.Println("Enter your task name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter task category Id")
	scanner.Scan()
	catId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	catFound := false
	for _, cat := range data.CategoryStorage {
		if cat.Id == catId {
			catFound = true

			break
		}
	}

	if !catFound {
		fmt.Println("category not found")
		return
	}

	fmt.Println("Enter the date which task must done")
	scanner.Scan()
	duedate = scanner.Text()
	task := Task{
		Name:       name,
		CategoryId: catId,
		Dudate:     duedate,
		Id:         len(tasks) + 1,
		UserId:     authenticatedUser.Id,
	}
	data.Tasks = append(data.Tasks, task)
	SaveData()
	fmt.Printf("Tasks:%+v\n", tasks)
}

func CreateCategoryHandler() {
	fmt.Println("create new category")
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Println("Enter category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Enter category color")
	scanner.Scan()
	color = scanner.Text()
	c := Category{
		Title:  title,
		Color:  color,
		Id:     len(categoryStorage) + 1,
		UserId: authenticatedUser.Id,
	}
	data.CategoryStorage = append(data.CategoryStorage, c)
	SaveData()
}

func RegisterUserHandler() error {
	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string
	fmt.Println("Enter your name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter your email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please set a password")
	scanner.Scan()
	password = scanner.Text()
	hashedPasswrod, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error in hashing password:%v", err)
	}

	newUser := User{
		Name:     name,
		Email:    email,
		Id:       len(users) + 1,
		Password: string(hashedPasswrod),
	}

	data.Users = append(data.Users, newUser)
	SaveData()
	fmt.Printf("User %s created successfully\n", name)
	fmt.Println("users", users)
	return nil

}

func LoginUserHandler() {
	scanner := bufio.NewScanner(os.Stdin)
	var username, password string
	fmt.Println("Enter your email")
	scanner.Scan()
	username = scanner.Text()
	for _, user := range data.Users {
		if user.Email == username {
			authenticatedUser = &user
			break
		}
	}
	if authenticatedUser == nil {
		log.Fatal("user not found")
	}

	fmt.Println("Enter your password")
	scanner.Scan()
	password = scanner.Text()
	err := bcrypt.CompareHashAndPassword([]byte(authenticatedUser.Password), []byte(password))
	if err != nil {
		log.Fatal("password is incorrect")
	}
	fmt.Println("You login successfully")
	// for _, user := range data.Users {
	// 	if user.Email == username && user.Password == password {
	// 		authenticatedUser = &user
	// 		break
	// 	}
	//

}

func ListTask() {
	userId := authenticatedUser.Id
	userTasks := []Task{}
	for _, task := range data.Tasks {
		if task.UserId == userId {
			userTasks = append(userTasks, task)
		}
	}
	fmt.Printf("usersTask = %+v\n", userTasks)
}

func LoadData() error {
	file, err := os.Open("db.json")
	if err != nil {
		return fmt.Errorf("error in opening file:%v", err)

	}
	defer file.Close()
	byteVaue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error in reading file content:%v", err)

	}
	err = json.Unmarshal(byteVaue, &data)
	if err != nil {
		return fmt.Errorf("error in parsing json file:%v", err)
	}
	fmt.Printf("data:%+v\n", data)
	return nil
}

func SaveData() error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {

		return fmt.Errorf("error in convert data to json:%v", err)
	}
	file, err := os.Create("db.json")
	if err != nil {

		return fmt.Errorf("error in creating or opening existing file:%v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error in writing data to db.json file:%v", err)
	}
	fmt.Println("data successfully write to db.json")
	return nil
}
