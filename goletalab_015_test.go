package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual

func TestGoletaCopier015(t *testing.T){src:=[]int{4,7,9};var dst []int;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if !reflect.DeepEqual(dst,src){t.Fatalf("dst=%v",dst)}}

func TestGoletaCopier015Boundary(t *testing.T) {
 src:=[]string{"first"}
 var dst []string
 if err:=Copy(&dst,src);err!=nil{t.Fatal(err)}
 if !reflect.DeepEqual(dst,src){t.Fatalf("dst=%v",dst)}
}
