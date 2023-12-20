package main
import (
	"Dictionnary/dict"
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("-------- Calling rest api ---------------------------\n")
	d := dict.Init()
	http.HandleFunc("/get", d.GetHandler) //http:localhost:8030/get?word=your_word
	http.HandleFunc("/list", d.ListHandler) //http:localhost:8030/list
	http.HandleFunc("/remove", d.RemoveHandler) //http:localhost:8030/remove?key=your_key
	http.HandleFunc("/add", d.AddHandler) // curl -X POST -H "Content-Type: application/json" -d "{\"key\": \"VB.Net\", \"value\": \"VB.Net is a programming language\"}" http://localhost:8083/add
	http.HandleFunc("/update", d.UpdateHanlder) // curl -X PUT -H "Content-Type: application/json" -d "{\"key\": \"existing_key\", \"newValue\": \"updated_value\"}" http://localhost:8080/update
	port := 8083
	fmt.Printf("Server is listening on port :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}