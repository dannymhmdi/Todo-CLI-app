package filestore

import (
	"mymodule/pkg/entity"
	"mymodule/pkg/jsondesrialize"
	"mymodule/pkg/jsonserialize"
)

var data = entity.Data{}

func New() entity.Data {
	return data
}

type FileStore struct {
	FilePath string
}

func (f FileStore) SaveUser(u entity.User) {
	data.Users = append(data.Users, u)
	jsonserialize.SaveData(data)
}

func (f FileStore) SaveCategory(c entity.Category) {
	data.CategoryStorage = append(data.CategoryStorage, c)
	jsonserialize.SaveData(data)
}

func (f FileStore) SaveTask(t entity.Task) {
	data.Tasks = append(data.Tasks, t)
	jsonserialize.SaveData(data)
}

func (f FileStore) Load() {
	jsondesrialize.LoadData(&data)
}
