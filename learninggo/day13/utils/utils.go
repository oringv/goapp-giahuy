package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := Reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		val, err := strconv.Atoi(input)
		if err == nil && val >= 0 {
			return val
		}
		fmt.Println("❌ Lỗi: Vui lòng nhập số nguyên dương!")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		val, err := strconv.ParseFloat(input, 64)
		if err == nil && val >= 0 {
			return val
		}
		fmt.Println("❌ Lỗi: Vui lòng nhập số dương!")
	}
}

// Các hàm bổ trợ để không bị lỗi code cũ
func ReadInt(p string) int       { return GetPositiveInt(p) }
func ReadString(p string) string { return ReadInput(p) }
