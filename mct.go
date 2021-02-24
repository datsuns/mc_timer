package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/rivo/tview"
)

var TotalWaitTimes = 6
var MinutesPerWait time.Duration = 1

var WaitToMessage = time.Minute * MinutesPerWait

//var WaitToMessage = time.Second * 10

var JavaExec = "java"
var ServerJarName = "server.jar"
var ServerRunParams = []string{
	"-Xmx1024M",
	"-Xms1024M",
	"-jar",
	ServerJarName,
	"nogui",
}

func runserver() {
	exepath := os.Args[0]
	targetpath := filepath.Dir(exepath)

	os.Chdir(targetpath)
	cmd := exec.Command(JavaExec, ServerRunParams...)
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
		myin.Scan()
		fmt.Println("wait start")
		var elapsed time.Duration = 0
		for i := 0; i < TotalWaitTimes; i++ {
			time.Sleep(WaitToMessage)
			elapsed += MinutesPerWait
			text := fmt.Sprintf("/say %d minutes\n", elapsed)
			fmt.Println("exec say")
			writer.Write([]byte(text))
		}
	}()
	cmd.Start()
	<-done
	cmd.Wait()
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

func app_on_tview() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	grid := tview.NewGrid().
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("hello")
	//runserver()
	app_on_tview()
}
