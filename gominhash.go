package gominhash


//Minhash now only for two sets
func MinHash(set1, set2 [] string) {
	//sign1 := findMin(set1)
	//sign2 := findMin(set2)
	totalElements(set1, set2)

}

func Hash(item string)int64 {
	return 0
}

func findMin(set1 []string)int64{
	min := int64(1000000000000)
	for _, elem := range set1 {
		hash_result := Hash(elem)
		if hash_result < min {
			min = hash_result
		}
	}
	return min
}

func jaccardSim(count int, set1, set2 []string) float64 {
	total := totalElements(set1, set2)
	inter := 0
	for i := 0; i < count; i++ {
		if set1[i] == set2[i] {
			inter += 1
		}
	}
	return float64(inter)/float64(len(total))

}

func totalElements(set1, set2 []string)[] string{
	total := []string{}
	for _, s1 := range set1 {
		if !findElement(total, s1) {
			total = append(total, s1)
		}
	}
	return total
}

func findElement(items []string, elem string) bool{
	for _, item := range items {
		if item == elem {
			return true
		}
	}

	return false
}