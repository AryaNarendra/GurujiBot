package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/strike-official/go-sdk/strike"
)

type Strike_Meta_Request_Structure struct {

	// Bybrisk variable from strike bot
	//
	Bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables"`

	// Our own variable from previous API
	//
	User_session_variables User_session_variables_struct `json: "user_session_variables"`
}

type Bybrisk_session_variables_struct struct {

	// User ID on Bybrisk
	//
	UserId string `json:"userId"`

	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId"`

	// Handler Name for the API chain
	//
	Handler string `json:"handler"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location"`

	// Username of the user
	//
	Username string `json:"username"`

	// Address of the user
	//
	Address string `json:"address"`

	// Phone number of the user
	//
	Phone string `json:"phone"`
}

type GeoLocation_struct struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type User_session_variables_struct struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Role       []string `json:"role"`
	Dept       string   `json:"dept"`
	Sub        string   `json:"sub"`
	Enrollment string   `json:"enrollment"`
	Semester   string   `json:"semester"`
	Branch     string   `json:"branch"`
	Email      string   `json:"email"`
	Phone      int64    `json:"phone"`
}

type AppConfig struct {
	Port  string `json:"port"`
	APIEp string `json:"apiep"`
}

type Teacher struct {
	ID       int64  `json:"teacher_id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Dept     string `json:"dept"`
	Sub      string `json:"sub"`
}

type Student struct {
	Name       string `json:"username"`
	Password   string `json:"password"`
	Enrollment string `json:"enrollment"`
	Semester   string `json:"semester"`
	Branch     string `json:"branch"`
	Email      string `json:"email"`
	Phone      int64  `json:"phone"`
}

type Admin struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

var conf *AppConfig

// This will be your API base link. Below we have used ngrok to make our bot public fast.
//var baseAPI = "http://8e28-2405-201-a407-908e-4c-9fe9-ad8b-43c8.ngrok.io"
var baseAPI = "https://7b2f-27-56-240-216.in.ngrok.io"

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "admin",
		Passwd:               "haCk!567",
		Net:                  "tcp",
		Addr:                 "first-hackathon.cepuilwl2joi.us-east-2.rds.amazonaws.com:3306",
		DBName:               "devengers",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err1 error
	db, err1 = sql.Open("mysql", cfg.FormatDSN())
	if err1 != nil {
		log.Fatal(err1)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	conf = &AppConfig{Port: ":7001", APIEp: ""}
	// Init Routes
	router := gin.Default()
	router.POST("/register", Register)
	router.POST("/registration", Registration)
	router.POST("/login", Login)
	router.POST("/login_as", LoginAs)
	router.POST("/add_user", addUser)

	// Start serving the application
	err := router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}

}

func Register(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	strikeObj := strike.Create("started", baseAPI+"/registration")
	quesObj := strikeObj.Question("role").
		QuestionText().
		SetTextToQuestion("Register As", "desc")

	quesObj.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Teacher", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Student", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Admin", "#3b5375", true)

	ctx.JSON(200, strikeObj)
}

func Login(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	strikeObj := strike.Create("started", baseAPI+"/login_as")
	quesObj := strikeObj.Question("role").
		QuestionText().
		SetTextToQuestion("Login As", "desc")

	quesObj.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Teacher", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Student", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Admin", "#3b5375", true)

	ctx.JSON(200, strikeObj)

}

func Registration(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure

	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}

	role := request.User_session_variables.Role[0]
	fmt.Println("Role " + role)
	strikeObj := strike.Create("started", baseAPI+"/add_user?role="+role)
	if role == "Teacher" {

		quesObj1 := strikeObj.Question("username").
			QuestionText().
			SetTextToQuestion("Please provide Your Username", "desc")

		quesObj1.Answer(true).TextInput("")

		quesObj2 := strikeObj.Question("password").
			QuestionText().
			SetTextToQuestion("Please provide Your Password", "desc")

		quesObj2.Answer(true).TextInput("")

		quesObj3 := strikeObj.Question("dept").
			QuestionText().
			SetTextToQuestion("Name your Department", "desc")

		quesObj3.Answer(true).TextInput("")

		quesObj4 := strikeObj.Question("sub").
			QuestionText().
			SetTextToQuestion("What are you teaching?", "desc")

		quesObj4.Answer(true).TextInput("")

	}
	if role == "Student" {

		quesObj1 := strikeObj.Question("username").
			QuestionText().
			SetTextToQuestion("Please provide Your Username", "desc")

		quesObj1.Answer(true).TextInput("")

		quesObj2 := strikeObj.Question("password").
			QuestionText().
			SetTextToQuestion("Please provide Your Password", "desc")

		quesObj2.Answer(true).TextInput("")

		quesObj3 := strikeObj.Question("enrollment").
			QuestionText().
			SetTextToQuestion("Enter Your Enrollment Number", "desc")
		quesObj3.Answer(true).TextInput("")

		quesObj4 := strikeObj.Question("semester").
			QuestionText().
			SetTextToQuestion("Enter Your Semester", "desc")
		quesObj4.Answer(true).TextInput("")

		quesObj5 := strikeObj.Question("branch").
			QuestionText().
			SetTextToQuestion("Enter Your Branch", "desc")
		quesObj5.Answer(true).TextInput("")

		quesObj6 := strikeObj.Question("email").
			QuestionText().
			SetTextToQuestion("Enter Your Email ID", "desc")
		quesObj6.Answer(true).TextInput("")

		quesObj7 := strikeObj.Question("phone").
			QuestionText().
			SetTextToQuestion("Enter Your Phone Number", "desc")
		quesObj7.Answer(true).TextInput("")

	}
	if role == "Admin" {

		quesObj8 := strikeObj.Question("username").
			QuestionText().
			SetTextToQuestion("Please provide Your Username", "desc")

		quesObj8.Answer(true).TextInput("")

		quesObj9 := strikeObj.Question("password").
			QuestionText().
			SetTextToQuestion("Please provide Your Password", "desc")

		quesObj9.Answer(true).TextInput("")
	}

	ctx.JSON(200, strikeObj)
}

func addUser(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure

	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	role := ctx.Query("role")
	fmt.Println("Role " + role)
	switch role {
	case "Teacher":
		tId, err := addTeacher(Teacher{
			Name:     request.User_session_variables.Username,
			Password: request.User_session_variables.Password,
			Dept:     request.User_session_variables.Dept,
			Sub:      request.User_session_variables.Sub,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added album: %v\n", tId)

	case "Admin":
		tId, err := addAdmin(Admin{
			Name:     request.User_session_variables.Username,
			Password: request.User_session_variables.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added album: %v\n", tId)

	case "Student":
		tId, err := addStudent(Student{
			Name:       request.User_session_variables.Username,
			Password:   request.User_session_variables.Password,
			Enrollment: request.User_session_variables.Enrollment,
			Semester:   request.User_session_variables.Semester,
			Branch:     request.User_session_variables.Branch,
			Email:      request.User_session_variables.Email,
			Phone:      request.User_session_variables.Phone,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added album: %v\n", tId)
	}

}

func LoginAs(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}

}

func addTeacher(tch Teacher) (int64, error) {
	result, err := db.Exec("INSERT INTO Teachers (name, password, department, subject) VALUES (?, ?, ?, ?)", tch.Name, tch.Password, tch.Dept, tch.Sub)
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	return id, nil
}

func addStudent(st Student) (int64, error) {
	result, err := db.Exec("INSERT INTO Students (enrollment_no, name, password, semester, branch, email, phone) VALUES (?, ?, ?, ?, ?, ?, ?)",
		st.Enrollment, st.Name, st.Password, st.Semester, st.Branch, st.Email, st.Phone)
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	return id, nil
}

func addAdmin(adm Admin) (int64, error) {
	result, err := db.Exec("INSERT INTO Admins (name, password) VALUES (?, ?)", adm.Name, adm.Password)
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addATeacher: %v", err)
	}
	return id, nil
}
