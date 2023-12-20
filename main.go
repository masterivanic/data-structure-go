package main
import (
    "fmt"
	"Dictionnary/dict"
)

func main()  {
	d := dict.New()
	fmt.Println(d)
	d.Add("Go", "go is a programming language made by google")
	fmt.Println(d.Get("Go"))
	fmt.Println(d.Get("Poc"))
	d.List()
}