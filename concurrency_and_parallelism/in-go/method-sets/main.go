package main

import "fmt"

type user1 struct {
	Name string
	Age  int
}

func (u user1) Details1() string {

	return fmt.Sprintf("Name and Age of this user is %s and %d\n", u.Name, u.Age)

}

func (u *user1) Details2() string {

	return fmt.Sprintf("Name and Age of this gentleman is %s and  %v\n", u.Name, u.Age)

}

type user interface {
	Details1() string
	Details2() string
}

func Details(u user) {

	fmt.Printf("There are two output:-  %s and %s \n", u.Details1(),u.Details2())

}

func main() {
 
	u1 := user1{

		Name: "Jha",
		Age:  34,
	}

	u2 := &user1{
		Name: "hello",
		Age:  89,
	}

	fmt.Println(u1)

	fmt.Println(&u1)

	fmt.Println(u2.Details1(), u2.Details2())

	fmt.Println(u1.Details1(), u1.Details2())

	Details(u2)
}

// Method set of an Type T must be set of methods defined with receiver type T. Like u1 can access Both Details1() and Details2().
// And Method set of Type *T must be set of methods defined with receiver type T or *T. Means value defined with type *T is also can access the method set of T. Like above u2 with *user1 can acces Details1() which have user1 as receiver type.


// Methos set determines the interface that type implements.

// This whole fiasco of using pointer type values if the receiver type is *T is fully explained in the method-set.org read it to get full clearity. And more examples will be there. 
