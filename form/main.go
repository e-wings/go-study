package main

import (
	"fmt"
	"go-tools/eform"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("register.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		//=========== 必填的判断 ===========
		if len(r.Form.Get("username")) == 0 {
			fmt.Println("用户名是必填的")
		}
		/*
			//如果是map的值，则必须通过这种方法来判断
			if len(r.Form["username"][0]) == 0 {
				log.Fatal("用户名是必填的")
			}
		*/

		//=========== 数字 ===========
		ageInt, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("年龄必须是数字")
		} else {
			if ageInt < 18 {
				fmt.Println("你年龄太小")
			}
		}

		//=========== 中文 ===========
		if r.Form.Get("name-cn") != "" {
			if m, _ := eform.IsFormat(r.Form.Get("name-cn"), "cn"); !m {
				fmt.Println("中文姓名必须由中文字符构成")
			}
		}

		//=========== 英文（允许出现空格） ===========
		if r.Form.Get("name-en") != "" {
			if m, _ := eform.IsFormat(r.Form.Get("name-en"), "en"); !m {
				fmt.Println("英文姓名必须由英文字母构成")
			}
		}

		//=========== 电子邮件地址 ===========
		if r.Form.Get("email") != "" {
			if m, _ := eform.IsFormat(r.Form.Get("email"), "email"); !m {
				fmt.Println("电子邮件地址不规范")
			}
		}

		//=========== 手机号码 ===========
		if r.Form.Get("mobile") != "" {
			if m, _ := eform.IsFormat(r.Form.Get("mobile"), "mobile"); !m {
				fmt.Println("手机号码不规范")
			}
		}

		//=========== 身份证号 ===========
		if r.Form.Get("id-card") != "" {
			if m, _ := eform.IsFormat(r.Form.Get("id-card"), "idCard"); !m {
				fmt.Println("身份证号不正确")
			}
		}

		//=========== 单选按钮 ===========
		if r.Form.Get("gender") != "" {
			if m, _ := eform.InRange(r.Form.Get("gender"), []string{"0", "1"}); !m {
				fmt.Println("性别错误")
			}
		}

		//=========== 多选按钮 ===========
		if r.Form.Get("interest") != "" {
			if m, _ := eform.RangeInRange(r.Form["interest"], []string{"football", "basketball", "tennis"}); !m {
				fmt.Println("喜爱的运动错误")
			}
		}

		//=========== 下拉菜单 ===========
		if r.Form.Get("fruit") != "" {
			if m, _ := eform.InRange(r.Form.Get("fruit"), []string{"0", "1", "2"}); !m {
				fmt.Println("最喜欢的水果超出范围")
			}
		}

	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/register", register)   //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
