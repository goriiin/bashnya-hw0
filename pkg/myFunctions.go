package pkg

const standingTimeOfTrain = 1

func FillArray(interval int, numberOfTrains int) [2]int {
	var array [2]int

	array[0] = MinWaitingTime(interval, numberOfTrains)
	array[1] = MinWaitingTime(interval, numberOfTrains) + 2*interval
	return array
}

func MinWaitingTime(interval int, numberOfTrains int) int {
	return (numberOfTrains-1)*interval + numberOfTrains*standingTimeOfTrain
}

func Swap(first *int, second *int) {
	n := *first
	*first = *second
	*second = n
}

func MinNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}
