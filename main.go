package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"qam-demo/pkg/modulation"
	"qam-demo/pkg/utils"
)

func main() {
	// Инициализация ГПСЧ один раз для всего приложения
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Парсинг аргументов
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "Usage: qam-demo <message> <sigma>")
		os.Exit(1)
	}

	message := flag.Arg(0)
	sigma := parseFloat(flag.Arg(1))

	// 1. Конвертация в биты
	originalBits := utils.StringToBits(message)
	originalCharCount := len(message) // Сохраняем исходную длину

	// 2. Модуляция
	symbols := modulation.Modulate16(originalBits)

	// 3. Добавление шума
	noisySymbols := modulation.AddGaussianNoise(symbols, sigma, rng)

	// 4. Демодуляция
	receivedBits := modulation.Demodulate16(noisySymbols)

	// 5. Конвертация в строку с учётом исходной длины
	receivedMessage := utils.BitsToString(receivedBits, originalCharCount)

	// 6. Подсчёт ошибок
	bitErrors := countBitErrors(originalBits, receivedBits, originalCharCount*8)

	// Вывод результатов
	fmt.Printf("Original message: %s\n", message)
	fmt.Printf("Received message: %s\n", receivedMessage)
	fmt.Printf("Bit errors: %d\n", bitErrors)
	fmt.Printf("BER: %.2f%%\n", float64(bitErrors)*100.0/float64(originalCharCount*8))
}

func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid sigma value: %s\n", s)
		os.Exit(1)
	}
	if val < 0 {
		fmt.Fprintln(os.Stderr, "Sigma must be non-negative")
		os.Exit(1)
	}
	return val
}

func countBitErrors(original, received []bool, bitLimit int) int {
	errors := 0
	for i := 0; i < bitLimit && i < len(original) && i < len(received); i++ {
		if original[i] != received[i] {
			errors++
		}
	}
	return errors
}
