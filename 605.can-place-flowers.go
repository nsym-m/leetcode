package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	l := len(flowerbed)
	count := 0
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}
		if i == 0 {
			if (l > 1 && flowerbed[i+1] != 1) ||
				l == 1 {
				flowerbed[i] = 1
				count++
			}
			continue
		}
		if i+1 == l {
			if flowerbed[i-1] != 1 {
				flowerbed[i] = 1
				count++
			}
			continue
		}
		if flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
			flowerbed[i] = 1
			count++
		}
	}

	return count >= n
}
