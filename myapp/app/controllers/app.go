package controllers

import (
	"github.com/revel/revel"
	"go-tools/enet"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Hooot"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

// func (c App) Help(username string, password string) revel.Result {
// 	revel.INFO.Println(username, password)
// 	return c.Render()
// }

func (c App) Help() revel.Result {
	var username, password string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	revel.INFO.Println(username, password)
	return c.Render(username, password)
}

func (c App) Form() revel.Result {

	c.Request.ParseForm()
	UserName := c.Request.Form.Get("username")
	Password := c.Request.Form.Get("password")
	revel.INFO.Printf("UserName :%s,Password :%s\n", UserName, Password)
	return c.RenderText("username :%s ,password :%s", UserName, Password)
}

func (c App) SetSession() revel.Result {
	var username, password string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	c.Session["username"] = username
	c.Session["password"] = password
	return nil
}

func (c App) DeleteSession() revel.Result {
	for k, _ := range c.Session {
		if k == "username" || k == "password" {
			delete(c.Session, k)
		}
	}
	return nil
}

func (c App) ShowSession() revel.Result {
	username := c.Session["username"]
	return c.RenderText("username from cookie is %s", username)
}

func (c App) SetCookie() revel.Result {
	var username, password string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	enet.SetCookie("username", username, c.Response, 7)
	enet.SetCookie("password", password, c.Response, 7)
	return nil
}
