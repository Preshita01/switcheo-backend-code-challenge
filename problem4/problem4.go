package challenge

// Sum from 1 to n using a while loop, in increasing order
func sum_to_n_a(n int) int {
	sum := 0
	curr_num := 1

	for (curr_num <= n) {
		sum += curr_num
		curr_num += 1
	}
	
	return sum
}

// Sum from n to 1 using a while loop, in decreasing order
func sum_to_n_b(n int) int {
	sum := 0
	curr_num := 1

	for (n >= 1) {
		sum += curr_num
		curr_num -= 1
	}
	
	return sum
}

// Sum from 1 to n using a for loop
func sum_to_n_c(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += 1
	}

	return sum
}
