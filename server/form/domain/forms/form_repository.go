package forms

import (
	"fmt"
	"github.com/ashkan-maleki/go-form-builder/server/pkg/models/lists"
)

type FormRepository struct {
}

func (repo FormRepository) Get(Id string) (Form, error) {
	fmt.Println(Id)
	return Form{}, nil
}

func (repo FormRepository) All(pag lists.Pagination) ([]Form, error) {
	fmt.Println(pag)
	return []Form{}, nil
}

func (repo FormRepository) Create(form Form) error {
	fmt.Println(form)
	return nil
}

func (repo FormRepository) Update(Id string, form Form) error {
	fmt.Println(Id, form)
	return nil
}

func (repo FormRepository) Delete(Id string) error {
	fmt.Println(Id)
	return nil
}
