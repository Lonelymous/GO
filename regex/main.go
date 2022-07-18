package main

import (
	"fmt"
	"regexp"
)

func findImages(htm string) []string {
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

func main() {
	fmt.Printf("\n\n\n")
	strArray := []string{"a3g32gg43", "2vhh", "aaasd", "lm", "asdfbhagfalgner", "abc"}

	strArray = append(strArray, "2342523")
	strArray = append(strArray, "69420")

	isDigitsOnly := regexp.MustCompile(`^\d+$`)
	isStartsWithDigit := regexp.MustCompile(`^\d+`)
	isEndsWithDigit := regexp.MustCompile(`\d+$`)
	isLetterOnly := regexp.MustCompile(`^[a-z]+$`)
	isLetterOnlyABC := regexp.MustCompile(`^[abc]+$`)
	isLetterOnlyLM := regexp.MustCompile(`[lm]{2}`)

	for _, element := range strArray {
		fmt.Println(element, "\tisDigitsOnly: ", isDigitsOnly.MatchString(element))
		fmt.Println(element, "\tisStartsWithDigit: ", isStartsWithDigit.MatchString(element))
		fmt.Println(element, "\tisEndsWithDigit: ", isEndsWithDigit.MatchString(element))
		fmt.Println(element, "\tisLetterOnly: ", isLetterOnly.MatchString(element))
		fmt.Println(element, "\tisLetterOnlyABC: ", isLetterOnlyABC.MatchString(element))
		fmt.Println(element, "\tisLetterOnlyLM: ", isLetterOnlyLM.MatchString(element))
	}
	fmt.Printf("\n\n\n")

	// `\s+(?:19\d{2}|20[01][0-9]|2022)[-/.](?:0[1-9]|1[012])[-/.](?:0[1-9]|[12][0-9]|3[01])\b`

	isDateValid := regexp.MustCompile(`^[1-9]{3}(-|.)(0[1-9]|1[0-2])(-|.)\d{2}$`)
	fmt.Println(isDateValid.MatchString("2004-05.01"))

	fmt.Printf("\n\n\n")

	// isEmailValid := regexp.MustCompile(`^.+@.+\..+$`)
	isEmailValid := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z]+(\.[a-zA-Z]+)*`)

	fmt.Println(isEmailValid.MatchString("szabo.armin.andras@gmail.com"))
	fmt.Println(isEmailValid.MatchString("anyadseszeretregex@szar.pornhub"))

	fmt.Printf("\n\n\n")

	isLakcimValid := regexp.MustCompile(`[1-9]\d{3} (.*) (.*) (utca|tér|köz) \d+`)

	fmt.Println(isLakcimValid.MatchString("4029 Debrecen csapó utca 74"))
	fmt.Println(isLakcimValid.MatchString("4400 Nyíregyháza kossuth tér 2"))
	fmt.Println(isLakcimValid.MatchString("4021 Valami jászai köz 54"))

	fmt.Printf("\n\n\n")

	isNumberValid := regexp.MustCompile(`^(|-)(|[0-9])+(|\.|\,)\d+$`)

	fmt.Println(isNumberValid.MatchString("4029"))
	fmt.Println(isNumberValid.MatchString("0"))
	fmt.Println(isNumberValid.MatchString("0.123"))
	fmt.Println(isNumberValid.MatchString("0,123"))
	fmt.Println(isNumberValid.MatchString("-0.123"))
	fmt.Println(isNumberValid.MatchString("-.123"))
	fmt.Println(isNumberValid.MatchString("00000.123"))
	fmt.Println(isNumberValid.MatchString("-12345.123"))
	fmt.Println(isNumberValid.MatchString("000021.123"))

	fmt.Printf("\n\n\n")

	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	// if your img's are properly formed with doublequotes then use this, it's more efficient.
	// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)

	fmt.Printf("\n\n\n")
}
