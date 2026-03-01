var (
	name string
	age int
	location string
)

var (
	name, location string
	age int
)

var name string
var age int
var location string

var (
	name string = "Michael"
	age int = 25
	location string = "Earth"
)

var (
	name = "Michael"
	age = 22
	location = "Earth"
)

var (
	name, location, age = "Michael", "Earth", 22
)

package main
import (
	"fmt"
	"time"
	"math"
	"net/http"
)

func main() {
	name := "Michael"
	age := 22
	location := "Earth"
	fmt.Printf("%s age %d from %s", name, age, location)
}

func main() {
	action := func() {

	}
	action()
}

const Pi = 3.14
const (
	StatusOK = 200
	StatusNotFound = 404
	StatusServerError = 500
)

const (
	Pi = 3.14
	E = 2.718
	Truth = false
	Big = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Printf(Pi)
}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {
	foo := map[string]interface{}{
		"Matt": 43,
	}
	timeMap(foo)
	fmt.Println(foo)
}

type Stringer interface {
	String() string
}

type fakeString struct {
	constent string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	s := &fakeString{"Ceci n'est pas un strins"}
	printString(s)
	printString("Hello, Gophers")
}

if err != nil {
	if msqlerr, ok := err.(*mysql.MySQLError); ok && msqlerr.Number == 1062 {
		log.Println("We got a MySQL duplicate :)")
	} else {
		return err
	}
}

func main() {
	cylonModel := 6
	fmt.Println(cylonModel)
}

func main() {
	name := "Caprica-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also know as %s", name, aka)
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Noew you have %g problems.", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Printf("HTTP status OK uses code: %d", http.StatusOK)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(43, 13))
}

func location(city string) (string, string) {
	var region string
	var continent string 

	switch city {
		case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s %s", region, continent)
}

client := &http.Client{}
resp, err := client.Get("http://gobootcamp.com")

// Mutability
type Artist struct {
	Name, Genre string
	Songs int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}