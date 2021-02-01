package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//type Status2 struct {
//	//Text string
//}

// review, must to be uppercase, otherwise won't work
type User2 struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	/* perform an HTTP request for the twitter status of user: Googland */
	res, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	/* initialize the structure of the JSON response */
	user := User2{}
	// need to colse body
	defer res.Body.Close()
	/* unmarshal the JSON into our structures */
	temp, _ := ioutil.ReadAll(res.Body)
	body := []byte(temp)
	json.Unmarshal(body, &user)
	fmt.Printf("User: %d, %d, %s, %v\n", user.UserId, user.Id, user.Title, user.Completed)

	res2, _ := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	defer res2.Body.Close()
	user2 := new(User2)
	/* another way to print json, more convinient*/
	json.NewDecoder(res2.Body).Decode(&user2)
	fmt.Printf("User: %d, %d, %s, %v\n", user2.UserId, user2.Id, user2.Title, user2.Completed)

}
