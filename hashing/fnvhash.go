package hashing

import "hash/fnv"

func PerformHash(input string) uint32 {
	hashf := fnv.New32a()
	hashf.Write([]byte(input))
	return hashf.Sum32()
}
