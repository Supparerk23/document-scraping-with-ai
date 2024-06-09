package service

import(
	"os"
	// "os/exec"
	"bytes"
	"fmt"
)

func (s *service)ProcessPDF() error {
	// path, pathErr := os.Getwd()
	// if pathErr != nil {
	//     fmt.Print(pathErr)
	// }
	// handle err
    fmt.Println("Run python script")
    var out bytes.Buffer
	var stderr bytes.Buffer
	s.command.Stdout = &out
	s.command.Stderr = &stderr
	err := s.command.Run()
	if err != nil {
		fmt.Println("PDF Error")
	    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	    return err
	}
	// fmt.Println("Result: " + out.String())
	// fmt.Println("Command finished")
	return nil
}

func (s *service)WriteResult(path string, res string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// l, err := f.WriteString(res)
	_, err = f.WriteString(res)
	if err != nil {
        f.Close()
		return err
	}

	// fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}