package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier009(t *testing.T){n:=7;p:=&n;pp:=&p;if got:=indirect(reflect.ValueOf(pp));got.Kind()!=reflect.Int||got.Int()!=7{t.Fatalf("got=%v",got.Kind())}}
