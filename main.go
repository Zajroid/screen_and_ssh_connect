package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
)

// cat /proc/cpuinfo | grep name | wc -l
// cat /proc/cpuinfo | grep name | head -n 1 | awk '{print $4,$5,$6,$7,$8,$9;}'
// free -m | grep Mem | awk '{print $2 ,$3, $4, $5, $6, $7}'
// ls -lh $(which ls)
// crontab -l
// w
// uname -m
// cat /proc/cpuinfo | grep model | grep name | wc -l
// top
// uname
// lscpu | grep Model
var (
	execCommand     = []string{"ls", "uname", "w"}
	dateAboutSystem []string
)

func main() {
	for i := 0; i < len(execCommand); i++ {
		out, err := exec.Command(execCommand[i]).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		dateAboutSystem = append(dateAboutSystem, execCommand[i])
	}

	createScreenshot()

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	cliConf := new(ClientConfig)
	fmt.Println(
		cfg.Section("server").Key("ip_address").MustString(""),
		cfg.Section("server").Key("tcp_port").MustInt64(),
		cfg.Section("server").Key("username").MustString(""),
		cfg.Section("server").Key("password").MustString(""),
	)

	cliConf.createClient(
		cfg.Section("server").Key("ip_address").MustString(""),
		cfg.Section("server").Key("tcp_port").MustInt64(),
		cfg.Section("server").Key("username").MustString(""),
		cfg.Section("server").Key("password").MustString(""),
	)

	fmt.Println(cliConf.RunShell("ls -l"))
}
