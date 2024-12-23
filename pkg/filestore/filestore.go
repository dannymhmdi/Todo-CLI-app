package filestore

import (
	"encoding/json"
	"fmt"
	"io"
	"mymodule/pkg/entity"
	"os"
)

var data = entity.Data{}

func New() entity.Data {
	return data
}

func LoadData() error {
	file, err := os.Open("db.json")
	if err != nil {
		return fmt.Errorf("error in opening file:%v", err)

	}
	defer file.Close()
	byteVaue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error in reading file content:%v", err)

	}
	err = json.Unmarshal(byteVaue, &data)
	if err != nil {
		return fmt.Errorf("error in parsing json file:%v", err)
	}
	fmt.Printf("data:%+v\n", data)
	return nil
}

func SaveData() error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {

		return fmt.Errorf("error in convert data to json:%v", err)
	}
	file, err := os.Create("db.json")
	if err != nil {

		return fmt.Errorf("error in creating or opening existing file:%v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error in writing data to db.json file:%v", err)
	}
	fmt.Println("data successfully write to db.json")
	return nil
}

type FileStore struct {
	FilePath string
}

func (f FileStore) SaveUser(u entity.User) {
	data.Users = append(data.Users, u)
	SaveData()
}

func (f FileStore) SaveCategory(c entity.Category) {
	data.CategoryStorage = append(data.CategoryStorage, c)
	SaveData()
}

func (f FileStore) SaveTask(t entity.Task) {
	data.Tasks = append(data.Tasks, t)
	SaveData()
}

func (f FileStore) Load() {
	LoadData()
}
