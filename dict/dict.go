package dict
import (
    "fmt"
)


type Dict struct{
	entries map[string]string
}

func New() Dict {
	d := Dict{entries:make(map[string]string)}
	return d
}

func (any_dict *Dict) Add(key, value string) {
	any_dict.entries[key] = value
}

func (any_dict *Dict) Get(key string) string {
	value_search, is_exist := any_dict.entries[key]
	if is_exist == true {
		return value_search
	}
	return "Element not found"
}

func (any_dict *Dict) Remove(key string){
	delete(any_dict.entries, key)
}

func (any_dict *Dict) List() {
	for key, value := range any_dict.entries {
		fmt.Println(key + "has definition " + value)
	}
}

