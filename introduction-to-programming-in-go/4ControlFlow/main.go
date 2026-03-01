package main

import (
	"fmt"
	"math"
	"time"
)

if answer != 42 {
	return "Wrong answer"
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

if err := foo(); err != nil {
	panic(err)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim 
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// FOR Loop

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}
}

func main() {
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	}
}

// Exercise on For Loops
package main
import "fmt"
import "strconv"
import "encoding/json"

func GetSum(array [] int) int {
  //write code here
  sum := 0
  
  for _, v := range array {
    sum += v
  }
  return sum; //return the sum here
}


// ERROR
package main
import "fmt"
import "strconv"
import "encoding/json"
import "math/rand"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func GetResult([] string) string {
 //insert code here
  
  // for _, v := range users {
  //   coins -= 1
  //   distribution[v] = 1
  // }
  
  for coins > 0 {
    for _, v := range users {
    i := rand.Intn(10)
    coins -= i
    distribution[v] = 1 + i
    
    }
  }

  fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
  return strconv.Itoa(distribution["Matthew"])+" "+strconv.Itoa(distribution["Sarah"])+" "+strconv.Itoa(distribution["Augustus"])+" "+strconv.Itoa(distribution["Heidi"])+" "+strconv.Itoa(distribution["Emilie"])+" "+strconv.Itoa(distribution["Peter"])+" "+strconv.Itoa(distribution["Giana"])+" "+strconv.Itoa(distribution["Adriano"])+" "+strconv.Itoa(distribution["Aaron"])+" "+strconv.Itoa(distribution["Elizabeth"])
}


// Succeeded
package main
import "fmt"
import "strconv"
import "encoding/json"
import "strings"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func GetResult([] string) string {
 //insert code here

 vewelCoins := map[rune]int {
   'a': 1,
   'e': 1,
   'i': 2,
   'o': 3,
   'u': 4,
 }
 for _, name := range users {
   count := 0
   for _, ch := range strings.ToLower(name) {
     count += vewelCoins[ch]
   }
   if count > 10 {
     count = 10
   }
   distribution[name] = count
   coins -= count
 }

  fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
  return strconv.Itoa(distribution["Matthew"])+" "+strconv.Itoa(distribution["Sarah"])+" "+strconv.Itoa(distribution["Augustus"])+" "+strconv.Itoa(distribution["Heidi"])+" "+strconv.Itoa(distribution["Emilie"])+" "+strconv.Itoa(distribution["Peter"])+" "+strconv.Itoa(distribution["Giana"])+" "+strconv.Itoa(distribution["Adriano"])+" "+strconv.Itoa(distribution["Aaron"])+" "+strconv.Itoa(distribution["Elizabeth"])
}
