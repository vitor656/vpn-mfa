package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/xlzd/gotp"
)

func main() {
	user := os.Getenv("VPN_USER")
	if user == "" {
		log.Fatal("VPN_USER não informado")
	}
	prefix := os.Getenv("VPN_PREFIX")

	configPath := os.Getenv("VPN_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("VPN_CONFIG_PATH não informado")
	}

	mfaCode := os.Getenv("VPN_MFA")
	if mfaCode == "" {
		log.Fatal("VPN_MFA não informado")
	}

	for {
		loadCredentials(mfaCode, prefix, user)

		cmd := exec.Command("openvpn", "--config", configPath, "--auth-user-pass", "vpn_tmp")
		stdout, _ := cmd.StdoutPipe()

		fmt.Println("Conectando...")
		if err := cmd.Start(); err != nil {
			log.Fatal("Erro ao executar comando: ", err)
			break
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}

		cmd.Wait()
		fmt.Println("Desconectado. Tentando reconectar...")
	}

	fmt.Println("Finalizando...")
}

func loadCredentials(mfaKey, prefix, user string) {
	totp := gotp.NewDefaultTOTP(mfaKey)
	pass := prefix + totp.Now()

	if err := saveToFile(user, pass); err != nil {
		log.Fatal("Erro ao salvar arquivo de credenciais: ", err)
	}
}

func saveToFile(user, pass string) error {
	file, err := os.Create("vpn_tmp")
	if err != nil {
		return err
	}
	defer file.Close()

	content := user + "\n" + pass

	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}

	return nil
}
