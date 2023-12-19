package main
import (
    "fmt"
	"sort"
)


type DictType struct {
	key   string
	value string
}

type AnySortFunction func(map[string]string) []DictType

/*
	any_dict (map[string]string): a dictionnary with key value as tring.
	key (string): key to find
	return string
*/
func get_element_from_dict(any_dict map[string]string, key string) string{
	value_search, is_exist := any_dict[key]
	if is_exist == true {
		return value_search
	}
	return "Element not found"
} 

/*
	any_dict (map[string]string): a dictionnary with key value as tring.
	key (string): key to find
	return map[string]string
*/
func remove_element_from_dict(any_dict map[string]string, key string) map[string]string{
	delete(any_dict, key)
	return any_dict
}

func sort_dict_by_key(any_dict map[string]string) []DictType{
	var sort_list []DictType
	for key, value := range any_dict {
		sort_list = append(sort_list, DictType{key, value})
	}

	sort.Slice(sort_list, func(i, j int) bool {
		return sort_list[i].key < sort_list[j].key
	})
	return sort_list
}

// List method
func get_list_as_list(any_dict map[string]string, sort_func AnySortFunction) []DictType{
	return sort_func(any_dict)
}


func add_value_in_dict(any_dict map[string]string, element DictType){
	any_dict[element.key] = element.value
}


func main()  {
	n := map[string]string{"foo": "re", "bar": "e", "program":"go language"}
	value := get_element_from_dict(n, "oo")
	fmt.Println(value)

	value1 := remove_element_from_dict(n, "foo")
	fmt.Println(value1)

	result_list := get_list_as_list(n, sort_dict_by_key)
	fmt.Println(result_list)

	any_value := DictType{"Go", "go is a programming language made by google"}
	add_value_in_dict(n, any_value)
	fmt.Println(n)
}