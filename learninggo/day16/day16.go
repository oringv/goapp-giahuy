package main

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/utils"
)

func main() {

	for {
		fmt.Println("\n📚 CHUONG TRINH QUAN LY THU VIEN")
		fmt.Println("1. Them sach")
		fmt.Println("2. Xem danh sach sach")
		fmt.Println("3. Them nguoi muon")
		fmt.Println("4. Xem danh sach nguoi muon")
		fmt.Println("5. Muon sach")
		fmt.Println("6. Xem lich su muon")
		fmt.Println("7. Tra sach")
		fmt.Println("8. Tim kiem sach")
		fmt.Println("9. Thoat")

		choice := utils.GetPositiveInt("👉 Chon chu nang: ")

		utils.ClearScreen()

		switch choice {
		case 1:
			fmt.Println("-=-=-=-=- Them Sach -=-=-=-=-")
		case 2:
			fmt.Println("-=-=-=-=- Xem Danh Sach Sach -=-=-=-=-")
		case 3:
			fmt.Println("-=-=-=-=- Them Nguoi Muon Sach -=-=-=-=-")
		case 4:
			fmt.Println("-=-=-=-=- Xem Danh Sach Nguoi Muon -=-=-=-=-")
		case 5:
			fmt.Println("-=-=-=-=- Muon Sach -=-=-=-=-")
		case 6:
			fmt.Println("-=-=-=-= Xem Lich Su Muon Sach -=-=-=-=-=")
		case 7:
			fmt.Println("-=-=-=-=- Tra Sach -=-=-=-=-")
		case 8:
			fmt.Println("-=-=-=-=- Tim Kiem Sach -=-=-=-=-")
		case 9:
			fmt.Println("👋 Tam biet!")
			return
		default:
			fmt.Println("⚠️ Lua chon khong hop le!")
		}
		utils.ReadInput("\nNhấn Enter để tiếp tục...")
	}
}
