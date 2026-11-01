package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	// Проверяем, что функция создаёт слайс указанного размера.
	data := generateRandomElements(5)

	if len(data) != 5 {
		t.Errorf("ожидали 5 элементов, получили %d", len(data))
	}

	// Проверяем, что все сгенерированные числа положительные.
	for _, value := range data {
		if value <= 0 {
			t.Errorf("ожидали положительное число, получили %d", value)
		}
	}

	// Проверяем крайний случай: размер слайса равен нулю.
	emptyData := generateRandomElements(0)

	if len(emptyData) != 0 {
		t.Errorf(
			"для size = 0 ожидали пустой слайс, получили %d элементов",
			len(emptyData),
		)
	}
}

func TestMaximum(t *testing.T) {
	// Проверяем обычный слайс из нескольких элементов.
	data := []int{5, 17, 3, 42, 9}

	result := maximum(data)

	if result != 42 {
		t.Errorf("ожидали максимум 42, получили %d", result)
	}

	// Проверяем слайс, содержащий только один элемент.
	oneElement := []int{15}

	result = maximum(oneElement)

	if result != 15 {
		t.Errorf("ожидали максимум 15, получили %d", result)
	}

	// Проверяем крайний случай: пустой слайс.
	emptyData := []int{}

	result = maximum(emptyData)

	if result != 0 {
		t.Errorf("для пустого слайса ожидали 0, получили %d", result)
	}
}

func TestMaxChunks(t *testing.T) {
	// Проверяем обычный случай: количество элементов равно количеству частей.
	data := []int{5, 17, 3, 42, 9, 11, 8, 25}

	result := maxChunks(data)

	if result != 42 {
		t.Errorf("ожидали максимум 42, получили %d", result)
	}

	// Проверяем, что остаток от деления слайса на части не теряется.
	// Максимальное значение специально помещено в последний элемент.
	dataWithRemainder := []int{1, 2, 3, 4, 5, 6, 7, 8, 100}

	result = maxChunks(dataWithRemainder)

	if result != 100 {
		t.Errorf("ожидали максимум 100, получили %d", result)
	}

	// Проверяем случай, когда элементов меньше, чем CHUNKS.
	smallData := []int{3, 50, 7}

	result = maxChunks(smallData)

	if result != 50 {
		t.Errorf("ожидали максимум 50, получили %d", result)
	}

	// Проверяем крайний случай: пустой слайс.
	emptyData := []int{}

	result = maxChunks(emptyData)

	if result != 0 {
		t.Errorf("для пустого слайса ожидали 0, получили %d", result)
	}
}
