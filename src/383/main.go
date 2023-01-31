package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	dic := [26]uint8{}
	for _, b := range []byte(magazine) {
		dic[b-'a']++
	}
	for _, b := range []byte(ransomNote) {
		dic[b-'a']--
	}
	for _, count := range dic {
		if count < 0 {
			return false
		}
	}
	return true
}
