package gominhash
import
(
	"hash/adler32"
)

//Now, in case with one hash function

//Minhash now only for two sets
func MinHash(set1, set2 [] string) {
	//sign1 := findMin(set1)
	//sign2 := findMin(set2)
	totalElements(set1, set2)

}

func Hash(item string)uint32 {
	return adler32.Checksum([]byte(item))
}

func findMin(set1 []string)uint32{
	min := uint32(100000)
	for _, elem := range set1 {
		hash_result := Hash(elem)
		if hash_result < min {
			min = hash_result
		}
	}
	return min
}

func strToHash(set1 []string)[]uint32 {
	return []uint32{1}
}

func jaccardSim(count int, set1, set2 []uint32) float64 {
	//total := totalElements(set1, set2)
	inter := 0
	for i := 0; i < len(set1); i++ {
		if set1[i] == set2[i] {
			inter += 1
		}
	}
	return float64(inter)/float64(len(set1))

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