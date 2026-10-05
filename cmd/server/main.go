package main

import (
	"log"
	"quick/internal/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}
}
