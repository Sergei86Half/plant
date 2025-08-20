package main

import (
	"bufio" // Для чтения ввода пользователя
	"fmt"   // Для вывода в консоль
	"os"    // Для работы со стандартным вводом/выводом
	"strings" // Для обработки ввода пользователя
)

// Константы для логики игры
const waterNeededForStage = 3 // Количество поливов, необходимых для перехода на следующую стадию

// stageNames содержит названия стадий и их "визуальное" представление (эмодзи)
var stageNames = []string{
	"Стадия 1: Семечко 🌱",
	"Стадия 2: Молодое растение 🌿",
	"Стадия 3: Взрослое растение 🌳",
	"Стадия 4: Полностью выросшее! 🌻", // Финальная стадия после всех переходов
}

// Plant - структура, описывающая состояние растения
type Plant struct {
	currentStageIndex int // Текущая стадия (индекс в массиве stageNames)
	wateringsCount    int // Количество поливов на текущей стадии
}

// NewPlant создает и инициализирует новое растение
func NewPlant() *Plant {
	return &Plant{
		currentStageIndex: 0, // Начинаем с первой стадии (Семечко)
		wateringsCount:    0,
	}
}

// GetCurrentStageName возвращает название текущей стадии растения
func (p *Plant) GetCurrentStageName() string {
	if p.currentStageIndex >= len(stageNames) {
		return "Неизвестная стадия" // Этого не должно произойти при корректной логике
	}
	return stageNames[p.currentStageIndex]
}

// Water имитирует полив растения.
// Если достигнуто достаточное количество поливов, растение переходит на следующую стадию.
func (p *Plant) Water() {
	// Проверяем, достигло ли растение финальной стадии
	if p.currentStageIndex >= len(stageNames)-1 { // len(stageNames)-1 - это индекс последней стадии
		fmt.Println("✨ Растение уже полностью выросло! Больше поливать не нужно.")
		return
	}

	p.wateringsCount++
	fmt.Printf("💧 Вы полили растение! Полито: %d/%d раз\n", p.wateringsCount, waterNeededForStage)

	// Проверяем, достаточно ли поливов для перехода на следующую стадию
	if p.wateringsCount >= waterNeededForStage {
		p.currentStageIndex++ // Переходим на следующую стадию
		p.wateringsCount = 0 // Сбрасываем счетчик поливов для новой стадии

		if p.currentStageIndex < len(stageNames) {
			fmt.Printf("🎉 Растение выросло! Теперь это %s 🎉\n", p.GetCurrentStageName())
		}
	}
}

// DisplayStatus выводит текущее состояние растения в консоль (имитация UI)
func (p *Plant) DisplayStatus() {
	fmt.Println("\n--- Статус растения ---")
	fmt.Printf("Текущая стадия: %s\n", p.GetCurrentStageName())
	// Отображаем прогресс поливов только если растение ещё не полностью выросло
	if p.currentStageIndex < len(stageNames)-1 {
		fmt.Printf("Поливов до следующей стадии: %d/%d\n", p.wateringsCount, waterNeededForStage)
	} else {
		fmt.Println("Поздравляем! Ваше растение достигло финальной стадии роста!")
	}
	fmt.Println("----------------------\n")
}

func main() {
	fmt.Println("Добро пожаловать в простой симулятор роста растения!")

	plant := NewPlant()                 // Создаем новое растение
	reader := bufio.NewReader(os.Stdin) // Создаем считыватель для ввода пользователя

	for {
		plant.DisplayStatus() // Отображаем текущий статус растения

		// Проверяем, если растение полностью выросло, завершаем симуляцию
		if plant.currentStageIndex >= len(stageNames)-1 {
			fmt.Println("Игра завершена. Ваше растение полностью выросло. Ура!")
			break
		}

		fmt.Println("Нажмите ENTER, чтобы полить растение, или введите 'q' для выхода:")
		input, _ := reader.ReadString('\n') // Читаем ввод пользователя до символа новой строки

		// Обрезаем пробелы (включая символ новой строки) и проверяем, хочет ли пользователь выйти
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "q" {
			fmt.Println("Выход из симуляции. До свидания!")
			break
		}

		plant.Water() // Вызываем метод полива растения
	}
}

