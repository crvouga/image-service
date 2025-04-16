package uowInMemory

type UowInMemory struct {
	operations []func() error
}

func Begin() UowInMemory {
	return UowInMemory{
		operations: []func() error{},
	}
}

func (uow *UowInMemory) Rollback() error {
	for _, operation := range uow.operations {
		if err := operation(); err != nil {
			return err
		}
	}

	return nil
}

func (uow *UowInMemory) Commit() error {
	for _, operation := range uow.operations {
		if err := operation(); err != nil {
			return err
		}
	}
	return nil
}

func (uow *UowInMemory) Add(operation func() error) {
	uow.operations = append(uow.operations, operation)
}
