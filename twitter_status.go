package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*这个结构会保存解析后的返回数据。
他们会形成有层级的XML，可以忽略一些无用的数据*/
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {
	// 发起请求查询推特Goodland用户的状态
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	// 初始化XML返回值的结构
	user := User{xml.Name{"", "user"}, Status{""}}
	// 将XML解析为我们的结构
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())

	}
	xml.Unmarshal(body, &user)
	fmt.Println(string(body))
	fmt.Printf("status: %s", user.Status.Text)

}
