package dict
import (
    "fmt"
	"os"
	"sort"
	"io/ioutil"
	"encoding/json"
	"net/http"
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

func (anyDict *Dict) Add(key, value string) string {
	var message string = ""
	var dataSlice []map[string]string
	jsonContent, err := ioutil.ReadFile(anyDict.filePath)
	if err != nil {
		message = "Error reading file: " +err.Error()
	}

	_ = json.Unmarshal(jsonContent, &dataSlice)

	my_map := map[string]string{key: value}
	dataSlice = append(dataSlice, my_map)

	data, _ := json.MarshalIndent(dataSlice, "", "  ")
    err = os.WriteFile(anyDict.filePath, data, 0644)
    if err != nil {
        message = err.Error()
    } else {
		message = "Add successfully"
	}
	return message
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

func (anyDict *Dict) Remove(key string) string {
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	var data []map[string]string
	isRemove := false
	var message string = ""
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
			message = "Something wrong:" + err.Error() 
		} else {
			message = "Delete successfully" 
		}
	} else {
		message = "Cannot remove " + key + " coz dont exist in json file"
	}
	return message
}

func (anyDict *Dict) List() []map[string]string {
	var dataMaps []map[string]string
	var resultList []map[string]string
	jsonContent, _ := ioutil.ReadFile(anyDict.filePath)
	_ = json.Unmarshal(jsonContent, &dataMaps)
	mapOrder := sortDictByKey(dataMaps)
	for _, aMap := range mapOrder {
		resultMap := make(map[string]string)
		for key, value := range aMap {
			resultMap[key] = value
		}
		resultList = append(resultList, resultMap)
	}
	return resultList
}

func (anyDict *Dict) Update(key, newValue string) string{
	var message string = ""
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
			message = err.Error()
		} else {
			message = "update successfully"
		} 
	} else {
		message = "cannot update " + key + " coz dont exist in json file"
	}
	return message
} 


func (anyDict *Dict) GetHandler(writer http.ResponseWriter, req *http.Request) {
	word := req.URL.Query().Get("word")
	result := anyDict.Get(word)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"result": result})
}

func (anyDict *Dict) ListHandler(writer http.ResponseWriter, req *http.Request){
	result := anyDict.List()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(result)
}

func (anyDict *Dict) RemoveHandler(writer http.ResponseWriter, req *http.Request){
	key := req.URL.Query().Get("key")
	message := anyDict.Remove(key)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"result": message})
}

func (anyDict *Dict) AddHandler(writer http.ResponseWriter, req *http.Request){
		var requestData map[string]string
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(writer, "Failed to read request body", http.StatusBadRequest)
			return
		}
	
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}
	
		// Extract key and value from the request data
		key, keyExists := requestData["key"]
		value, valueExists := requestData["value"]
	
		if !keyExists || !valueExists {
			http.Error(writer, "Key and value are required", http.StatusBadRequest)
			return
		}
	
		message := anyDict.Add(key, value)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(map[string]string{"message": message})
}

func (anyDict *Dict) UpdateHanlder(writer http.ResponseWriter, req *http.Request)  {
	var requestData map[string]string

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	key, keyExists := requestData["key"]
	newValue, valueExists := requestData["newValue"]

	if !keyExists || !valueExists {
		http.Error(writer, "Key and newValue are required", http.StatusBadRequest)
		return
	}

	message := anyDict.Update(key, newValue)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": message})
}





