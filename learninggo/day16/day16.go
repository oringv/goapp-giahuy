package day16

import "fmt"

func mani() {

	for {
		fmt.Println(" CHUONG TRINH QUAN LY THU VIEN")
		fmt.Println("1. Them sach")
		fmt.Println("2. Xem danh sach sach")
		fmt.Println("3. Them nguoi muon")
		fmt.Println("4. Xem danh sach nguoi muon")
		fmt.Println("5. Muon sach")
		fmt.Println("6. Xem lich su muon")
		fmt.Println("7. Tra sach")
		fmt.Println("8. Tim kiem sach")
		fmt.Println("9. Thoat")
		choice := util.GetPostiveInt("Chon chu nang :")

		switch choice {
		case 1:
			fmt.Println("-=-=-=-=- Them Sach -=-=-=-=-")
		case 2:
			fmt.Println("-=-=-=-=- Xem Danh Sach Sach -=-=-=-=-")
		case 3:
			fmt.Println("-=-=-=-=- Them Nguoi Muon Sach -=-=-=-=-")
		case 4:
			fmt.Println("-=-=-=-=- Xem Danh Sach Nguoi Muon -=-=-=-=-")

		}
	}
}
