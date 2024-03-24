package service

import(
	// "os"
	// "os/exec"
	"bytes"
	"fmt"
)

func (s *service)ProcessPDF() {
	// path, pathErr := os.Getwd()
	// if pathErr != nil {
	//     fmt.Print(pathErr)
	// }
	// handle err
    fmt.Println("Run python")
    var out bytes.Buffer
	var stderr bytes.Buffer
	s.command.Stdout = &out
	s.command.Stderr = &stderr
	err := s.command.Run()
	if err != nil {
		fmt.Println("PDF Error")
	    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	    return
	}
	fmt.Println("Result: " + out.String())
}