package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	UserName string `json:"name"`
	UserAge  string `json:"age"`
}

type AppConfig struct {
	Port  string `json:"port"`
	APIEp string `json:"apiep"`
}

var conf *AppConfig

// This will be your API base link. Below we have used ngrok to make our bot public fast.
//var baseAPI = "http://8e28-2405-201-a407-908e-4c-9fe9-ad8b-43c8.ngrok.io"
var baseAPI = "https://7b2f-27-56-240-216.in.ngrok.io"

func main() {
	conf = &AppConfig{Port: ":7001", APIEp: ""}
	// Init Routes
	router := gin.Default()
	router.POST("/", GettingStarted)
	router.POST("/respondBack", RespondBack)

	// Start serving the application
	err := router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}

}

func GettingStarted(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	// Core Logic
	strikeObj := strike.Create("getting_started", baseAPI+"/respondBack")

	// First Question: Whats your name?
	quesObj1 := strikeObj.Question("name").
		QuestionText().
		SetTextToQuestion("Hi! What is your name? Here to help, dude !!", "")
	// Prompt the user to give his answer as a text.
	quesObj1.Answer(true).TextInput("")

	// Second Question: Whats your age?
	quesObj2 := strikeObj.Question("age").
		QuestionText().
		SetTextToQuestion("What is you age", "desc")
	// Prompt the user to give his answer as a number.
	quesObj2.Answer(true).NumberInput("")

	ctx.JSON(200, strikeObj)
}

func RespondBack(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}

	name := request.User_session_variables.UserName
	age := request.User_session_variables.UserAge

	// Core Logic (Not giving any return API as this is the last response to the User.)
	strikeObj := strike.Create("getting_started", "")

	// Respond back
	strikeObj.Question("").
		QuestionText().
		SetTextToQuestion("Hi! "+name+" You are the choosen few. Congratulation on such a feat at the age of "+age+".", "")

	ctx.JSON(200, strikeObj)
}
