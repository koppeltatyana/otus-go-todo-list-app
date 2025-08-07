package utils

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomStr(length int) string {
	source := rand.NewSource(time.Now().UnixNano()) // Создание нового источника случайных чисел
	r := rand.New(source)                           // Создание нового генератора случайных чисел
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))] // Генерация случайного символа
	}
	return string(b)
}

func GetRandomEmail() string {
	length := rand.Intn(10)
	str := GetRandomStr(length)
	domains := []string{"email.ru", "gmail.com", "yandex.ru"}
	return str + "@" + domains[rand.Intn(len(domains))]
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// GetMaxVal - Получение максимального значения в слайсе интов
func GetMaxVal(arr []int) int {
	maxVal := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	return maxVal
}
