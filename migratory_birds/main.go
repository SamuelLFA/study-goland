// https://www.hackerrank.com/challenges/migratory-birds/problem

package main

import "fmt"

type bird struct {
	id   int32
	freq int32
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	birds := []bird{}

	for _, s := range arr {
		exists := false
		for i := 0; i < len(birds); i++ {
			if birds[i].id == s {
				birds[i].freq += int32(1)
				exists = true
			}
		}
		if exists == false {
			birds = append(birds, bird{id: int32(s), freq: 1})
		}
	}
	id_most_frequent := int32(99999999)
	freq_most_frequent := int32(0)
	for _, bird := range birds {
		if bird.freq > freq_most_frequent {
			freq_most_frequent = bird.freq
			id_most_frequent = bird.id
		} else if bird.freq == freq_most_frequent && bird.id < id_most_frequent {
			freq_most_frequent = bird.freq
			id_most_frequent = bird.id
		}
	}

	return id_most_frequent
}

func main() {
	// 4
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}
