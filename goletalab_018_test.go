package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual
type goletaSetterSource018 struct{Name string}
type goletaSetterDest018 struct{got string}
func (d *goletaSetterDest018) Name(v string){d.got=v}

func TestGoletaCopier018(t *testing.T){src:=goletaSetterSource018{Name:"alice"};var dst goletaSetterDest018;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if dst.got!="alice"{t.Fatalf("got=%q",dst.got)}}
