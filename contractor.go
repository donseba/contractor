package contractor

import (
	"errors"
	"reflect"
)

// Dispatch a new contractor.
func NewContractor(i map[string]interface{}, args ...string) *Contractor {
	return &Contractor{i}
}

// Contractor struct.
// Where does a contractor hold it's contracts? In his briefcase!
type Contractor struct {
	Briefcase map[string]interface{}
}

// Validate the contracts and make a new case.
func (C *Contractor) Read(contract string) (ContractorCase, error) {
	Case := C.validateContract(contract)

	if nil != Case {
		return NewContractorCase(Case), nil
	}

	return NewContractorCase(Case), errors.New("Contractor: Cannot read the contracts, It does not exists, or I need some glasses")
}

// Batch allows you to create multiple cases at once.
func (C *Contractor) Batch(contract string, amount int) (ContractorCaseBatch, error) {
	// new slice of Cases
	array := make([]ContractorCase, amount)

	for i := 0; i < amount; i++ {
		Case := C.validateContract(contract)
		array[i] = NewContractorCase(Case)
	}

	return NewContractorCaseBatch(array), nil
}

// Tell the contractor to empty its Briefcase (garbage collecting?)
func (C *Contractor) Destroy() {
	C.Briefcase = map[string]interface{}{}
}

// Validate the contract and return a new instance of it.
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
