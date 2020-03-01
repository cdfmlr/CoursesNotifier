package wxPlatformServer

import (
	"fmt"
	"math/rand"
	"time"
)

type CoursesSerSession struct {
	verification   string
	databaseSource string
}

func (s *CoursesSerSession) GenerateVerification() {
	randI := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000) // 4位随机数
	randS4 := fmt.Sprintf("%04v", randI)                                   // 4位随机数字字符串
	s.verification = randS4
}

type VerifySerSession interface {
	GenerateVerification()
	Verify() string
	Continue(verificationCode string) string
}