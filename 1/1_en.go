package main

import "fmt"

// WeddingInvitation holds the details of our special day
type WeddingInvitation struct {
	Groom    string // 신랑 이름
	Bride    string // 신부 이름
	Date     string // 결혼식 날짜
	Time     string // 결혼식 시간
	Location string // 결혼식 장소
}

// TogetherForever symbolizes the union of two hearts
const TogetherForever = true

func main() {
	invitation := WeddingInvitation{
		Groom:    "이병우",
		Bride:    "오예영",
		Date:     "2025-05-17",
		Time:     "12:20 PM",
		Location: "The Link Hotel 7F White Hall",
	}

	fmt.Println("=======================================")
	fmt.Println("💌   You're Invited to Our Wedding   💌")
	fmt.Println("=======================================")
	fmt.Printf("👰 신부: %s\n", invitation.Bride)
	fmt.Printf("🤵 신랑: %s\n", invitation.Groom)
	fmt.Println("────────────────────────────────────────")

	fmt.Printf("📅 날짜 : %s\n", invitation.Date)
	fmt.Printf("⏰ 시간 : %s\n", invitation.Time)
	fmt.Printf("📍 장소 : %s\n", invitation.Location)

	fmt.Println("────────────────────────────────────────")

	// A special message to our guests
	message := `
We invite you to kindly join us in 
celebrating this special moment that marks the 
beginning of our love story."
We would be grateful if you could add warmth and 
blessings to make this day even more meaningful.
`
	fmt.Println(message)
}
