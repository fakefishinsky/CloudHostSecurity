package compliance

import (
    "fmt"
)

type AuditCommandUnit struct {
    Id string
    Type string
    Command string
    ExpectReg string
    ExpectFlag string
    CommandOutput string
    CommandError string
}

type AuditItem struct {
    Description string
    Commands []AuditCommandUnit
    Run string
    Result bool
}

func (self *AuditItem) Audit() error {
    for _, auditcmd := range self.Commands {
        cmd := &Command{auditcmd.Type, auditcmd.Command}
        _, auditcmd.CommandOutput, auditcmd.CommandError = cmd.Run()
        fmt.Println(auditcmd.CommandOutput)
    }
    return nil
}
