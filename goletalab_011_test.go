package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier011(t *testing.T){to:=reflect.ValueOf(new(int)).Elem();ok,e:=set(to,reflect.Value{},false,nil);if e!=nil||!ok{t.Fatalf("ok=%v err=%v",ok,e)}}
