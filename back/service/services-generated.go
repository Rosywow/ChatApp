//Package services generated by service-enroll-generate.go;
// !DO NOT EDIT THIS FILE!
package service

import (
	
	"back/serve/chat" //chat ,  auth, , XUnion@GMail.com
	"back/serve/downloadfile" //downloadfile ,  auth, , XUnion@GMail.com
	"back/serve/login" //login ,  auth, , XUnion@GMail.com
	"back/serve/uploadfile" //uploadfile ,  auth, , XUnion@GMail.com
	"back/serve/usermanage" //user ,  auth, , XUnion@GMail.com
	"back/serve/xxx" //xxx ,  xxx, , XUnion@GMail.com
)

//Enroll will be called from serve cmd
func Enroll(){
  
  chat.Enroll(`{"name":"auth","email":"XUnion@GMail.com"}`)
  downloadfile.Enroll(`{"name":"auth","email":"XUnion@GMail.com"}`)
  login.Enroll(`{"name":"auth","email":"XUnion@GMail.com"}`)
  uploadfile.Enroll(`{"name":"auth","email":"XUnion@GMail.com"}`)
  usermanage.Enroll(`{"name":"auth","email":"XUnion@GMail.com"}`)
  xxx.Enroll(`{"name":"xxx","email":"XUnion@GMail.com"}`)
}