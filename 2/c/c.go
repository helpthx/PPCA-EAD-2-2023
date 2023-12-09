package main

import (
	"fmt"
)

// func main() {
//     for {
//         var n int
//         fmt.Scanf("%d", &n)

//         v := make([]int, n)
//         for i := 0; i < n; i++{
//             fmt.Scanf("%d", &v[i])
//         }

//         if n == 0 {
//             break
//         }

//         mostCommonNum := findMostCommon(v)
// 	    fmt.Println(mostCommonNum) // Output: 3
//     }
// }

func main() {
    var n int
    fmt.Scanf("%d", &n)
    solver(n)
}

func solver(n int) {
    if n == 0 {
        return
    }

    ans := 0
    counter := 0
    for i := 0; i < n; i++{
        var x int
        fmt.Scanf("%d", &x)
        if counter == 0{
            ans = x
            counter++
        } else {

            if ans == x {
                counter += 1
            } else {
                counter += -1
            }

        }
    }
    fmt.Printf("%d\n", ans)
    fmt.Scanf("%d", &n)
    solver(n)
}

func findMostCommon(slice []int) int {
    // Create a map to store the frequency of each number
    frequencyMap := make(map[int]int)

    // Count the frequency of each number
    for _, num := range slice {
        frequencyMap[num]++
    }

    // Find the number with the highest frequency
    mostCommonNum := 0
    maxFrequency := 0
    for num, frequency := range frequencyMap {
        if frequency > maxFrequency {
            mostCommonNum = num
            maxFrequency = frequency
        }
    }

    return mostCommonNum
}
