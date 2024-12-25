package jsonserialize

import (
	"encoding/json"
	"fmt"
	"mymodule/pkg/entity"
	"os"
)

func SaveData(data entity.Data) error {
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
