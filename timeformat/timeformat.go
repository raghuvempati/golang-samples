package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//fmt.Println(time.Now().Format("200601021200"))

	out, err := exec.Command("kubectl", "version", "--client").Output()
	//err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	log.Print(string(out))
	fmt.Println(string(out))
}
