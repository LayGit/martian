package main

import (
	"fmt"
)
import "syscall/js"

func encryptToken(_ js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}
	token := args[0].String()
	globalKey := args[1].String()

	// encrypt token by GlobalKey
	md5Key := MD5(globalKey)
	iv := ReverseString(md5Key)
	res, err := EncryptAES_CBC(token, md5Key, iv)
	if err != nil {
		return nil
	}
	return res
}

func decryptToken(_ js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}
	encryptedToken := args[0].String()
	globalKey := args[1].String()

	// decrypt token by GlobalKey
	md5Key := MD5(globalKey)
	iv := ReverseString(md5Key)
	res, err := DecryptAES_CBC(encryptedToken, md5Key, iv)
	if err != nil {
		return nil
	}
	return res
}

func registerCallbacks() {
	js.Global().Set("encryptToken", js.FuncOf(encryptToken))
	js.Global().Set("decryptToken", js.FuncOf(decryptToken))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Hello, Martian!")
	registerCallbacks()
	<-c
}