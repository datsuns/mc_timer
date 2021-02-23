package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/rivo/tview"
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

func maybe_chdir_to_executable_file_path() {
	for i, a := range os.Args {
		fmt.Printf("%v) %v\n", i, a)
	}
}

func notify_error_when_java_not_found() {
	s, err := exec.LookPath("java")
	if err != nil {
		panic(err)
	}
	fmt.Printf("java is [%s]\n", s)
}

func main() {
	fmt.Println("hello")
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
