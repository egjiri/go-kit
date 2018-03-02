package exec

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Execute runs a command directly on the underlying OS
func Execute(commands ...string) {
	joinedCommands := strings.Join(commands, " ")
	args := strings.Split(joinedCommands, " ")
	name, err := exec.LookPath(args[0])
	if err != nil {
		log.Fatal(err)
	}
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	p, err := os.StartProcess(name, args, &procAttr)
	if err != nil {
		log.Fatal(err)
	}
	p.Wait()
}
