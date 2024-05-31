package main

// OS: /etc/os-release
// Host: /etc/hostname
// User: $USER
// Linux Version: /etc/os-release
// Packages:
// Shell: $SHELL.
// CPU: /proc/cpuinfo.
// Memory: /proc/meminfo.
// Swap: /proc/swaps
// Disk: /etc/fstab, /proc/mounts
// Local IP: /etc/network/interfaces.
// Battery: /sys/class/power_supply/ (e.g., /sys/class/power_supply/BAT0/ for the first battery).
// Power Adapter: /sys/class/power_supply/ (e.g., /sys/class/power_supply/AC/ for the AC adapter).
// Locale: $LANG.

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reset string = "\033[0m"
	red   string = "\033[31m"
)

func main() {
	get_info("/etc/os-release", "NAME=", "Distro:")
	get_env("USER", "User:")
	get_env("SHELL", "Shell:")
	get_info("/proc/cpuinfo", "model name\t:", "CPU:")
	get_memory()
	get_env("LANG", "Locale:")
}

func errDo(err error, tag string) {
	if err != nil {
		errMsg := errors.New("\n" + red + "---> " + tag + "INFO NOT FOUND" + reset)
		fmt.Println(errMsg)
		os.Exit(1)
	}
}

func get_env(env string, tag string) {
	var envCmd string = os.Getenv(env)
	if envCmd == "" {
		errMsg := errors.New("\n" + red + "---> " + tag + "INFO NOT FOUND" + reset)
		fmt.Println(errMsg)
		os.Exit(1)
	}
	fmt.Println(red+tag+reset, os.Getenv(env))
}

func get_info(path string, prefix string, tag string) {
	file, err := os.ReadFile(path)
	errDo(err, tag)
	content := string(file)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			info := strings.TrimPrefix(line, prefix)
			info = strings.Trim(info, "\"")
			info = strings.Trim(info, "  ")
			fmt.Println(red+tag+reset, info)
			return
		}
	}
}

func get_memory() {
	var prefix string = "MemTotal:"
	var tag string = "Memory:"
	file, err := os.ReadFile("/proc/meminfo")
	errDo(err, tag)
	content := string(file)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			info := strings.TrimPrefix(line, prefix)
			info = strings.Trim(info, "\"")
			info = strings.Trim(info, "  ")
			info = strings.Trim(info, " kB")
			infoFloat, _ := strconv.ParseFloat(info, 64)
			fmt.Printf("%s%s%s %.2fGB\n", red, tag, reset, (infoFloat/1024)/1024)
			return
		}
	}
}
