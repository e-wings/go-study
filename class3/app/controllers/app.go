package controllers

import (
	"class3/app/models"
	"github.com/revel/revel"
	"net/http"
	"net/url"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Help() revel.Result {
	var UserName, Password string
	c.Params.Bind(&UserName, "username")
	c.Params.Bind(&Password, "password")
	revel.INFO.Printf("UserName :%s,Password :%s\n", UserName, Password)
	return c.Render(UserName, Password)
}

// func (c App) Help(username string, password string) revel.Result {
// 	revel.INFO.Printf("username :%s , password :%s\n", username, password)
// 	return c.Render()
// }

func (c App) Form() revel.Result {

	c.Request.ParseForm()
	UserName := c.Request.Form.Get("username")
	Password := c.Request.Form.Get("password")
	revel.INFO.Printf("UserName :%s,Password :%s\n", UserName, Password)

	return c.RenderText("username :%s ,password :%s", UserName, Password)
}
func (c App) SetSession() revel.Result {
	var UserName, Password string
	c.Params.Bind(&UserName, "username")
	c.Params.Bind(&Password, "password")

	c.Session["username"] = UserName
	c.Session["password"] = Password

	revel.INFO.Printf("username :%s,password :%s\n", UserName, Password)
	return nil
}

func (c App) DeleteSession() revel.Result {
	for k, _ := range c.Session {
		if k == "username" ||
			k == "password" {
			delete(c.Session, k)
		}
	}
	return nil
}
func (c App) ShowUserName() revel.Result {
	UserName := c.Session["username"]
	return c.RenderText(" Current UserName %s", UserName)
}

func (c App) SetCookie() revel.Result {
	var UserName, Password string
	c.Params.Bind(&UserName, "username")
	c.Params.Bind(&Password, "password")
	UserName = models.Base64Encode(UserName)
	Password = models.Base64Encode(Password)
	c.setCookie("username", UserName, c.Response)
	c.setCookie("password", Password, c.Response)
	return nil
}
func (c App) DeleteCookie() revel.Result {
	c.setCookie("password", "", c.Response)
	return nil
}
func (c App) DecodeCookie() revel.Result {
	UserNameCookie, _ := c.Request.Cookie("username")
	PasswordCookie, _ := c.Request.Cookie("password")
	UserName, _ := url.QueryUnescape(UserNameCookie.Value)
	Password, _ := url.QueryUnescape(PasswordCookie.Value)
	UserName = models.Base64Decode(UserName)
	Password = models.Base64Decode(Password)
	return c.RenderText("username :%s,password :%s\n", UserName, Password)
}
func (c App) setCookie(name, value string, dest *revel.Response) {

	ExpireTime := time.Now().AddDate(0, 0, 7)

	var sessionValue string
	sessionValue = value

	sessionData := url.QueryEscape(sessionValue)
	cookie := &http.Cookie{
		Name:     name,
		Value:    sessionData,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  ExpireTime.UTC(),
	}
	http.SetCookie(dest.Out, cookie)
}
func (c App) ShowVersion() revel.Result {
	return nil
}
