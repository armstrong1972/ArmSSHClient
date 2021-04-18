package main

import (
	cpt "armstrong/arm_crypt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type SshConfig struct {
	Mod    string
	Addr   string
	User   string
	Cipher string
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	return err
}

var sExePath string

func main() {

	if len(os.Args) < 2 {
		fmt.Println("> sh config-name")
		os.Exit(1)
	}

	sExePath = getCurrentDirectory()
	//fmt.Println(sExePath)

	sfConfig := "./config/" + os.Args[1] + ".json"
	//sfConfig := "./config/sample.json"
	//fmt.Println(sfConfig)

	if !FileExist(sfConfig) {
		sfConfig = sExePath + "/config/" + os.Args[1] + ".json"
	}

	JsonParse := NewJsonStruct()
	sshC := SshConfig{}

	e := JsonParse.Load(sfConfig, &sshC)
	if e != nil {
		os.Exit(2)
	}
	//fmt.Println(sshC)

	ConnetSSH(sshC)
}

func ConnetSSH(sshC SshConfig) {
	cliConfig := ssh.ClientConfig{
		User:            sshC.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshC.Mod == "pem" {
		sfPem := "./config/pem/" + sshC.Cipher
		if !FileExist(sfPem) {
			sfPem = sExePath + "/config/pem/" + sshC.Cipher
		}
		cliConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sfPem)}
	} else {
		// sshC.Mod == "pwd"
		sPriKey := "MIIEpAIBAAKCAQEAwIc4VUrADMGHk7ywstSENsswvAsau0q8Rui+zbPC6OsyAcXAo/MFBlZv+G5YSz1yQqC/b9asMs2Y3QO9grhdPOdfhgkw14Ok/IGTewFKEEyO4evdeRtKlPcjiRP/RVwiAa6hoFlfEDkQlUu5h0taZz7yKkkAL33CC32vPTJCj/U+7Y1YPpGUnbzkZVE0Q2wpu9i73niM9nwhcrm2XR7K6Hn3uXV5UcmuBEELw7ZfXl2RsQmFHlNTbO8n7aTXLQAReOi/JDEcRoWkfazdqsEA8YppVbWue4wqth42SzEcgR7Otv8aVvxNDl4fO2kzZjfx/jiL/EaZdFfpHwLZOEFb7wIDAQABAoIBAQCWTHmjzBklXCfVI6bl1JXqmLFQ+3JA6FUXHjdmKoFsh7+gNpru4pb3nZ8H2EzBYBTFAuywCGyhtHMzhun5FKhlFVMzNhWVCUZRgW56xkTleH7Ky/E3zCBrLovlRWEw6n35xjSTE3HtHj9d7aHGhwLqOu+eJyMn8Ar+IX4eThA8hKJWp+us8Bhteg8JjtzStBA/ZNpIbo+dnGKizXfENGCIRctQwUp9ds+wifSj8G2ZtTMrO4ktUlVcSwN/y1fOxic8V0JIbJHEczdEwq9c/HCHMwdoDNaf3Ml9jTjyc8EmMkkJbD06BLzQGsHvjSAiwblnY01xU7hRrqTdZl2u6OURAoGBAMvGvzAxuPlEF7MLi1DstW5wiab/M15LRb4Skl3rM3D7LsjMKrNy29A9vMrtKQV40DXIbq+3+DKwpUxOiq5ngYfpTj60QSMl0oQq55AAfDQJ28as3fJiYKLPDE+iJvPA95P4/T8I9OB31cgYN1AdYeqfbwiUdxYo9fPDw8AlPlwTAoGBAPHegqYUVTPV4pNlNkUA3ZsQHpZP0c5yT2pf9olSqllu7aLM2E/PNgObVhpmbHBkiE0GgORtO8f8i1j4cbx7Alrp8rbMl4bMO2s4EjW4D1ppbkEr+pPGXxzFez8oWcUv2+W7B1ZlFhdFPpC/qms52wqtWhixcw5XNORp5CUgeAQ1AoGBAMZ0A2RbOErBztbXClHAhZ1NyjEx3oQZiI1optUCp3GDnnd9pqRm2r4+Mevq9gVXOGb4kRtebkBnmkPrRdI+CX8kbshQYhNp1VBUPHoYjt1bMIeXePuCZZyJSfMP9yFyr2qTlYbbISiubCvXsZ67Ts3hgY/4jtWtBILnB2/MlaOzAoGAIt1nSddufL1dHtAdJY89YofNUO+Kl87Egdn5aMwgwCmMWcZbfA2rhJOUstOG2CC1wRyp4cOZNXfDAn01r+yvZzmIAi+1u/meqxL3FQVGOUsvWHeldD3JIhWZcX44ioMemJwAL8T8jTgvD6CBSVmaqIxai2qw4iVR+4cEGxH2Gu0CgYBoCODl23tLzA8J4Ix/MIsRdqm+NH2JDi/xAU3JizagCzuQt8ZL5HDKRve/oFT9x22AW+J4hNeT/eyXTi6taucnHOc1lMQIenwofrqAVm/ox3iy5EL1ourkZI08njdA/cqj8LyJ63l6jBaM3br3pv1z53KCnywyIfBConJcrTEY0w=="
		PriKey := cpt.Str2Buff(sPriKey)

		encBuff := cpt.Str2Buff(sshC.Cipher)
		sPwd := string(cpt.RSA_Decrypt(encBuff, PriKey))
		cliConfig.Auth = []ssh.AuthMethod{ssh.Password(sPwd)}
	}

	client, err := ssh.Dial("tcp", sshC.Addr, &cliConfig)
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	//Clean Screen
	for i := 0; i <= 100; i++ {
		fmt.Println()
	}
	fmt.Println("******************************************************")
	fmt.Println("*                                                    *")
	fmt.Println("*    Welcome Armstrong's SSH Client                  *")
	fmt.Println("*                                                    *")
	fmt.Println("*                     Copyright by ArtCubeVR Ltd.    *")
	fmt.Println("*                                                    *")
	fmt.Println("******************************************************")

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}
	defer session.Close()

	// Fixed Ctrl+C to exit app
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
	session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}
	if err = session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}
	if err = session.Wait(); err != nil {
		log.Fatalf("return error: %s", err.Error())
	}
}

func publicKeyAuthFunc(keyPath string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Get path of Exe failed", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
