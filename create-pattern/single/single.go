package single

import (
	"fmt"
	"sync"
)

/*
单例模式
@data: 2023.8.25
@author: lzw
*/

var once sync.Once

type Single struct {
	Name string
	Age  int
}

var SingleInstance *Single

func GetSingleInstance() *Single {
	if SingleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			SingleInstance = &Single{
				Name: "lzw",
				Age:  18,
			}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return SingleInstance
}
