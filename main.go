package main
import (
	"Dictionnary/dict"
	"fmt"
)

func main()  {
	fmt.Println("-------- before removing ---------------------------\n")
	d := dict.Init()
	d.Add("Java", "Java is a programming language")
	d.List()
	fmt.Println(d.Get("Node") + "\n")
	fmt.Println("-------- after removing ---------------------------\n")
	d.Remove("C")
	fmt.Println(d.Get("Go"))
	d.List()
	d.Update("Java", "This is a bullshit language")
}