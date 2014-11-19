package contractor

import (
	"fmt"
	"reflect"
)

func NewContractorCase(i interface{}) *ContractorCase {

	return &ContractorCase{i}
}

type ContractorCase struct {
	Case interface{}
}

// Set the values for the current Case
func (C *ContractorCase) Set(fields map[string]interface{}) {
	if len(fields) > 0 {
		t := reflect.TypeOf(C.Case)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			fmt.Printf("Contractor: Case must be a pointer, but got: %t", C.Case)
		}

		if t.Kind() == reflect.Struct {
			dest := reflect.ValueOf(C.Case)

			for field, val := range fields {

				destField := dest.Elem().FieldByName(field)
				rval := reflect.ValueOf(val)

				switch destField.Type().Kind() {
				case reflect.String:
					destField.SetString(rval.String())

				case reflect.Bool:
					destField.Set(reflect.ValueOf(val))

				case reflect.Int, reflect.Int32, reflect.Int64:

					destField.SetInt(reflect.ValueOf(val).Int())

				case reflect.Float32, reflect.Float64:
					destField.Set(reflect.ValueOf(val))

				case reflect.Slice:

					switch reflect.ValueOf(val).Type().Kind() {
					case reflect.Slice:
						sliceVal := reflect.ValueOf(val)

						if sliceVal.CanInterface() {
							sliceLen := sliceVal.Len()

							if sliceLen > 0 {
								SliceInd := reflect.Indirect(reflect.ValueOf(val))

								for i := 0; i < sliceLen; i++ {
									destField.Set(reflect.Append(destField, reflect.ValueOf(SliceInd.Index(i).Interface())))
								}
							}
						}
					}
				case reflect.Struct, reflect.Ptr:

					switch reflect.ValueOf(val).Type().Kind() {
					case reflect.Struct:
						destField.Set(reflect.ValueOf(val))
					case reflect.Ptr:
						destField.Set(reflect.ValueOf(val))
					}
				default:
					fmt.Println("unknown kind to set: ")
					fmt.Println(destField.Type().Kind())
				}
			}
		}

	} else {
		fmt.Println("Cannot set empty values.")
	}
}

func (C *ContractorCase) Get() interface{} {
	return reflect.ValueOf(C.Case).Interface()
}

func (C *ContractorCase) CaseItem(field string) interface{} {
	t := reflect.TypeOf(C.Case)

	if t.Kind() != reflect.Ptr {
		fmt.Printf("Contractor CaseItem: Case must be a pointer, but got: %t", C.Case)
	}

	dest := reflect.ValueOf(C.Case)
	destField := dest.Elem().FieldByName(field)

	return destField.Interface()

}
