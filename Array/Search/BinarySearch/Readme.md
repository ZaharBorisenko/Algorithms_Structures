
# Binary Search (Бинарный поиск)

Бинарный поиск (Binary Search) - это эффективный алгоритм поиска элемента в отсортированном списке, массиве или коллекции. Он использует стратегию разделения коллекции пополам, что позволяет быстро находить искомый элемент.
## Как работает:

Если ```arr[mid] == target``` — элемент найден.

Если ```arr[mid] > target``` — ищем в левой половине коллекции (right = mid - 1).

Если ```arr[mid] < target``` — ищем в правой половине коллекции (left = mid + 1).

Если ```left > right``` — элемент отсутствует -- завершаем поиск ```return - 1```
## Базовая реализация

```bash
func BinarySearch(arr []int, target int) int {
	//two pointers
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}

	}

	return -1
}
```


# Асимптотика
**Время**: ```O(log n)``` — так как диапазон делится пополам на каждом шаге.

**Память**: ```O(1)``` — не требует дополнительной памяти.

