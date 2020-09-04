package main

import "golang.org/x/crypto/ssh"

type Tunnel struct {
	Local string
	Remote string
	SSHClient *ssh.Client
}

func NewTunnel(username string, password string) (*Tunnel, error) {
	tun := new(Tunnel)
	cfg := &ssh.ClientConfig{
		User:              username,
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
	}

	cfg.Auth = append(cfg.Auth, ssh.Password(password))

	client, err := ssh.Dial("tcp", tun.Local, cfg)
	if err != nil {
		return nil, err
	}

	client.Dial("tcp", tun.Remote)

	tun.SSHClient = client

	return tun, nil
}