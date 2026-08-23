package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier001(t *testing.T){fn:=TypeConverter{SrcType:int(0),DstType:"",Fn:func(v interface{})(interface{},error){return "ok",nil}};m:=(Option{Converters:[]TypeConverter{fn}}).converters();if len(m)!=1{t.Fatalf("len=%d",len(m))}}
