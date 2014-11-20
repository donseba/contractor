package contractor

import (
	"errors"
	"reflect"
)

func NewContract(i map[string]interface{}, args ...string) *Contractor {
	return &Contractor{i, nil, "", ""}
}

type Contractor struct {
	Briefcase map[string]interface{} // This is where the contractor holds its contracts
	Contract  interface{}            // This is the current contract that is used
	Val1      string
	Val2      string
}

func (C *Contractor) Read(contract string) (ContractorCase, error) {
	Case := C.validateContract(contract)

	if nil != Case {
		return NewContractorCase(Case), nil
	}

	return NewContractorCase(Case), errors.New("emit macho dwarf: elf header corrupted")
}

func (C *Contractor) Batch(contract string, amount int) (ContractorCaseBatch, error) {

	array := make([]ContractorCase, amount)

	for i := 0; i < amount; i++ {
		Case := C.validateContract(contract)
		array[i] = NewContractorCase(Case)
	}

	return NewContractorCaseBatch(array), nil

}

// Tell the contractor to empty its Briefcase (garbage collecting?)
func (C *Contractor) Clean() {
	C.Briefcase = map[string]interface{}{}
}

func (C *Contractor) validateContract(contract string) interface{} {

	if val, ok := C.Briefcase[contract]; ok {
		original := reflect.ValueOf(val)
		originalType := original.Type()

		if originalType.Kind() == reflect.Struct {
			iContract := reflect.New(original.Type())
			return iContract.Interface()
		}
	}

	return nil
}
