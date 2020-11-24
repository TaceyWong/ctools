package utils

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
)

// ExecShellCMD 执行Sheel命令
func ExecShellCMD(shellType, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellType, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// RegexGetGroup 实现正则group功能
func RegexGetGroup(compRegEx *regexp.Regexp, text string) (map[string]string, error) {
	match := compRegEx.FindStringSubmatch(text)
	groupMap := make(map[string]string)
	err := errors.New("no match")
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			groupMap[name] = match[i]
			err = nil
		}
	}
	return groupMap, err
}
