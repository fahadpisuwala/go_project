package basics

import (
	"fmt"
	iqra "net/http"
)

func main() {
	fmt.Println("hello from another go")

	resp, err := iqra.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

}
