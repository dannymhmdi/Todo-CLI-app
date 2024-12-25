package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mymodule/pkg/contract"
	"mymodule/pkg/entity"
	"mymodule/pkg/filestore"
	"mymodule/pkg/textcolor"
	"os"
	"strconv"
)

var (
	authenticatedUser *entity.User
	data              = entity.Data{}
)

func main() {

	DataFileStore.Load()
	data = filestore.New()
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

var UserFileStore contract.DataStore = filestore.FileStore{
	FilePath: "db.json",
}

var DataFileStore contract.DataLoad = filestore.FileStore{}

func RunCommand(cmd string) {
	DataFileStore.Load()
	data = filestore.New()
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
		CreateTaskHandler(UserFileStore)
	case "listTask":
		ListTask()
	case "createCategory":
		CreateCategoryHandler(UserFileStore)
	case "registerUser":
		RegisterUserHandler(UserFileStore)
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("No valid command!")

	}
}

func CreateTaskHandler(store contract.DataStore) {
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
		fmt.Println(textcolor.Red + "category not found" + textcolor.Reset)
		return
	}

	fmt.Println("Enter the date which task must done")
	scanner.Scan()
	duedate = scanner.Text()
	task := entity.Task{
		Name:       name,
		CategoryId: catId,
		Dudate:     duedate,
		Id:         len(data.Tasks) + 1,
		UserId:     authenticatedUser.Id,
	}

	store.SaveTask(task)
}

func CreateCategoryHandler(store contract.DataStore) {
	fmt.Println("create new category")
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Println("Enter category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Enter category color")
	scanner.Scan()
	color = scanner.Text()
	c := entity.Category{
		Title:  title,
		Color:  color,
		Id:     len(data.CategoryStorage) + 1,
		UserId: authenticatedUser.Id,
	}
	store.SaveCategory(c)
}

func RegisterUserHandler(store contract.DataStore) error {
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error in hashing password:%v", err)
	}

	newUser := entity.User{
		Name:     name,
		Email:    email,
		Id:       len(data.Users) + 1,
		Password: string(hashedPassword),
	}

	store.SaveUser(newUser)
	fmt.Printf(textcolor.Green+"User %s created successfully\n"+textcolor.Reset, name)
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
		log.Fatal(textcolor.Red + "user not found" + textcolor.Reset)
	}

	fmt.Println("Enter your password")
	scanner.Scan()
	password = scanner.Text()
	err := bcrypt.CompareHashAndPassword([]byte(authenticatedUser.Password), []byte(password))
	if err != nil {
		log.Fatal(textcolor.Red + "password is incorrect" + textcolor.Reset)
	}
	fmt.Println(textcolor.Green + "You login successfully" + textcolor.Reset)
}

func ListTask() {
	userId := authenticatedUser.Id
	userTasks := []entity.Task{}
	for _, task := range data.Tasks {
		if task.UserId == userId {
			userTasks = append(userTasks, task)
		}
	}
	fmt.Printf(textcolor.Cyan+"usersTask = %+v\n"+textcolor.Reset, userTasks)
}
