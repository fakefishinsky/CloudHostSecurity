package compliance

import (
    "bytes"
    "strings"
    "os/exec"
    "errors"
)

type Command struct {
    Type string
    Cmd string
}

func ShellCommand(command string, shell string) (error, string, string) {
    var cmdout bytes.Buffer
    var cmderr bytes.Buffer

    cmd := exec.Command(shell, "-c", command)
    cmd.Stdout = &cmdout
    cmd.Stderr = &cmderr

    err := cmd.Run()
    return err, cmdout.String(), cmderr.String()
}

func (self *Command) Run() (error, string, string) {
    if strings.Contains(strings.ToLower(self.Type), "shell:") {
        cmdtype := strings.Split(self.Type, ":")
        if (len(cmdtype) == 2) && (strings.ToLower(cmdtype[0]) == "shell") {
            return ShellCommand(self.Cmd, cmdtype[1])
        }
    } else if self.Type == "builtin" {
        return nil, "", ""
    }
    
    return errors.New("Unable to identify the command type(" + self.Type + ")."), "", ""
}
