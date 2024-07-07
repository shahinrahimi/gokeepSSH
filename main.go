package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	vpsAddress := os.Getenv("VPS_ADDRESS")
	localPort := os.Getenv("LOCAL_PORT")
	sshKey := os.Getenv("SSH_KEY_PATH")
	if vpsAddress == "" || localPort == "" || sshKey == "" {
		log.Fatal("Wrong local variables.")
	}
	for {
		err := startSSHTunnel(vpsAddress, localPort, sshKey)
		if err != nil {
			log.Printf("SSH tunnel error: %v", err)
		}
		log.Println("SSH trunnel terminated, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

}

func startSSHTunnel(vpsAddress, localPort, sshKey string) error {
	cmd := exec.Command("ssh", "-i", sshKey, "-D", localPort, "-N", vpsAddress)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	log.Println("Starting SSH tunnel to %s on local port %s", vpsAddress, localPort)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start SSH tunnel: %w", err)
	}
	return nil
}
