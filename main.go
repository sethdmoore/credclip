package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func main() {
	var string_builder string
	section := "default"

	aws_config := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")

	if len(os.Args) > 1 {
		section = os.Args[1]
	}

	ini, err := ini.Load(aws_config)
	if err != nil {
		fmt.Printf("Error loading INI: %v\n", err)
		os.Exit(2)
	}

	access_key := ini.Section(section).Key("aws_access_key_id").String()
	secret_key := ini.Section(section).Key("aws_secret_access_key").String()

	string_builder = fmt.Sprintf("unset HISTFILE; export AWS_ACCESS_KEY_ID=\"%s\"; export AWS_SECRET_ACCESS_KEY=\"%s\";",
		access_key,
		secret_key,
	)

	clipboard.WriteAll(string_builder)
}
