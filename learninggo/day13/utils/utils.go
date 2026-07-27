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

// ✅ THÊM HÀM NÀY: Nhập chuỗi, nhấn Enter để giữ nguyên giá trị cũ
func GetOptionalString(prompt string, oldValue string) string {
	fmt.Printf("%s [%s]: ", prompt, oldValue)
	input, _ := Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return oldValue
	}
	return input
}

// ✅ THÊM HÀM NÀY: Nhập số thực, nhấn Enter để giữ nguyên giá trị cũ
func GetOptionalPositiveFloat(prompt string, oldValue float64) float64 {
	for {
		fmt.Printf("%s [%.2f]: ", prompt, oldValue)
		input, _ := Reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			return oldValue
		}
		val, err := strconv.ParseFloat(input, 64)
		if err == nil && val >= 0 {
			return val
		}
		fmt.Println("❌ Lỗi: Nhập số dương hoặc nhấn Enter để bỏ qua!")
	}
}
