package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier020(t *testing.T){v:=reflect.ValueOf(struct{Name string}{"x"});if got:=fieldByName(v,"name",true);got.IsValid(){t.Fatalf("got=%v",got)}}
