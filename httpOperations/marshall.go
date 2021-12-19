package httpOperations

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Age      int
	Name     string
	Lastname string
	Phone    string
}

func Marshalling() {

	data, _ := json.Marshal(&MyObject{Age: 1, Name: "Jair David", Lastname: "Vasquez Martinez", Phone: "7774173332"})
	fmt.Print(string(data))
}

func Unmarshalling() {
	jsonHttp := `{"Age":1,"Name":"Jair David","Lastname":"Vasquez Martinez","Phone":"7774173332"}`
	var data MyObject
	json.Unmarshal([]byte(jsonHttp), &data)

	fmt.Println(data)
}
