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
	str, err := ReplaceVarWithDef(testSrc1)
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
				
				
				func main(){
					var a = A{}
					a.build().self().self().self().print().var
				}`
	str, err := ReplaceVarWithDef(testSrc1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}
