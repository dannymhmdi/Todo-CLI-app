package todo

// this part in main func

  // command := flag.String("name","no command","what do you want to do")
	// scanner := bufio.NewScanner(os.Stdin)
	// flag.Parse()
	// for {
	// 	 fmt.Println("please enter your command:")
	// 	 scanner.Scan()
	// 	 *command = scanner.Text()
	// 	 RunCommand(*command)
	//  }



// this part out of main func


// func RunCommand(cmd string) {
// 	fmt.Println("cmd",cmd)
// 	switch cmd {
// 	case "createTask":
// 		CreateTaskHandler()
// 	case  "createCategory":
//         CreateCategoryHandler()
// 	case "registerUser":
// 		RegisterUserHandler()
// 	case "login":
// 		LoginUserHandler()
// 	case "exit":
// 		os.Exit(0)
// 	default:
// 		fmt.Println("No valid command!")

// 	}
// }

// type User struct {
// 	name , email string
// 	id int 
// 	password string
// }

// var users = []User{}


// func CreateTaskHandler () {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var name , category, duedate string
// 	fmt.Println("Enter your task name")
// 	scanner.Scan()
// 	name = scanner.Text()

// 	fmt.Println("Enter task category")
// 	scanner.Scan()
// 	category = scanner.Text()

// 	fmt.Println("Enter the date which task must done")
// 	scanner.Scan()
// 	duedate = scanner.Text()
     
// 	fmt.Printf("Task:%s ,category:%s , duedate: %s\n",name,category,duedate)
// }


// func CreateCategoryHandler () {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var name , color string
// 	fmt.Println("Enter category name")
// 	scanner.Scan()
// 	name = scanner.Text()

// 	fmt.Println("Enter category color")
// 	scanner.Scan()
// 	color = scanner.Text()
//     fmt.Printf("category name:%s , category color:%s\n",name,color)
// }


// func RegisterUserHandler () {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var name , email , password string
// 	fmt.Println("Enter your name")
// 	scanner.Scan()
// 	name = scanner.Text()

// 	fmt.Println("Enter your email")
// 	scanner.Scan()
// 	email = scanner.Text()

// 	fmt.Println("Please set a password")
// 	scanner.Scan()
// 	password = scanner.Text()

// 	newUser:= User{
// 		name: name,
// 		email: email,
// 		id: len(users) + 1,
// 		password: password,
// 	}

// 	users = append(users,newUser)

// 	fmt.Printf("User %s created successfully\n",name)
// 	fmt.Println("users",users)

// }

// func LoginUserHandler () {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var username , _ string
// 	fmt.Println("Enter your email")
// 	scanner.Scan()
// 	username = scanner.Text()

// 	fmt.Println("Enter your password")
// 	scanner.Scan()
// 	_ = scanner.Text()
//     //  findUser := []User{}
// 	// for _,user:= range users {
// 	// 	if user.email == username {
//     //      findUser = append(findUser,user)
// 	// 	 break
// 	// 	}
// 	// }
   
// 	// if len(findUser) == 0 {
// 	// 	fmt.Println("There is no user with this information please register","users",users)
// 	// } else { fmt.Printf("User %s login successfully\n",username)}
// 	fmt.Println("user with email : %s logged in",username)
   
// }

