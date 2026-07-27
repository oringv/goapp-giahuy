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

// VIẾT HOA chữ cái đầu để các package khác có thể gọi được
func ClearScreen() {
	if runtime.GOOS == "windows" {
		exec.Command("cls").Run()
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
