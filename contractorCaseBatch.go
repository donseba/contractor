package contractor

func NewContractorCaseBatch(i []*ContractorCase) *ContractorCaseBatch {

	return &ContractorCaseBatch{i}
}

type ContractorCaseBatch struct {
	Batch []*ContractorCase
}

func (C ContractorCaseBatch) Get() []interface{} {
	Len := len(C.Batch)
	array := make([]interface{}, Len)

	for i := 0; i < Len; i++ {
		array[i] = C.Batch[i].Get()
	}

	return array
}
