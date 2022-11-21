package uploadfile


//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:uploadfile-service

import (
	"back/cmn"
	"back/serve/login"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
type ReplyProto struct {
	//Status, 0: success, others: fault
	Status int `json:"status"`

	//Msg, Action result describe by literal
	Msg string `json:"msg,omitempty"`

	//Data, operand
	Data types.JSONText `json:"data,omitempty"`

	//AllUser, storage the allUserList
	AllUser types.JSONText `json:"alluser,omitempty"`

	//PageCount ,just page count
	PageCount int64 `json:"pageCount,omitempty"`

	// RowCount, just row count
	RowCount int64 `json:"rowCount,omitempty"`

	//API, call target
	API string `json:"API,omitempty"`

	//Method, using http method
	Method string `json:"method,omitempty"`

	//SN, call order
	SN int `json:"SN,omitempty"`

	Qrcode []byte `json:"qrcode,omitempty"`
}
func Resp(w http.ResponseWriter, msg *ReplyProto) {

	if w == nil || msg == nil {
		fmt.Println("call respErr with invalid w/msg")
		return
	}

	//将msg结构体转化为json字符串
	buf, err := json.Marshal(&msg)
	fmt.Println(msg)

	//fmt.Println("buf---------",buf)
	//生成json字符串失败时
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code":-300,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
		return
	}

	w.Write(buf)
	//fmt.Println(string(buf))
}
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
		Fn: uploadFile,

		Path: "/api/uploadfile",
		Name: "uploadfile",

		Developer: developer,
	})

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: getFileList,

		Path: "/api/getFileList",
		Name: "getFileList",

		Developer: developer,
	})
}


func uploadFile(ctx context.Context) {
	fmt.Println("文件上传")
	q := cmn.GetCtxValue(ctx)
	r := q.R
	w := q.W

	//session, _ := login.Store.Get(q.R, "session-name")
	//uid := session.Values["userid"]
	uid := q.R.URL.Query().Get("uid")
	fmt.Println("uid:",uid)

	r.ParseMultipartForm(10 << 20) //maxMemory limit : 10MB
	md5key := r.FormValue("MD5")
	fmt.Println("md5key1:",md5key)

	var md5 interface{}
	s:=`select "MD5key" from file where "MD5key" = $1`
	row := login.Connection.QueryRow(context.Background(),s,md5key)
	err := row.Scan(&md5)
	fmt.Println("md5interface{}",md5)
	fmt.Println("err:",err)
	notExist := err == pgx.ErrNoRows
	fmt.Println("notExist:",notExist)

	var path string
	if notExist {
		//文件不存在，insert
		fmt.Println("文件不存在")
		file,_,err :=r.FormFile("file")
		if err!=nil {
			fmt.Println("Error Retrieving file from form-data")
			fmt.Println(err.Error())
			return
		}
		defer file.Close()

		//测试代码
		tempFile,err := os.Create("test.txt")
		if err!= nil {
			fmt.Println("1",err)
		}
		filename := r.FormValue("filename")
		path = tempFile.Name()


		//创建文件
		//tempFile,err :=ioutil.TempFile("/app/chatAppFile","upload-*.png")
		//if err!=nil {
		//	fmt.Println("1",err)
		//}
		//filename := r.FormValue("filename")
		//path = tempFile.Name()
		//path = strings.TrimPrefix(path, "temp-images\\")
		fmt.Println("path:",path)
		//读文件
		fileBytes,err :=ioutil.ReadAll(file)
		if err!=nil {
			fmt.Println("2",err)
		}
		//复制文件到指定文件
		tempFile.Write(fileBytes)
		//对file表进行insert
		s = `insert into file ("MD5key", "fileName", "filePath",citation) values
            ($1,$2,$3,1)`
		_,err = login.Connection.Exec(context.Background(),s,md5key,filename,path)
		if err!=nil {
			fmt.Println("1 err:",err)
		}

		//同时建立用户和文件的映射
		s = `insert into user_file (uid, "md5") values ($1,$2)`
		_,err = login.Connection.Exec(context.Background(),s,uid,md5key)
		if err!=nil {
			fmt.Println("2 err:",err)
		}
	}else {
		//文件存在，citation+1
		fmt.Println("文件存在")
		s = `update file set citation=citation+1 where "MD5key"= $1 `
		_,err = login.Connection.Exec(context.Background(),s,md5key)
		if err != nil {
			fmt.Println("引用数加一失败：", err)
			return
		}
		//查看当前用户有没有该文件
		fmt.Println("查看当前用户有没有该文件--uid:",uid)
		var Uid int
		s = `select Uid from user_file where md5=$1 and uid=$2`
		row = login.Connection.QueryRow(context.Background(),s,md5key,uid)
		err = row.Scan(&Uid)
		fmt.Println("Uid:",Uid)
		notExist = err == pgx.ErrNoRows
		fmt.Println("err:",err)
		fmt.Println("notExist:",notExist)
		if notExist{
			s = `insert into user_file (uid, "md5") values ($1,$2)`
			_,err := login.Connection.Exec(context.Background(),s,uid,md5key)
			if err!=nil {
				fmt.Println("3 err: ",err)
			}
		}
		s = `select "filePath" from file where "MD5key"=$1`
		row = login.Connection.QueryRow(context.Background(),s,md5key)
		row.Scan(&path)
	}

	data :=fmt.Sprintf(`{"filepath":"%s","status":200}`,path)
	w.Write([]byte(data))

}

func getFileList(ctx context.Context) {
	msg := ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	fmt.Println("getFileList")
	q := cmn.GetCtxValue(ctx)

	session, _ := login.Store.Get(q.R, "session-name")
	uid := session.Values["userid"]
	fmt.Println("uid:",uid)

	s:=`select file."fileName",file."filePath" from user_file,file where file."MD5key"=user_file.md5 and user_file.uid= $1 `
	rows,err :=login.Connection.Query(context.Background(),s,uid)
	if err!=nil {
		log.Println(err)
		return
	}
	var fileList []string
	var filename, filepath string
	for rows.Next() {
		err = rows.Scan(&filename,&filepath)
		if err!=nil {
			log.Println("读取文件信息时出错：",err)
			return
		}
		fileList = append(fileList,fmt.Sprintf(`{"filename":"%s","filepath":"%s"}`,filename,filepath))
	}
	//fmt.Println(fileList,fmt.Sprintf(`[%s]`, strings.Join(fileList, ",")))
	//fmt.Println([]byte(fmt.Sprintf(`[%s]`, strings.Join(fileList, ","))))
	//q.W.Write([]byte(fmt.Sprintf(`[%s]`, strings.Join(fileList, ","))))


	msg.Data=[]byte(fmt.Sprintf(`[%s]`, strings.Join(fileList, ",")))
	Resp(q.W,&msg)

}
