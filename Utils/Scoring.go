package Utils

/*
Scores known password
	total = -18
	It's present = -18
	There isn't = 0
*/
func KnownPwdScore(words []string, password string) float64 {
	found := FindExactly(words, password)
	if found {
		return -18
	}
	return 0
}

/*
Scores known reversed password
	total = -8
	It's present = -8
	There isn't = 0
*/
func KnownPwdReverseScore(words []string, password string) float64 {
	foundReverse := FindExactlyReversed(words, password)
	if foundReverse {
		return -8
	}
	return 0
}

/*
Scores password's length
	total = 20
	length <= 4 = 0
	length <= 6 = 4
	length==7 = 6
	length==8 = 10
	8 < length < 12 = 14
	11 < length < 15 = 22
	14 < length < 19 = 27
	length >=19 = 30
*/
func LengthScore(password string) float64 {
	length := len(password)
	if length <= 4 {
		return 0
	}
	if length <= 6 {
		return 4
	}
	if length == 7 {
		return 6
	}
	if length == 8 {
		return 10
	}
	if length > 8 && length < 12 {
		return 14
	}
	if length > 11 && length < 15 {
		return 22
	}
	if length > 14 && length < 19 {
		return 27
	}
	return 30
}

/*
Scores password's composition
	total = 20
	There is numbers = 5
	There is symbol = 5
	There is uppercase = 5
	There is lowercase = 5
*/
func CompositionPwdScore(password string) float64 {
	var result int
	numbers := IsThereNumber(password)
	upper := IsThereUpperCase(password)
	lower := IsThereLowerCase(password)
	symbols := IsThereSymbol(password)
	if numbers {
		result += 5
	}
	if upper {
		result += 5
	}
	if lower {
		result += 5
	}
	if symbols {
		result += 5
	}
	return float64(result)
}

/*
Scores How many different chars in relation to the length
	total = 15
	n = (different_chars*total)/total_chars
*/
func DifferentCharScore(password string) float64 {
	diffChars := HowManyDifferents(password)
	total := len(password)
	return float64((diffChars * 15) / total)
}

/*
Scores Entropy's password
	total = 35
	< 28 bits = 3
	28 - 35 bits = 8
	36 - 59 bits = 22
	60 - 80 bits = 28
	81 - 120 bits = 32
	120+ bits = 35
*/
func EntropyScore(password string) float64 {
	entropy := Entropy(password)
	if entropy <= 28 {
		return 3
	}
	if entropy > 28 && entropy <= 35 {
		return 8
	}
	if entropy > 35 && entropy <= 59 {
		return 22
	}
	if entropy > 59 && entropy <= 80 {
		return 28
	}
	if entropy > 80 && entropy <= 120 {
		return 32
	}
	return 35
}

//Grader
func Grader(words []string, password string) float64 {
	knownPwd := KnownPwdScore(words, password)
	knownPwdReverse := KnownPwdReverseScore(words, password)
	lengthScore := LengthScore(password)
	compositionPwdScore := CompositionPwdScore(password)
	differentCharScore := DifferentCharScore(password)
	entropyScore := EntropyScore(password)
	score := knownPwd + knownPwdReverse + lengthScore + compositionPwdScore + differentCharScore + entropyScore
	if score > 0 {
		return score
	}
	return 0
}