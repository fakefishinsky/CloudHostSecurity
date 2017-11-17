package compliance

type RemediationCommandUnit struct {
    Id string
    Type string
    Command string
    CommandOutput string
    CommandError string
}

type RemediationItem struct {
    Description string
    Commands []RemediationCommandUnit
    Run string
    Result bool
}

func (self *RemediationItem) Remedy() error {
    return nil
}