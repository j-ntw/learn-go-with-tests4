package main

func Sum(numbers []int) (res int) {
	for _, num := range numbers {
		res += num
	}
	return res
}
func SumAll(jobs ...[]int) (results []int) {
	for _, job := range jobs {
		results = append(results, Sum(job))
	}
	return
}

func SumAllTails(jobs ...[]int) (tails []int) {
	for _, job := range jobs {
		if len(job) == 0 {
			tails = append(tails, 0)
		} else {

			tails = append(tails, Sum(job[1:]))
		}
	}
	return tails
}
