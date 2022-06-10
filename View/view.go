package View

import (
	"fmt"
)

type viewTits struct {
	info interface{}
}

func ShowInfo(info interface{}) {
	var z = viewTits{info}
	fmt.Println(z.info)
}

/*
func GetCommand() {
	var command string
	fmt.Scanln(&command)
	Controller.Controller(command)
}
*/
/*
func View(show string) {
	switch show {
	case "start":
		fmt.Println("Ведите команду")
	}
}
*/
