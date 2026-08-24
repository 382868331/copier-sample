package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual

func TestGoletaCopier014(t *testing.T){src:=map[string]int{"a":1,"b":2};var dst map[string]int;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if !reflect.DeepEqual(dst,src){t.Fatalf("dst=%v",dst)}}

func TestGoletaCopier014Boundary(t *testing.T) {
 src:=map[int]string{2:"two"}
 dst:=map[int]string{1:"old"}
 if err:=Copy(&dst,src);err!=nil{t.Fatal(err)}
 if dst[2]!="two"{t.Fatalf("dst=%v",dst)}
}
