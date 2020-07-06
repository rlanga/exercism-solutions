package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	routineResult := make(chan FreqMap, len(s))
	result := FreqMap{}

	for _,st := range s {
		go func(st string) {
			routineResult <- Frequency(st)
		}(st)
	}
	recCount := 0
	for {
		if recCount == len(s) { break }
		val, ok := <- routineResult
		recCount++
		if !ok { break }
		for k, v := range val {
			result[k] += v
		}
	}

	close(routineResult)
	return result
}
