package main

import (
	"Mmx/Request"
	"Mmx/Util"
	"encoding/json"
	"fmt"
)

func ErrHandler(err error) {
	if err != nil {
		fmt.Println("Error occurred")
		panic(err)
	}
}

func main() {
	if Util.Checker.NetOk() {
		fmt.Println("There's no need to login")
		return
	}

	G := Util.Config.Init()

	fmt.Println("Step1: Get local ip returned from srun server.")
	{
		body, err := Request.Get(G.UrlLoginPage, nil)
		ErrHandler(err)
		G.Ip, err = Util.GetIp(body)
		ErrHandler(err)
	}
	fmt.Println("Step2: Get token by resolving challenge result.")
	{
		data, err := Request.Get(G.UrlGetChallengeApi, map[string]string{
			"callback": "jsonp1583251661367",
			"username": G.Form.UserName,
			"ip":       G.Ip,
		})
		ErrHandler(err)
		G.Token, err = Util.GetToken(data)
		ErrHandler(err)
	}
	fmt.Println("Step3: Loggin and resolve response.")
	{
		info, err := json.Marshal(map[string]string{
			"username": G.Form.UserName,
			"password": G.Form.PassWord,
			"ip":       G.Ip,
			"acid":     G.Meta.Acid,
			"enc_ver":  G.Meta.Enc,
		})
		ErrHandler(err)
		G.EncryptedInfo = "{SRBX1}" + Util.Base64(Util.XEncode(string(info), G.Token))
		G.Md5 = Util.Md5(G.Token)
		G.EncryptedMd5 = "{MD5}" + G.Md5

		var chkstr string
		chkstr = G.Token + G.Form.UserName
		chkstr += G.Token + G.Md5
		chkstr += G.Token + G.Meta.Acid
		chkstr += G.Token + G.Ip
		chkstr += G.Token + G.Meta.N
		chkstr += G.Token + G.Meta.VType
		chkstr += G.Token + G.EncryptedInfo
		G.EncryptedChkstr = Util.Sha1(chkstr)

		res, err := Request.Get(G.UrlLoginApi, map[string]string{
			"callback":     "jQuery1124011576657442209481_1602812074032",
			"action":       "login",
			"username":     G.Form.UserName,
			"password":     G.EncryptedMd5,
			"ac_id":        G.Meta.Acid,
			"ip":           G.Ip,
			"info":         G.EncryptedInfo,
			"chksum":       G.EncryptedChkstr,
			"n":            G.Meta.N,
			"type":         G.Meta.VType,
			"os":           "Windows 10",
			"name":         "windows",
			"double_stack": "0",
			"_":            "1602812428675",
		})
		ErrHandler(err)
		G.LoginResult, err = Util.GetResult(res)
		ErrHandler(err)
	}
	fmt.Println("The loggin result is: " + G.LoginResult)
}
