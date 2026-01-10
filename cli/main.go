package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Select an option:")
	fmt.Println("1. List contents of /tmp")
	fmt.Println("2. Create a 'test' folder in /tmp")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	switch choice {
	case 1:
		cmd := exec.Command("ls", "/tmp")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing ls command:", err)
			return
		}
		fmt.Println(string(output))
	case 2:
		testDir := "/tmp/test"
		if _, err := os.Stat(testDir); !os.IsNotExist(err) {
			fmt.Println("Error: directory", testDir, "already exists")
			return
		}
		err := os.Mkdir(testDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory 'test' created in /tmp")
	default:
		fmt.Println("Invalid option")
	}
}
