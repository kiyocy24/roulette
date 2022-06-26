package helper

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func Total(m map[string]int) int {
	var total int
	for _, v := range m {
		total += v
	}

	return total
}

func Rand(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Load(filepath string) (map[string]int, error) {
	items := map[string]int{}
	_, err := os.Stat(filepath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return map[string]int{}, nil
	} else {
		b, err := ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &items)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func Lot(items map[string]int) string {
	totalWeight := Total(items)
	r := Rand(totalWeight)
	var tmp int
	for k, v := range items {
		if tmp <= r && r < tmp+v {
			return k
		}
		tmp += v
	}

	return ""
}
