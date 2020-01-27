package vars

import (
	"testing"
)

func TestReplaceVarWithDef(t *testing.T) {
	testSrc1 := `package main
				 
				 import "fmt"
				 import "strings"

				 func main() {
					strings.Split(fmt.Sprintf("%d", 1),"").var
				 }
				`
	str, err := ReplaceVarWithDef("", testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}

func TestReplaceVarWithDef1(t *testing.T) {
	testSrc1 := `package main

				import "fmt"
	
				type A struct{}
				func (a A) self() A{
					return a
				}

				func (a A) print() A{
					fmt.Println("a")
					return a
				}

				func (a A) print1() (a A, b string) {
					return a, ""
				}
				
				
				func main(){
					var a = A{}
					a.self().self().print().var
				}`
	str, err := ReplaceVarWithDef("", testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}

func TestReplaceVarWithDef2(t *testing.T) {
	testSrc1 := `package main

				import "fmt"
	
				type A struct{}
				func (a A) self() A{
					return a
				}

				func (a A) print() A{
					fmt.Println("a")
					return a
				}

				func print1(a A) (a A, b string) {
					return a, ""
				}
				
				
				func main(){
					var a = A{}
					print1(a).var
				}`
	str, err := ReplaceVarWithDef("", testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}

func TestReplaceVarWithDef3(t *testing.T) {
	testSrc1 := `package main

				import "fmt"
	
				type A struct{}
				func (a A) self() A{
					return a
				}

				func (a A) print() A{
					fmt.Println("a")
					return a
				}

				func print1(a A) (a A, b string) {
					return a, ""
				}
				
				
				func main(){
					100.var
				}`
	str, err := ReplaceVarWithDef("", testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}

func TestReplaceVarWithDef4(t *testing.T) {
	testSrc1 := `package main
				 
				 import "fmt"
				 import "strings"

				 type A struct{}
				 func print() []A{
					 return nil
				 }

				 func print1() map[int]string {
					 return nil
				 }

				 func print2() func() {
					return nil
				}

				func print3() []func() {
					return nil
				}

				 func main() {
					tmp := A{}
				    tmp.self().self().String()
					fmt.Println("aa")
					tmp.var
				 }
				`
	str, err := ReplaceVarWithDef("", testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}
