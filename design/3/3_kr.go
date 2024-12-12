package main

import "fmt"

// WeddingInvitation은 특별한 날의 세부 정보를 나타냅니다.
type WeddingInvitation struct {
	Groom    string // 신랑 이름
	Bride    string // 신부 이름
	Date     string // 결혼식 날짜
	Time     string // 결혼식 시간
	Location string // 결혼식 장소
}

// TogetherForever는 두 마음의 결합을 상징합니다.
const TogetherForever = true

func main() {
	// 결혼식 초대장 세부 정보
	invitation := WeddingInvitation{
		Groom:    "이병우",
		Bride:    "오예영",
		Date:     "2025-05-17",
		Time:     "오후 12:20",
		Location: "더 링크 호텔 7층 화이트 홀",
	}

	// 🌸 하객들에게 전하는 특별한 메시지 🌸
	message := `
	저희 사랑의 새로운 시작을 함께 축하해 주시면 감사하겠습니다. 
	이 날이 더욱 특별하고 따뜻한 추억으로 남을 수 있도록 
	귀한 축복으로 자리 빛내 주시길 바랍니다.
	`

	fmt.Println("=======================================")
	fmt.Println("💍   저희 결혼식에 초대합니다!  💍")
	fmt.Println("=======================================")
	fmt.Printf("👰 신부: %s\n", invitation.Bride)
	fmt.Printf("🤵 신랑: %s\n", invitation.Groom)
	fmt.Println("---------------------------------------")
	fmt.Printf("📅 날짜 : %s\n", invitation.Date)
	fmt.Printf("⏰ 시간 : %s\n", invitation.Time)
	fmt.Printf("📍 장소 : %s\n", invitation.Location)
	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Println("🌸 특별한 메시지 🌸")
	fmt.Println(message)
	fmt.Println("---------------------------------------")
}
