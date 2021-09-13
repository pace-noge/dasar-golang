package  main

import "fmt"

type Student struct {
    name string
    grade int
}

func (s Student) changeName1(name string) {
    fmt.Println("----> on changeName1, name changed to", name)
    s.name = name
}

func (s *Student) changeName2(name string) {
    fmt.Println("-----> on changeName2 name change to", name)
    s.name = name
}

func main() {
    var s1 = Student{"John Wick", 21}
    fmt.Println("s1 before", s1.name)

    s1.changeName1("Jason Bourne")
    fmt.Println("After changeName1", s1.name)

    s1.changeName2("ethan hunt")
    fmt.Println("s1 after changeName 2", s1.name)
}
