package repository

import (
	"fmt"
	"strings"
)

type (
	FilterCriteriaName  string
	FilterCriteriaValue string

	FilterCriteria struct {
		Name  FilterCriteriaName
		Value FilterCriteriaValue
	}

	FilterCriterias []FilterCriteria

	TableFilter[T any] interface {
		Filter(T) (acceptable bool)
		SetCriterias(FilterCriterias)
		Criterias() FilterCriterias
	}

	TableSubstrFilter[T any] struct {
		criterias FilterCriterias
	}
)

func NewTableSubstrFilter[T any](fc FilterCriterias) *TableSubstrFilter[T] {
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

func (s *TableSubstrFilter[T]) Criterias() FilterCriterias {
	return s.criterias
}

func (s *TableSubstrFilter[T]) SetCriterias(fc FilterCriterias) {
	s.criterias = fc
}
