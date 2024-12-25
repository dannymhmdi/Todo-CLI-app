package jsondesrialize

import (
	"encoding/json"
	"fmt"
	"io"
	"mymodule/pkg/constant"
	"mymodule/pkg/entity"
	"os"
)

func LoadData(data *entity.Data) error {
	file, err := os.Open(constant.StorePath)
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
	//fmt.Printf("data:%+v\n", data)
	return nil
}
