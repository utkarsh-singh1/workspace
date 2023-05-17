
/* _   _ _        _             _  _   
| \ | (_)_ __  (_) __ _      | || |  
|  \| | | '_ \ | |/ _` |_____| || |_ 
| |\  | | | | || | (_| |_____|__   _|
|_| \_|_|_| |_|/ |\__,_|        |_|  
             |__/                    

*/

package main

import "fmt"

func main() {
    
    // problem 1
    
    var x [5]int
   
    y := []int{1,2,3,4,5}
    
    fmt.Println(x,y)

    z := [6]int{}

    fmt.Println(z)

    for i:=0 ; i < len(z) ; i++{
        z[i] = i
    }

    fmt.Println(z)

    for i , v := range z{
        fmt.Println(i,v)
    }

    fmt.Printf("%T\n",z)
    
    // problem 2

    a := []int{1,2,3,5,6,7,8,4,9,0,67}

    for i,v :=  range a{
        fmt.Println(i,v)
    }

    fmt.Printf("%T\n",a)

    // problem 3

    b := a[:5]
    c := a[5:10]
    d := a[2:7]
    e := a[1:6]
    fmt.Println(b,c,d,e)

    // problem 4

    a = append(a, 78)

    fmt.Println(a)

    a = append(a, 53,54,55)

    fmt.Println(a)

    q := []int{22,23,34,56,57}

    a = append( a, q...)

    fmt.Println(a)

    // problem 5

    j := []int{42,43,44,45,46,47,48,49,50,51}

    // [42,43,44,48,50,51]

    j = append(j[:3],j[6:]...)

    fmt.Println(j)

    // problem 6


    r := make([]string, 50, 50)


    states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}


    for i:=0 ; i < len(states) ; i++{
        r[i] = states[i]
    }

    fmt.Println(r)

    for k:=1 ; k < len(r)+1; k++{
        fmt.Println(k,r[k-1] )
    }

    // problem 7

    s1 := []string{"Utkarsh","singh","Not","stirred"}

    s2 := []string{"Miss","Moneypenny", "Hellooo", "James"}

    s3 := [][]string{s1,s2}

    fmt.Println(s3)

    for u , v := range s3 {
        fmt.Println("record :",u)
        for w , t := range v{
            fmt.Printf("\t for index - %v value is - %v\n",w,t)
        }
    }

    // problem 8

    a1 := map[string][]string{
        "bond_james" : {`Shaken, not stirred`,`Martins`,`Women`},
        "moneypenny_miss" : {"James Bond","Literature","Cs"},
        `no_dr` : {`Being evil`, `Ice cream`, `Sunsets`},
    }

    fmt.Println(a1)
    
    a1["singh_utkarsh"] = []string{"The great Jigyasu", "Spiderman", "Daal-rice"} // solution of problem 9

    delete(a1, "bond_james")

    for o , p := range a1{
        fmt.Println("for the key:-",o)
        for _, v1 := range p{
            fmt.Printf("\t Attached value is:- %v\n",v1)
        }  
    }
    


}
