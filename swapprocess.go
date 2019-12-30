package main

import "fmt"
//import "io/ioutil"
import "os/exec"

func main() {

    dateCmd := exec.Command("bash","-c", "date && ls -l /tmp")

    // `.Output` is another helper that handles the common
    // case of running a command, waiting for it to finish,
    // and collecting its output. If there were no errors,
    // `dateOut` will hold bytes with the date info.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))
}
