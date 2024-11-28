package infoblox

import "fmt"

func numberOfAnagram(s, t string) int {
	freqt := make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		freqt[t[i]]++
	}
	fmt.Printf("freqt=%v\n", freqt)
	freqs := make(map[byte]int, 0)

	totalanagrams := 0
	for i := 0; i < len(t); i++ {
		freqs[s[i]]++
	}
	if isAnagramMap(freqt, freqs) {
		totalanagrams++
	}

	for i := len(t); i < len(s); i++ {
		outinded := i - len(t)
		outbyte := s[outinded]
		if _, ok := freqs[outbyte]; ok {
			freqs[outbyte]--
			if freqs[outbyte] == 0 {
				delete(freqs, outbyte)
			}
		}
		inbyte := s[i]
		freqs[inbyte]++
		if isAnagramMap(freqt, freqs) {
			totalanagrams++
		}
	}

	return totalanagrams

}

func isAnagramMap(freqt map[byte]int, freqs map[byte]int) bool {
	for key, val := range freqs {
		if freqt[key] != val {
			return false
		}
	}
	return true
}
