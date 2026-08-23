package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier006(t *testing.T){v:=reflect.ValueOf(struct{Name string}{"alice"});got:=fieldByNameOrZeroValue(v,"Name");if !got.IsValid()||got.String()!="alice"{t.Fatalf("got=%v",got)}}

func TestGoletaCopier006AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	v:=reflect.ValueOf(struct{Age int}{8});got:=fieldByNameOrZeroValue(v,"Age");if !got.IsValid()||got.Int()!=8{t.Fatalf("got=%v",got)}
}
