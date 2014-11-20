package contractor

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func NewContractorCase(i interface{}) ContractorCase {

	return ContractorCase{i}
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
				var destTempField reflect.Value

				// check if field contains a dot.
				matchDot, _ := regexp.MatchString("\\.", field)
				if matchDot == true {
					parts := strings.Split(field, ".")

					// For now only 1 sublevel.. gotta make this recursive...
					if len(parts) == 2 {
						destTempField = dest.Elem().FieldByName(parts[0])

						if destTempField.IsValid() {
							if destTempField.Type().Kind() == reflect.Struct {
								sublevel := destTempField.FieldByName(parts[1])

								if sublevel.IsValid() {
									if sublevel.CanSet() {
										destTempField = sublevel
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								continue
							}
						} else {
							continue
						}
					} else {
						continue
					}
				} else {
					destTempField = dest.Elem().FieldByName(field)
				}

				destField := destTempField
				rval := reflect.ValueOf(val)

				if destField.IsValid() {
					if destField.CanSet() {

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
								CurVal := reflect.ValueOf(val)

								if CurVal.CanInterface() {
									CurLen := CurVal.Len()

									if CurLen > 0 {
										for i := 0; i < CurLen; i++ {
											// Since it is an pointer, get the Inderect value.
											CurSlice := reflect.Indirect(reflect.ValueOf(CurVal.Index(i).Interface()))

											destField.Set(reflect.Append(destField, CurSlice))
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
							fmt.Printf("Contractor: unknown kind to set: ` %v ` . please file a request to http://github.com/donseba/contractor \n", destField.Type().Kind())
						}
					}
				}
			}
		}
	}
}

func (C *ContractorCase) Get() interface{} {
	return reflect.ValueOf(C.Case).Interface()
}

func (C *ContractorCase) CaseItem(field string) interface{} {
	t := reflect.TypeOf(C.Case)

	if t.Kind() != reflect.Ptr {
		fmt.Printf("Contractor: `CaseItem` Case must be a pointer, but got: %t", C.Case)
	}

	dest := reflect.ValueOf(C.Case)
	destField := dest.Elem().FieldByName(field)

	return destField.Interface()

}
