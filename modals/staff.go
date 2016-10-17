package modals

import "fmt"

type rect struct {
	width  int
	height int
}

// Skills string
type Skills []string

// Human struct
type Human struct {
	name   string
	weight int
}

// Person struct
type Person struct {
	Human
	name, phone string
	age         int
	Skills
	int
}

// func older(p1, p2 Person) (Person, int) {
// 	if p1.age > p2.age {
// 		return p1, p1.age - p2.age
// 	}
// 	return p2, p2.age - p1.age
// }

func (r *rect) mj() int {
	return r.width * r.height
}

// Log msg
func Log() {
	// cong := Person{Human{"lei", 120}, "cong", "136888", 35, 100}
	P2 := Person{name: "longge", age: 8, Skills: Skills{"dafeiji"}}
	P2.Skills = append(P2.Skills, "jerk off", "daren")
	P2.int = 888
	P2.Human.name = "Cong"
	fmt.Println(P2)

	r := rect{width: 10, height: 20}

	fmt.Println(r.mj())
}
