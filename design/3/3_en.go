package main

import "fmt"

// WeddingInvitation represents the details of a special day
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
	// Wedding invitation details
	invitation := WeddingInvitation{
		Groom:    "이병우",
		Bride:    "오예영",
		Date:     "2025-05-17",
		Time:     "12:20 PM",
		Location: "The Link Hotel 7F White Hall",
	}

	// A special message to our guests
	message := `
	"Together, we begin a new chapter of our love story.
	We warmly invite you to join us on this joyous day 
	to celebrate our union and share in our happiness."
`

	fmt.Println("=======================================")
	fmt.Println("💍   You're Invited to Our Wedding!  💍")
	fmt.Println("=======================================")
	fmt.Printf("👰 Bride: %s\n", invitation.Bride)
	fmt.Printf("🤵 Groom: %s\n", invitation.Groom)
	fmt.Println("---------------------------------------")
	fmt.Printf("📅 Date : %s\n", invitation.Date)
	fmt.Printf("⏰ Time : %s\n", invitation.Time)
	fmt.Printf("📍 Venue: %s\n", invitation.Location)
	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Println("🌸 A Special Message 🌸")
	fmt.Println(message) // 하객들에게 전하는 메시지 출력
	fmt.Println("---------------------------------------")
	fmt.Println("💖 We look forward to celebrating with you! 💖")
}
