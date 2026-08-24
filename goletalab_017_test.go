package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual
type goletaMethodSource017 struct{Base int}
func (s goletaMethodSource017) Score()int{return s.Base*2}
type goletaMethodDest017 struct{Score int}

func TestGoletaCopier017(t *testing.T){src:=goletaMethodSource017{Base:6};var dst goletaMethodDest017;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if dst.Score!=12{t.Fatalf("score=%d",dst.Score)}}

func TestGoletaCopier017Boundary(t *testing.T) {
 src:=goletaMethodSource017{Base:-3}
 var dst goletaMethodDest017
 if err:=Copy(&dst,src);err!=nil{t.Fatal(err)}
 if dst.Score!=-6{t.Fatalf("score=%d",dst.Score)}
}
