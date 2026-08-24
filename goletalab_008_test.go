package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier008(t *testing.T){typ:=reflect.TypeOf(struct{A int}{ });_ = deepFields(typ);if got:=deepFields(typ);len(got)!=1{t.Fatalf("fields=%v",got)}}
