package main

import "fmt"

func getNextGap(gap float64) (float64) {
	gap = (gap*10)/13

	if(gap < 1) {
		return 1
	} else {
		return gap
	}
}

func sort(arr []int, n int) {
	var gap float64
	swapped := true

	for gap!=1 || swapped==true {
		gap = getNextGap(gap)
		intGap := int(gap)

		swapped = false

		for i:=0; i<n-intGap; i++ {
			if(arr[i] > arr[i+intGap]) {
				arr[i], arr[i+intGap] = arr[i+intGap], arr[i]

				swapped = true
			}
		}
	}
}
