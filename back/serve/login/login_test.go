package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type ResBody struct {
	Status int `json:"status"`
}

func TestLogin(t *testing.T) {
	requestBody :=[]byte(`{"username":"欧文","password":"123456789"}`)
	url := "http://localhost:9090/api/login"
	response, err := http.Post(url,"application/json",bytes.NewBuffer(requestBody))
	if err!=nil {
		t.Error("Post err:",err)
	}

	body,err := ioutil.ReadAll(response.Body)
	if err!=nil {
		t.Error("ioutil.ReadAll(response.Body) err:",err)
	}
	var data ResBody
	json.Unmarshal(body,&data)
	fmt.Println("body:",data)
	if data.Status == 200 {
		t.Log("Test Pass")
	} else {
		t.Error("username or password is wrong")
	}
}
