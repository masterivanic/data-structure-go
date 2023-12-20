package dict
import (
    "fmt"
	"os"
	"log"
	"sort"
	"io/ioutil"
	"encoding/json"
	"strconv"
)


type Dict struct{
	entries map[string]string
	filePath string
}


func removeDuplicates(elements []string) []string {
	encountered := make(map[string]bool)
	result := []string{}

	for _, element := range elements {
		if !encountered[element] {
			encountered[element] = true
			result = append(result, element)
		}
	}
	return result
}


func sortDictByKey(anyDictList []map[string]string) []map[string]string{
	var keys []string
	var sortList []map[string]string
	for _, anyDict := range anyDictList {
		for key := range anyDict {
			keys = append(keys, key)
		}
	}

	// Remove duplicate keys
	keys = removeDuplicates(keys)

	// Sort the keys
	sort.Strings(keys)

	// Create the sorted list of maps
	for _, key := range keys {
		for _, anyDict := range anyDictList {
			if value, ok := anyDict[key]; ok {
				myMap := map[string]string{key: value}
				sortList = append(sortList, myMap)
			}
		}
	}
	return sortList
}

func createFile() (string, bool) {
	filePath := "dict.json"
	var isFileCreate = false
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Erreur lors de la creation du fichier", err)
		} else {
			isFileCreate = true
		}
		defer file.Close()
	}
	return filePath, isFileCreate
}


func Init() Dict {
	filePath, _ := createFile() 
	d := Dict{entries:make(map[string]string), filePath:filePath}
	return d
}

func (anyDict *Dict) Add(key, value string) {
	var dataSlice []map[string]string
	jsonContent, err := ioutil.ReadFile(anyDict.filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	_ = json.Unmarshal(jsonContent, &dataSlice)

	my_map := map[string]string{key: value}
	dataSlice = append(dataSlice, my_map)

	data, err := json.MarshalIndent(dataSlice, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    err = os.WriteFile(anyDict.filePath, data, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func (anyDict *Dict) Get(word string) string {
	var dataMaps []map[string]string
	var wordFind string = "Element not found"
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	_ = json.Unmarshal(jsonContent, &dataMaps)

	for _, aMap := range dataMaps {
		for key, value := range aMap {
			if word == key {
				wordFind = word + " : " + value
				break
			}
		}
	}
	return wordFind
}

func (anyDict *Dict) Remove(key string){
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	var data []map[string]string
	isRemove := false
	_ = json.Unmarshal(jsonContent, &data)
	for i := len(data) - 1; i >= 0; i-- {
		if _, ok := data[i][key]; ok {
			data = append(data[:i], data[i+1:]...)
			isRemove = true
		}
	}

	if isRemove {
		updatedData, _ := json.MarshalIndent(data, "", "    ")
		err := ioutil.WriteFile(anyDict.filePath, updatedData, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("remove successfully")
		}
	} else {
		fmt.Println("cannot remove " + key + " coz dont exist in json file")
	}
}

func (anyDict *Dict) List() {
	var count int = 1
	var dataMaps []map[string]string
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	_ = json.Unmarshal(jsonContent, &dataMaps)
	mapOrder := sortDictByKey(dataMaps)
	for _, aMap := range mapOrder {
		for key, value := range aMap {
			fmt.Println(strconv.Itoa(count) + ") " + key + " : " + value)
			count += 1
		}
	}
}

func (anyDict *Dict) Update(key, newValue string){
	var data []map[string]string
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	isUpdate := false
	_ = json.Unmarshal(jsonContent, &data)
	for _, item := range data {
		if _, ok := item[key]; ok {
			item[key] = newValue
			isUpdate = true
			break
		}
	}

	if isUpdate {
		updatedData, _ := json.MarshalIndent(data, "", "    ")
		err := ioutil.WriteFile(anyDict.filePath, updatedData, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("update successfully")
		} 
	} else {
		fmt.Println("cannot update " + key + " coz dont exist in json file")
	}
	
} 


