package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"caleox-spaceforum/models"
	"caleox-spaceforum/utils"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	data, _ := utils.GetBin()
	users := data["users"].([]interface{})

	for _, u := range users {
		user := u.(map[string]interface{})
		if user["username"] == username {
			color.Red("Username already exists!")
			return
		}
	}

	newUser := models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hash),
		Joined:   time.Now().Format("2006-01-02"),
	}

	users = append(users, newUser)
	data["users"] = users
	utils.UpdateBin(data)
	color.Green("Signup successful! You can now login.")
}

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	data, _ := utils.GetBin()
	users := data["users"].([]interface{})

	for _, u := range users {
		user := u.(map[string]interface{})
		if user["username"] == username {
			err := bcrypt.CompareHashAndPassword([]byte(user["password"].(string)), []byte(password))
			if err != nil {
				color.Red("Wrong password!")
				return
			}
			utils.SaveSession(user["id"].(string), username)
			color.Green("Logged in as %s", username)
			return
		}
	}
	color.Red("Username not found!")
}
