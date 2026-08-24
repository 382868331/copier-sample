package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual

func TestGoletaCopier016(t *testing.T){type s struct{N int};type d struct{N int64};src:=[]s{{1},{2},{3}};var dst []d;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if len(dst)!=3||dst[2].N!=3{t.Fatalf("dst=%v",dst)}}
