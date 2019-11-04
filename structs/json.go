package main

import ("fmt"; "encoding/json")

type Car struct{
	Make string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
}


func main() {
	car := `{"make":"Toyota", "model":"Camry_LE", "year":2011}`
	var jsonToObject *Car
	json.Unmarshal([]byte(car), &jsonToObject)
	fmt.Println(*jsonToObject)

	//another way to unmarshall json to car struct
	var jsonToObject2 Car
	json.Unmarshal([]byte(car), &jsonToObject2)
	fmt.Println(jsonToObject2)

	//convert struct to json
	car2,_ := json.Marshal(jsonToObject)
	fmt.Println(string(car2))

}
