package fingerprint

import (
	"crypto/sha256"
	"fmt"
	"os/exec"
	"strings"
)

func getOutput(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func GetMachineFingerprint() string {
	board := getOutput("dmidecode", "-s", "baseboard-serial-number")
	cpu := getOutput("cat", "/proc/cpuinfo")
	disk := getOutput("lsblk", "-o", "NAME,SERIAL")

	raw := board + cpu + disk + "cc"
	hash := sha256.Sum256([]byte(raw))
	return fmt.Sprintf("%x", hash[:])
}
