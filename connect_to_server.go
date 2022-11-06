package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

// ClientConfig Connection configuration
type ClientConfig struct {
	Host       string      //ip
	Port       int64       // Port
	Username   string      //Username
	Password   string      //Password
	Client     *ssh.Client //ssh client
	LastResult string      //Result of the last run
}

func (cliConf *ClientConfig) createClient(host string, port int64, username, password string) {
	var (
		client *ssh.Client
		err    error
	)
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password
	cliConf.Port = port

	//Generally pass in four parameters: user, []ssh.AuthMethod{ssh.Password(password)}, HostKeyCallback, timeout,
	config := ssh.ClientConfig{
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	//Get client
	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred:", err)
	}

	cliConf.Client = client
}

func (cliConf *ClientConfig) RunShell(shell string) string {
	var (
		session *ssh.Session
		err     error
	)

	//Get session, this session is used to perform operations remotely
	if session, err = cliConf.Client.NewSession(); err != nil {
		log.Fatalln("error occurred:", err)
	}

	//Execute shell
	if output, err := session.CombinedOutput(shell); err != nil {
		log.Fatalln("error occurred:", err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}
