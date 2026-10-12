package main
import "net/http"
func FetchUser(id string){ http.Get("http://provider/api/v1/users/42") }
func main(){ FetchUser("42") }
