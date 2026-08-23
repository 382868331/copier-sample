package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier001(t *testing.T){fn:=TypeConverter{SrcType:int(0),DstType:"",Fn:func(v interface{})(interface{},error){return "ok",nil}};m:=(Option{Converters:[]TypeConverter{fn}}).converters();if len(m)!=1{t.Fatalf("len=%d",len(m))}}

func TestGoletaCopier001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	a:=TypeConverter{SrcType:"",DstType:int(0),Fn:func(v interface{})(interface{},error){return 1,nil}};if len((Option{Converters:[]TypeConverter{a}}).converters())!=1{t.Fatal("converter missing")}
}
