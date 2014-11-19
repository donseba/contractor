package contractor

func NewContractorCaseBatch(i []*ContractorCase) *ContractorCaseBatch {

	return &ContractorCaseBatch{i}
}

type ContractorCaseBatch struct {
	Batch []*ContractorCase
}

func (C ContractorCaseBatch) Get() []interface{} {
	array := make([]interface{}, len(C.Batch))

	for key, _ := range C.Batch {
		array[key] = C.Batch[key].Get()
	}

	return array
}
