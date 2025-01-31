package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var soundFiles embed.FS

func playSound(filename string) {
	data, err := soundFiles.ReadFile("sounds/" + filename)
	if err != nil {
		fmt.Println("Error reading embedded sound file:", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "sound-*.mp3")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.Write(data)
	tmpFile.Close()

	filename = tmpFile.Name()

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", filename)
	case "linux":
		cmd = exec.Command("aplay", filename)
	case "windows":
		cmd = exec.Command("powershell", "-c", "(New-Object Media.SoundPlayer '"+filename+"').PlaySync();")
	default:
		fmt.Println("Unsupported OS for playing audio.")
		return
	}

	cmd.Start()
}

func pingHost(host string, interval int, alertType string) {
	for {
		cmd := exec.Command("ping", "-c", "1", host)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("ping", "-n", "1", host)
		}

		if err := cmd.Run(); err == nil {
			fmt.Println("Host is up!")
			if alertType == "up" {
				playSound("upAlert.mp3")
			}
		} else {
			fmt.Println("Host is down!")
			if alertType == "down" {
				playSound("downAlert.mp3")
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func main() {
	host := flag.String("host", "", "Host to ping")
	shortHost := flag.String("h", "", "Host to ping (shorthand)")
	timeInterval := flag.Int("time", 5, "Ping interval in seconds")
	shortTime := flag.Int("t", 5, "Ping interval in seconds (shorthand)")
	alertType := flag.String("alert", "up", "Alert type: up/down")

	flag.Parse()

	finalHost := *host
	if *shortHost != "" {
		finalHost = *shortHost
	}

	interval := *timeInterval
	if *shortTime != 5 {
		interval = *shortTime
	}

	if finalHost == "" {
		fmt.Println("Please provide a host using --host or -h")
		os.Exit(1)
	}

	fmt.Println("Pinging", finalHost, "every", interval, "seconds...")
	pingHost(finalHost, interval, *alertType)
}
