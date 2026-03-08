package task

import (
	"errors"
	"gorm.io/gorm"
)

type Task struct {
	running bool
}

func (receiver *Task) Run(db *gorm.DB, f func(db *gorm.DB) error) (err error) {
	if receiver.running {
		err = errors.New("正在运行")
		return
	}
	receiver.running = true
	defer func() {
		receiver.running = false
	}()

	return f(db)
}
