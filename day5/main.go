package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type value struct {
	in  []string
	has []string
}

func stringInArray(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}

func main() {
	part_one_total := 0
	sep_hash := map[string]value{}
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
	re := regexp.MustCompile(`(\d+)\|(\d+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	skip_characters := 0

	for _, match := range matches {
		if len(match) == 3 {
			prev_map, _ := sep_hash[match[1]]

			prev_map.has = append(prev_map.has, match[2])
			prev_map.in = append(prev_map.in, match[1])

			sep_hash[match[1]] = prev_map
			skip_characters += len(match[0]) + 1
		}

	}

	re_comman := regexp.MustCompile(`\d+`)

	for _, match := range strings.Split(input[skip_characters:], "\n") {
		number_matches := re_comman.FindAllStringSubmatch(match, -1)

		total_numbers := len(number_matches)

		has_value := true

		for k_p, v := range number_matches {

			map_array, ok := sep_hash[v[0]]

			if !has_value {

			}

			if !ok {
				continue
			}

			wg := sync.WaitGroup{}

			for i := k_p + 1; i < total_numbers; i++ {

				wg.Add(1)
				go func() {
					defer wg.Done()
					if has_value && !stringInArray(map_array.has, number_matches[i][0]) {
						has_value = false
					}
				}()

			}

			wg.Wait()
		}

		if total_numbers > 1 {
			if has_value {
				int_value, _ := strconv.Atoi(number_matches[total_numbers/2][0])

				fmt.Println(int_value)
				part_one_total += int_value
			} else {

			}

		}

	}

	fmt.Println(part_one_total)
}
