package usermanage
//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:user-service

import (
	"back/cmn"
	"back/serve/login"
	"back/serve/uploadfile"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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
		Fn: getuser,

		Path: "/api/getuser",
		Name: "getuser",

		Developer: developer,
	})
	cmn.AddService(&cmn.ServeEndPoint{
		Fn: changeUser,

		Path: "/api/changeUser",
		Name: "changeUser",

		Developer: developer,
	})
	cmn.AddService(&cmn.ServeEndPoint{
		Fn: deleteUser,

		Path: "/api/deleteUser",
		Name: "deleteUser",

		Developer: developer,
	})
}


func getuser(ctx context.Context) {
	msg := uploadfile.ReplyProto{
		Status: 200,
		Msg:    "success",
	}
	fmt.Println("getFileList")
	q := cmn.GetCtxValue(ctx)

	s:=`select uid,username,password from user_login`
	rows,err :=login.Connection.Query(context.Background(),s)
	if err!=nil {
		log.Println(err)
		return
	}
	var userList []string
	var uid,password int
	var username string
	for rows.Next() {
		err = rows.Scan(&uid,&username,&password)
		if err!=nil {
			log.Println("读取文件信息时出错：",err)
			return
		}
		userList = append(userList,fmt.Sprintf(`{"uid":%d,"username":"%s","password":%d}`,uid,username,password))
	}

	msg.Data=[]byte(fmt.Sprintf(`[%s]`, strings.Join(userList, ",")))
	uploadfile.Resp(q.W,&msg)
}

func changeUser(ctx context.Context){
	fmt.Println("修改用户信息")
	msg := uploadfile.ReplyProto{
		Status: 200,
		Msg:    "success",
	}
	q := cmn.GetCtxValue(ctx)

	buf,err :=ioutil.ReadAll(q.R.Body)
	if err!=nil {
		fmt.Println("ioutil.ReadAll(q.R.Body) err:",err)
		return
	}
	jsonmap := make(map[string]interface{})
	err = json.Unmarshal(buf,&jsonmap)
	if err!=nil {
		fmt.Println("json.Unmarshal err :",err)
		return
	}

	uid:=jsonmap["uid"]
	username:=jsonmap["username"]
	password:=jsonmap["password"]
	fmt.Println("uid",uid)
	fmt.Println("username:",username)
	fmt.Println("password:",password)

	s:=`update user_login set username=$1,password=$2 where uid=$3`
	_,err = login.Connection.Exec(context.Background(),s,username,password,uid)
	if err!=nil {
		log.Println("更新用户基本信息失败")
		return
	}
	uploadfile.Resp(q.W,&msg)
}

func deleteUser(ctx context.Context){
	fmt.Println("getFileList")
	msg := uploadfile.ReplyProto{
		Status: 200,
		Msg:    "success",
	}


	q := cmn.GetCtxValue(ctx)
    uid:=q.R.URL.Query().Get("uid")
	s:=`delete from user_login where uid = $1 `
	_,err :=login.Connection.Exec(context.Background(),s,uid)
	if err!=nil {
		log.Println(err)
		return
	}

	uploadfile.Resp(q.W,&msg)
}

