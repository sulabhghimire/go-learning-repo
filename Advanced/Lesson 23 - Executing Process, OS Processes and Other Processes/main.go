package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("output:", string(output))

	// // ========= INTERACTION VIA PIPE
	// pr, pw := io.Pipe()

	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("foo is bad\nbar\nbax\n"))
	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("output:", string(output))

	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error starting command:", err)
	// 	return
	// }
	// fmt.Println("output:", string(output))

	// // ============= KILLING A PROCESS
	// cmd := exec.Command("sleep", "5")
	// // Start command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("error starting command:", err)
	// 	return
	// }

	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("error killing process:", err)
	// 	return
	// }
	// fmt.Println("Process killed")

	// // ===================== WAITING FOR PROCESS TO COMPLETE
	// cmd := exec.Command("sleep", "5")
	// // Start command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("error starting command:", err)
	// 	return
	// }

	// // waiting for program to finish
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("error waiting:", err)
	// 	return
	// }
	// fmt.Println("Process completed")

	// // ===================== PROVIDING READER TO OUR PROCESS
	// cmd := exec.Command("grep", "foo")
	// // Set input for our command
	// cmd.Stdin = strings.NewReader("foo is \nbar\nbaz\nfoo")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	// // ===================== RUNNING A SIMPLE PROCESS
	// cmd := exec.Command("echo", "Hello World!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

}
