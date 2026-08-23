package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier006(t *testing.T){v:=reflect.ValueOf(struct{Name string}{"alice"});got:=fieldByNameOrZeroValue(v,"Name");if !got.IsValid()||got.String()!="alice"{t.Fatalf("got=%v",got)}}
