package login

//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:login-service

import (
	"back/cmn"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func Enroll(author string) {
	var developer *cmn.ModuleAuthor
	if author != "" {
		var d cmn.ModuleAuthor
		err := json.Unmarshal([]byte(author), &d)
		if err != nil {
			log.Println(err.Error())
			return
		}
		developer = &d
	}

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: login,

		Path: "/api/login",
		Name: "login",

		Developer: developer,
	})

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: signup,

		Path: "/api/signup",
		Name: "signup",

		Developer: developer,
	})
}

var Connection *pgxpool.Pool
func init() {
	var errors error//// urlExample := "postgres://username:password@localhost:5432/database_name"
	configData,err :=cmn.GetViperValue()
	if err!=nil {
		fmt.Println("错误：",err)
	}
	username := configData.DB.LocalDB.User
	password := configData.DB.LocalDB.Password
	ip := configData.DB.LocalDB.IpAddress
	port := configData.DB.LocalDB.Port
	dbname := configData.DB.LocalDB.Dbname
	fmt.Println("username,password,ip,port,dbname",username,password,ip,port,dbname)
	s :=fmt.Sprintf("postgres://%s:%s@%s:%s/%s",username,password,ip,port,dbname)
	config, err := pgxpool.ParseConfig(s)
	//config, err := pgxpool.ParseConfig("postgres://postgres:123456789@postgres:5432/postgres")
	if err != nil {
		fmt.Println("err1:",err)
	}
	Connection, errors = pgxpool.ConnectConfig(context.Background(), config)

	//处理连接失败的情况
	if errors != nil {
		fmt.Println("err2:",errors)
	}
}

func login(ctx context.Context) {
	fmt.Println("登录")
	q := cmn.GetCtxValue(ctx)
	//读取用户名密码
	buf,err :=ioutil.ReadAll(q.R.Body)
	jsonMap :=make(map[string]interface{})
	err = json.Unmarshal(buf,&jsonMap)
	if err!=nil {
		fmt.Println("json unmarshal err ")
		q.W.Write([]byte(fmt.Sprintf(`{"status":300,"err":%s}`,err.Error())))
		return
	}
	username :=jsonMap["username"]
	password :=jsonMap["password"]
	if username==""||password=="" {
		log.Println("用户名或密码为空")
		q.W.Write([]byte(`{"status":300,"err":"用户名或密码为空"}`))
		return
	}

	//查询数据库
	var Uid int
	s:= `select uid from user_login where username = $1 and password = $2 `
	row := Connection.QueryRow(context.Background(), s,username,password)
	err = row.Scan(&Uid)
	if err==pgx.ErrNoRows {
		log.Println("用户不存在，请注册")
		q.W.Write([]byte(`{"status":300,"err":"用户不存在，请注册"}`))
		return
	}

	//登录生成session，存储userid
	session, _ := Store.Get(q.R, "session-name")
	q.Session = session
	session.Values["userid"]=Uid
	err = q.Session.Save(q.R,q.W)
	if err!=nil {
		http.Error(q.W, err.Error(), http.StatusInternalServerError)
		return
	}

	//z.Info("---->" + cmn.FncName())
	//q.Msg.Msg = cmn.FncName()
	//q.Resp()

	q.W.Write([]byte(fmt.Sprintf(`{"status":200,"uid":%d}`,Uid)))
}

func signup(ctx context.Context) {
	fmt.Println("注册")
	q := cmn.GetCtxValue(ctx)
	//读取用户名密码
	buf,err :=ioutil.ReadAll(q.R.Body)
	jsonMap :=make(map[string]interface{})
	err = json.Unmarshal(buf,&jsonMap)

	if err!=nil {
		fmt.Println("json unmarshal err ")
		q.W.Write([]byte(fmt.Sprintf(`{"status":300,"err":%s}`,err.Error())))
		return
	}

	username :=jsonMap["username"]
	password :=jsonMap["password"]
	fmt.Println(username)
	var Uid int
	s:= `select uid from user_login where username = $1 `
	row := Connection.QueryRow(context.Background(), s,username)
	fmt.Println("3")
	err = row.Scan(&Uid)


	if err == pgx.ErrNoRows {
		s = `insert into user_login (username,password) values ($1,$2)`
		_,err = Connection.Exec(context.Background(),s,username,password)
		if err!=nil {
			fmt.Println("Exec: ",err)
		}
		_,err1 := q.W.Write([]byte(`{"status":200,"err":"SignIn successfully"}`))
		if err1!= nil {
			fmt.Println("err1:",err1)
		}
	} else {
		_,err1 := q.W.Write([]byte(fmt.Sprintf(`{"status":300,"err":"username already exist"}`)))
		if err1!= nil {
			fmt.Println("err2:",err1)
		}
	}
}
