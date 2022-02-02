package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	outputString := string(output)

	re := regexp.MustCompile("<vault-kv-get>.*</vault-kv-get>")

	matches := re.FindAllString(outputString, -1)

	for _, match := range matches {
		args := strings.Replace(strings.Replace(match, "<vault-kv-get>", "", -1), "</vault-kv-get>", "", -1)
		err, secret, _ := bashOut(fmt.Sprintf("vault kv get %s", args))
		if err != nil {
			panic("some error found")
		}
		outputString = strings.Replace(outputString, match, secret, -1)
	}

	fmt.Println(outputString)

}

func bashOut(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
