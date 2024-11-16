package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type Word struct {
	Word  string
	Count int
}

var re = regexp.MustCompile(`^[[:punct:]]+|[[:punct:]]+$`)

func PrepareWord(w string) (string, error) {
	if len(w) < 2 {
		for _, s := range w {
			if !unicode.IsLetter(s) {
				return "", errors.New("is not word")
			}
		}
	}
	return strings.ToLower(re.ReplaceAllString(w, "")), nil
}

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}

	uniqueWords := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		// проверить что слово является именно словом и отрезать от него знаки препинания, если они есть
		prepWord, err := PrepareWord(word)
		if err != nil {
			continue
		}

		_, ok := uniqueWords[prepWord]

		if !ok {
			uniqueWords[prepWord] = 1
		} else {
			uniqueWords[prepWord]++
		}
	}

	wordsSlice := make([]Word, len(uniqueWords)) // слайс отсортированных по количеству повторяющихся слов элементов

	// приведения к структуре Word
	i := 0
	for word, count := range uniqueWords {
		wordsSlice[i].Count = count
		wordsSlice[i].Word = word
		i++
	}

	length := len(wordsSlice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if wordsSlice[j].Count < wordsSlice[j+1].Count {
				wordsSlice[j], wordsSlice[j+1] = wordsSlice[j+1], wordsSlice[j]
			}

			// если слова встретились одинаковое кол-во раз, то они сортируются лексикографически
			if wordsSlice[j].Count == wordsSlice[j+1].Count {
				if wordsSlice[j].Word > wordsSlice[j+1].Word {
					wordsSlice[j], wordsSlice[j+1] = wordsSlice[j+1], wordsSlice[j]
				}
			}
		}
	}

	// вычисление на сколько элементов надо создать слайс для конечного результата
	resultSliceLen := 10
	if len(wordsSlice) < resultSliceLen {
		resultSliceLen = len(wordsSlice)
	}

	result := make([]string, resultSliceLen)
	for i := range resultSliceLen {
		result[i] = wordsSlice[i].Word
	}

	return result
}
