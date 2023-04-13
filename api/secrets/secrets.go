package secrets

import (
	"log"
	"os"
)

var (
	IpmiUsername string
	IpmiPassword string
)

func ReadSecret() {
	// Read the credentials from the secret volume
	ipmiUsernameByte, err := os.ReadFile("/etc/secret-volume/ipmiUsername")
	if err != nil {
		log.Fatalf("Failed to read ipmiUsername: %s", err)
	}

	ipmiPasswordByte, err := os.ReadFile("/etc/secret-volume/ipmiPassword")
	if err != nil {
		log.Fatalf("Failed to read ipmiPassword: %s", err)
	}

	IpmiUsername = string(ipmiUsernameByte)
	IpmiPassword = string(ipmiPasswordByte)

}
