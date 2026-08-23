package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier004(t *testing.T){src:=struct{A int}{0};dst:=struct{A int}{9};if e:=CopyWithOption(&dst,src,Option{IgnoreEmpty:true});e!=nil||dst.A!=9{t.Fatalf("dst=%+v err=%v",dst,e)}}
