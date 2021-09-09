// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {

// 	url := "http://localhost:8080/videos"

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("Authorization", "Basic YXp1cmE6MTIz")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println(res)
// 	fmt.Println(string(body))

// }
