package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual

func TestGoletaCopier014(t *testing.T){src:=map[string]int{"a":1,"b":2};var dst map[string]int;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if !reflect.DeepEqual(dst,src){t.Fatalf("dst=%v",dst)}}
