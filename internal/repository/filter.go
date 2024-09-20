package repository

import (
	"fmt"
	"strings"
)

type (
	FilterCriteriaName  string
	FilterCriteriaValue string

	FilerCriteria struct {
		Name  FilterCriteriaName
		Value FilterCriteriaValue
	}

	FilerCriterias []FilerCriteria

	TableFilter[T any] interface {
		Filter(T) (acceptable bool)
		SetCriterias(FilerCriterias)
		Criterias() FilerCriterias
	}

	TableSubstrFilter[T any] struct {
		criterias FilerCriterias
	}
)

func NewTableSubstrFilter[T any](fc FilerCriterias) *TableSubstrFilter[T] {
	return &TableSubstrFilter[T]{
		criterias: fc,
	}
}

func (s *TableSubstrFilter[T]) Filter(x T) (acceptable bool) {
	data := fmt.Sprintf("%v", x)
	for _, v := range s.criterias {
		if strings.Contains(data, string(v.Value)) {
			return true
		}
	}
	return false
}

func (s *TableSubstrFilter[T]) Criterias() FilerCriterias {
	return s.criterias
}

func (s *TableSubstrFilter[T]) SetCriterias(fc FilerCriterias) {
	s.criterias = fc
}
