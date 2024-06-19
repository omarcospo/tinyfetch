package main

// DONE:
// User: $USER
// Shell: $SHELL.
// Memory: /proc/meminfo.
// CPU: /proc/cpuinfo.
// Locale: $LANG.
// Uptime: /proc/uptime
// OS: /etc/os-release
// FIXME:
// TODO:
// Local IP: /etc/network/interfaces.
// Swap: /proc/swaps
// Disk: /etc/fstab, /proc/mounts
// Kernel Version: /proc/version
// Battery: /sys/class/power_supply/ (e.g., /sys/class/power_supply/BAT0/ for the first battery).
// Power Adapter: /sys/class/power_supply/ (e.g., /sys/class/power_supply/AC/ for the AC adapter).
// Host: /etc/hostname

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	reset string = "\033[0m"
	red   string = "\033[31m"
)

func ReadInfoFile(path string, tag string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		errMsg := errors.New("\n" + red + "---> " + tag + "INFO NOT FOUND" + reset)
		fmt.Println(errMsg)
		os.Exit(1)
	}
	return string(file)
}

func CutStrPrefix(content string, prefix string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			info := strings.TrimPrefix(line, prefix)
			info = strings.Trim(strings.Trim(info, "\""), "  ")
			return info
		}
	}
	return ""
}

func GetInfo(path string, prefix string, tag string) {
	file := ReadInfoFile(path, tag)
	info := CutStrPrefix(file, prefix)
	fmt.Println(red+tag+reset, info)
}

func GetEnv(env string, tag string) {
	var envCmd string = os.Getenv(env)
	if envCmd == "" {
		errMsg := errors.New("\n" + red + "---> " + tag + "INFO NOT FOUND" + reset)
		fmt.Println(errMsg)
		os.Exit(1)
	}
	fmt.Println(red+tag+reset, envCmd)
}

func GetUptime() {
	var tag string = "Uptime:"
	var file = ReadInfoFile("/proc/uptime", tag)
	info, _ := strconv.ParseFloat(strings.Split(file, " ")[0], 64)
	if info >= 3600 {
		fmt.Printf("%s%s%s %.0fh\n", red, tag, reset, info/3600)
	} else {
		fmt.Printf("%s%s%s %.2fmin\n", red, tag, reset, info/60)
	}
	return
}

func GetMemory() {
	var tag string = "Memory:"
	var prefix string = "MemTotal:"
	var file = ReadInfoFile("/proc/meminfo", tag)
	for _, line := range strings.Split(file, "\n") {
		var info = CutStrPrefix(line, prefix)
		infoFloat, _ := strconv.ParseFloat(strings.Trim(info, " kB"), 64)
		fmt.Printf("%s%s%s %.2fGB\n", red, tag, reset, (infoFloat/1024)/1024)
		return
	}
}

func main() {
	GetInfo("/etc/os-release", "NAME=", "Distro:")
	GetEnv("USER", "User:")
	GetEnv("SHELL", "Shell:")
	GetInfo("/proc/cpuinfo", "model name\t:", "CPU:")
	GetMemory()
	GetUptime()
	GetEnv("LANG", "Locale:")
}
