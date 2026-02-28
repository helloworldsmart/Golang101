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

type Bootcamp struct {
	Lat float64
	Lon float64
	Date time.Time
}

func main() {
	fmt.Println(Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
		Date: time.Now()
	})
}

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{x: 1}
	s = Point{}
)

func main() {
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495339
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s, location(%f, %f)",
	event.Date, event.Lat, event.Lon)
}

x := new(int)

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(*x == *y)
}

type User struct {
	Id int
	Name string
	Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{
		User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}

	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
		p.Id, p.Name, p.location, p.GameId
	)
	p.Id = 11
	fmt.Printf("%+v", p)
}