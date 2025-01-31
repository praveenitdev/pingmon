package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func playSound(soundFile string) {
	var cmd *exec.Cmd

	if _, err := os.Stat(soundFile); os.IsNotExist(err) {
		fmt.Println("Audio file not found:", soundFile)
		return
	}

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("afplay", soundFile)
	case "linux":
		cmd = exec.Command("aplay", soundFile) // Or "play"
	case "windows":
		cmd = exec.Command("powershell", "-c", "(New-Object Media.SoundPlayer '"+soundFile+"').PlaySync();")
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
				playSound("up-audio.mp3")
			}
		} else {
			fmt.Println("Host is down!")
			if alertType == "down" {
				playSound("down-audio.mp3")
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
