package downloadfile



//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:downloadfile-service

import (
	"back/cmn"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		Fn: downloadFile,

		Path: "/api/downloadfile",
		Name: "downloadfile",

		Developer: developer,
	})
}

// authMgmt authenticate/authorization management
func downloadFile(ctx context.Context) {
	fmt.Println("文件下载")
	q := cmn.GetCtxValue(ctx)
	//session, _ := login.Store.Get(q.R, "session-name")
	//fmt.Println("q.Session.Values[username]",session.Values["username"])
	//files, err := ioutil.ReadDir("./")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, f := range files {
	//	fmt.Println(f.Name())
	//}
	filepath:=q.R.URL.Query().Get("filepath")
	fmt.Println("filepath:",filepath)
	file,err := os.Open(filepath)
	if err!=nil {
		fmt.Println("Open File Failed")
		fmt.Println(err)
		return
	}

	fileBytes,err :=ioutil.ReadAll(file)
	if err!=nil {
		fmt.Println("Read File Failed",err)
	}

	//s := `{"status":200}`
	//fmt.Println(s)
	q.W.Write(fileBytes)
	fmt.Println("成功")
}
