package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier020(t *testing.T){v:=reflect.ValueOf(struct{Name string}{"x"});if got:=fieldByName(v,"name",true);got.IsValid(){t.Fatalf("got=%v",got)}}

func TestGoletaCopier020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	v:=reflect.ValueOf(struct{Name string}{"x"});got:=fieldByName(v,"name",false);if !got.IsValid()||got.String()!="x"{t.Fatalf("got=%v",got)}
}
