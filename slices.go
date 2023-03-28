// package main

// import "fmt"

// var pl = fmt.Println

// func main() {

// 	// var name[] datatype
// 	sl1 := make([]string, 6)
// 	sl1[0] = "Society"
// 	sl1[1] = "Of"
// 	sl1[2] = "the"
// 	sl1[3] = "Simulated"
// 	sl1[4] = "Universe"
// 	pl("SLice size: ", len(sl1))
// 	for i := 0; i < len(sl1); i++ {
// 		pl(sl1[i])
// 	}
// 	for _, x := range sl1 {
// 		pl(x)
// 	}
// 	sArr := [5]int {1,2,3,4,5}
// 	sl3 := sArr[0:2]
// 	pl(sl3)
// 	pl("1st 3 item of the array: ", sArr[:3])
// 	pl("1st 3 item of the array: ", sArr[2:5])
// 	sArr[0] = 10
// 	sl3[0] = 20
// 	pl("sArr now :", sArr)
// 	pl("sArr now :", sl3)

// 	sl3 = append(sl3, 12)
// 	pl("sl3 now :", sl3)

// 	sl4 := make([]string, 6)
// 	pl(sl4)
// 	pl(sl4[0])
// }
