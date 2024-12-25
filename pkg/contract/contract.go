package contract

import "mymodule/pkg/entity"

type DataStore interface {
	SaveUser(u entity.User)
	SaveCategory(c entity.Category)
	SaveTask(t entity.Task)
}

type DataLoad interface {
	Load()
}
