package wxinterface

import (
	"crypto/md5"
	"encoding/hex"
	//"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"net/url"
	//"regexp"
	"strings"
	//"time"
)

type WebWeChat struct {
	token   string
	context string
	cookies []*http.Cookie
	isLogin bool
}

type MsgItem struct {
	FakeId   string `json:"fakeid"`
	NickName string `json:"nick_name"`
	DateTime int    `json:"date_time"`
	Content  string `json:"content"`
}

/*
func NewWebWeChat(context string) *WebWeChat {
	w := new(WebWeChat)
	w.context = context
	w.isLogin = w.login()
	return w
}
*/

//为WebWeChat添加login方法,w.login()调用
func (w *WebWeChat) login() bool {
	client := new(http.Client)
	login_url := "https://mp.weixin.qq.com/cgi-bin/login?lang=zh_CN"
	email := "major_360mb@qq.com"
	pwd := "major_360mbadmin"
	/*
		email := "shyashimeifa@163.com"
		pwd := "ij123456n"
	*/
	h := md5.New()
	h.Write([]byte(pwd))
	pwd = hex.EncodeToString(h.Sum(nil))
	post_arg := url.Values{"username": {email}, "pwd": {pwd}, "imgcode": {""}, "f": {"json"}}
	fmt.Print(post_arg)
	/*验证码部分暂时不要
	verifycode_url := "http://mp.weixin.qq.com/cgi-bin/verifycode?username=%s&r=%v"
	req, err := http.NewRequest("GET", fmt.Sprintf(verifycode_url, email, time.Now().Unix()), nil)

	if err != nil {
		fmt.Print("Error:", err)
		return false
	}
	*/
	req, err := http.NewRequest("POST", login_url, strings.NewReader(post_arg.Encode()))
	if err != nil {
		fmt.Print("error:", err)
		return false
	}

	req.Header.Add("Referer", "https://mp.weixin.qq.com/")
	req.Header.Add("Host", "mp.weixin.qq.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36")

	/*
		w.cookies = req.Cookies()
		for i := range req.Cookies(){
			fmt.Print("tt")
			fmt.Printf("arr:", req.Cookies()[i])
		}
	*/

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	return true
}
