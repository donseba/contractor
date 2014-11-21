package contractor

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Initiate a new Case
func NewContractorCase(i interface{}) ContractorCase {
	return ContractorCase{i}
}

// The Case struct.
type ContractorCase struct {
	caseFiles interface{}
}

// Set the values for the current Case
func (C *ContractorCase) Set(fields map[string]interface{}) {
	if len(fields) > 0 {
		t := reflect.TypeOf(C.caseFiles)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			fmt.Printf("Contractor: Case must be a pointer, but got: %t", C.caseFiles)
		}

		if t.Kind() == reflect.Struct {
			dest := reflect.ValueOf(C.caseFiles)

			for field, val := range fields {
				var destTempField reflect.Value

				// check if field contains a dot.
				matchDot, _ := regexp.MatchString("\\.", field)
				if matchDot == true {
					destTempField = C.getNestedField(dest, field)
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

											// Append the slice item to its destination field
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
							fmt.Printf("Contractor: unknown kind to set: ` %+v ` . please file a request to http://github.com/donseba/contractor \n", destField.Type().Kind())
						}
					}
				}
			}
		}
	}
}

// Get the specific case values.
func (C *ContractorCase) Get() interface{} {
	return C.caseFiles
}

// Get a specific item inside a case.
func (C *ContractorCase) Item(field string) interface{} {
	t := reflect.TypeOf(C.caseFiles)

	if t.Kind() != reflect.Ptr {
		fmt.Printf("Contractor: `CaseItem` Case must be a pointer, but got: %t", C.caseFiles)
	}

	dest := reflect.ValueOf(C.caseFiles)
	destField := dest.Elem().FieldByName(field)

	return destField.Interface()
}

// Get the Json of an specific casefile
func (C ContractorCase) Json() []byte {
	jsonm, _ := json.Marshal(C.Get())
	return jsonm
}

// Try to reach the nested struct item value.
func (C *ContractorCase) getNestedField(dest reflect.Value, field string) reflect.Value {
	parts := strings.Split(field, ".")

	destTempField := dest.Elem().FieldByName(parts[0])

	for i := 1; i < len(parts); i++ {
		if destTempField.IsValid() {
			if destTempField.Type().Kind() == reflect.Struct {
				sublevel := destTempField.FieldByName(parts[i])

				if sublevel.IsValid() {
					if sublevel.CanSet() {
						destTempField = sublevel
					}
				}
			}
		}
	}

	return destTempField
}
