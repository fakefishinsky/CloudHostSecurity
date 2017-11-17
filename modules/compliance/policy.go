package compliance

import (
    io "io/ioutil"
    json "encoding/json"
)

type PolicyUnit struct {
    Id string
    Title string
    Description string
    Rationale string
    Audit AuditItem
    Remediation RemediationItem
}

type Policy struct {
    Policies []PolicyUnit
}

func (self *Policy) LoadPolicy(policyFile string) error {
    data, err := io.ReadFile(policyFile)
    if err != nil {
        return err
    }

    jsondata := []byte(data)
    return json.Unmarshal(jsondata, &self.Policies)
}

func (self *Policy) Audit() error {
    for _, policy := range self.Policies {
        auditItem := &policy.Audit
        auditItem.Audit()
    }
    return nil
}

func (self *Policy) Remediation() error {
    for _, policy := range self.Policies {
        remediationItem := &policy.Remediation
        remediationItem.Remedy()
    }
    return nil
}