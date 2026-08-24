package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual

func TestGoletaCopier016(t *testing.T){type s struct{N int};type d struct{N int64};src:=[]s{{1},{2},{3}};var dst []d;if err:=Copy(&dst,src);err!=nil{t.Fatal(err)};if len(dst)!=3||dst[2].N!=3{t.Fatalf("dst=%v",dst)}}

func TestGoletaCopier016Boundary(t *testing.T) {
 type s struct{V int};type d struct{V int64}
 src:=[]s{{42}};var dst []d
 if err:=Copy(&dst,src);err!=nil{t.Fatal(err)}
 if len(dst)!=1||dst[0].V!=42{t.Fatalf("dst=%v",dst)}
}
