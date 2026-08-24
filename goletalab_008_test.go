package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier008(t *testing.T){typ:=reflect.TypeOf(struct{A int}{ });_ = deepFields(typ);if got:=deepFields(typ);len(got)!=1{t.Fatalf("fields=%v",got)}}

func TestGoletaCopier008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	typ:=reflect.TypeOf(struct{A int;B string}{});_ = deepFields(typ);if len(deepFields(typ))!=2{t.Fatal("cache miss")}
}
