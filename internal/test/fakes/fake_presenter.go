package fakes

import "github.com/lccezinha/twodo/internal/twodo"

type FakePresenter struct {
	Todo twodo.Todo
	Errs []twodo.ValidationError
}

func (fp *FakePresenter) PresentCreatedTodo(todo twodo.Todo) {
	fp.Todo = todo
}

func (fp *FakePresenter) PresentErrors(errs []twodo.ValidationError) {
	fp.Errs = errs
}

func NewFakePresenter() *FakePresenter {
	return &FakePresenter{}
}