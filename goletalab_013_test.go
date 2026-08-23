package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier013(t *testing.T){to:=reflect.ValueOf(new(string)).Elem();to.SetString("old");from:=reflect.ValueOf(1);pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,e:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return nil,nil}}});if e!=nil||!ok||to.String()!=""{t.Fatalf("to=%q",to.String())}}
