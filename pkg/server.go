package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactUser struct {
	Login    string
	Password string
	//Success       bool
	//	StorageAccess string //пустота
	//Title     string
	FirstName string
	LastName  string
}

func Handler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})

	c.JSON(http.StatusOK, gin.H{
		"ups": "tra ly ly",
	})
}

func GetHandlerForm(c *gin.Context) {
	c.HTML(200, "form.html", gin.H{})

}

func PostHandlerForm(c *gin.Context) {
	c.HTML(200, "form.html", gin.H{})

	data := ContactUser{
		Login:     c.PostForm("Login"),
		Password:  c.PostForm("Password"),
		FirstName: c.PostForm("FirstName"),
		LastName:  c.PostForm("LastName"),
	}

	c.JSON(http.StatusOK, gin.H{
		"Login":     data.Login,
		"Password":  data.Password,
		"FirstName": data.FirstName,
		"LastName":  data.LastName,
	})
	//c.JSON(200, gin.H{
	//	"Login":     data.Login,
	//	"Password":  data.Password,
	//	"FirstName": data.FirstName,
	//	"LastName":  data.LastName,
	//})
	fmt.Printf("\n\n %s %s %s %s \n", data.Login, data.Password, data.FirstName, data.LastName)

}
