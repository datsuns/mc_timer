package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func runserver() {
	os.Chdir("D:\\Programs\\Minecraft_server")
	cmd := exec.Command("java", "-Xmx1024M", "-Xms1024M", "-jar", "server.jar", "nogui")
	cmdReader, _ := cmd.StdoutPipe()
	writer, _ := cmd.StdinPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	myin := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		done <- true
	}()
	go func() {
		for myin.Scan() {
			fmt.Println("exec say")
			writer.Write([]byte("/say hello\n"))
		}
	}()
	cmd.Start()
	<-done
	err = cmd.Wait()
}

func main() {
	fmt.Println("hello")
	for i, a := range os.Args {
		fmt.Printf("%v) %v\n", i, a)
	}
	s, err := exec.LookPath("java")
	if err != nil {
		panic(err)
	}
	fmt.Printf("java is [%s]\n", s)
}
