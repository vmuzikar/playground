package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	//http.DefaultTransport.(*http.Transport).ForceAttemptHTTP2 = false // uncomment this to "fix" this

	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "admin")
	form.Add("client_id", "admin-cli")
	form.Add("grant_type", "password")

	//res, _ := http.PostForm("http://localhost:8080/realms/master/protocol/openid-connect/token", form) // http works
	res, _ := http.PostForm("https://localhost:8443/realms/master/protocol/openid-connect/token", form) // https doesn't
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
