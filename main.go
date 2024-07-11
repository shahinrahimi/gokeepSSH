package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	l := log.New(os.Stdout, "[gokeepSSH] ", log.LstdFlags)
	if err := godotenv.Load(); err != nil {
		l.Fatal("Error loading .env file")
	}

	vpsAddress := os.Getenv("VPS_ADDRESS")
	localPort := os.Getenv("LOCAL_PORT")
	sshKey := os.Getenv("SSH_KEY_PATH")
	var counter int = 0
	if vpsAddress == "" || localPort == "" || sshKey == "" {
		l.Fatal("Wrong local variables.")
	}
	for {
		err := startSSHTunnel(vpsAddress, localPort, sshKey, l)
		counter++

		if err != nil {
			l.Printf("SSH tunnel error[%d]: %v", counter, err)
		}
		l.Println("SSH trunnel terminated, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

}

func startSSHTunnel(vpsAddress, localPort, sshKey string, l *log.Logger) error {
	cmd := exec.Command("ssh", "-i", sshKey, "-D", localPort, "-N", vpsAddress)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	l.Printf("Starting SSH tunnel to %s on local port %s", vpsAddress, localPort)
	return cmd.Run()
}
