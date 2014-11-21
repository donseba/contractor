package contractor

import "encoding/json"

// Initiate a new batch of cases
func NewContractorCaseBatch(i []ContractorCase) ContractorCaseBatch {
	return ContractorCaseBatch{i}
}

// Batch of Cases.
type ContractorCaseBatch struct {
	Batch []ContractorCase
}

// Get all the cases that are in the batch
func (C ContractorCaseBatch) All() []interface{} {

	// Get the length of the batch
	Len := len(C.Batch)

	// Create a new array
	array := make([]interface{}, Len)

	// Loop over the items.
	for i := 0; i < Len; i++ {
		array[i] = C.Batch[i].Get()
	}

	// return the value
	return array
}

// return the json values of the batch
func (C ContractorCaseBatch) Json() []byte {
	jsonm, _ := json.Marshal(C.All())
	return jsonm
}

// Get e specific case from the batch
func (C ContractorCaseBatch) Find(i int) interface{} {
	l := len(C.Batch)

	// Check if the case is within range, if so, return it
	if l >= i {
		return C.Batch[i].Get()
	}

	// nothin to return
	return nil
}

// return the json values of a specific item
func (C ContractorCaseBatch) FindJson(i int) []byte {
	jsonm, _ := json.Marshal(C.Find(i))
	return jsonm
}
