package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier003(t *testing.T){src:=struct{A int}{7};dst:=struct{A int}{};if e:=Copy(&dst,&src);e!=nil||dst.A!=7{t.Fatalf("dst=%+v err=%v",dst,e)}}
