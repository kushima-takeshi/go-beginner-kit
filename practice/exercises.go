package practice

import (
	"errors"
	"strings"
)

func Level(minutes int) int {
    // ここに実装
    if minutes >= 60 {
        return 2
    } else if minutes  <= 0 {
        return 0
    } else {
        return 1
    }
}

func Average(values []int) (float64, error) {
	if (len(values) == 0) {
		return 0, errors.New("no values to average")
	}
	sum := 0
	for _, value := range values {
		sum += value
	}
	return float64(sum) / float64(len(values)), nil
}

func Totals(records []Record) map[string]int {
    result := make(map[string]int)
    for _, record := range records {
        result[strings.TrimSpace(record.Title)] += record.Minutes
    }
    return result
}

type Record struct {
	Title   string
	Minutes int
}

func (r *Record) SetMinutes(minutes int) error {
	if minutes < 1 || minutes > 1440 {
		return errors.New("minutes must be between 1 and 1440")
	}
	r.Minutes = minutes
	return nil	
}
