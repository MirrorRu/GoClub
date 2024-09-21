package repository_test

import (
	"fmt"
	"goclub/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	struct4filter struct {
		stringFld string
		intFld    int
	}
)

func TestNewTableSubstrFilterTest(t *testing.T) {
	t.Parallel()
	criterias := repository.FilterCriterias{{Name: "Int", Value: "777"}, {Name: "Str", Value: "text to find"}}
	filter := repository.NewTableSubstrFilter[struct4filter](criterias)
	assert.NotNil(t, filter, "Ожидается значение, отличное от nil")
	_, ok := interface{}(filter).(repository.TableFilter[struct4filter])

	assert.True(t, ok, "Провекрка соответвия интерфейсу")

}

func TestTableSubstrFilterGetSetCriterias(t *testing.T) {
	t.Parallel()
	criterias1 := repository.FilterCriterias{{Name: "Int", Value: "777"}, {Name: "Str", Value: "text to find"}}
	criterias2 := repository.FilterCriterias{{Name: "", Value: "abc"}}
	filter := repository.TableSubstrFilter[struct4filter]{}
	assert.Nil(t, filter.Criterias(), "Критерии должны быть не инициализированы")
	filter.SetCriterias(criterias1)
	assert.Equal(t, criterias1, filter.Criterias(), "Сохраненниые критерии должны совпасть с исходными")

	filter.SetCriterias(criterias2)
	assert.NotEqual(t, criterias1, filter.Criterias(), "Сохраненниые критерии НЕ должны совпасть с исходными")
}

func TestTableSubstrFilterFiltering(t *testing.T) {
	t.Parallel()
	criterias1 := repository.FilterCriterias{{Name: "Int", Value: "777"}, {Name: "Str", Value: "text to find"}}
	criterias2 := repository.FilterCriterias{{Name: "trash", Value: "abcabcabc"}}
	type testCase struct {
		name        string
		fileringObj struct4filter
		criterias   repository.FilterCriterias
		expectation bool
	}

	testCases := []testCase{
		{name: "OK by Str",
			fileringObj: struct4filter{stringFld: "Text with text to find!", intFld: 777},
			criterias:   criterias1,
			expectation: true},
		{name: "OK by Int",
			fileringObj: struct4filter{stringFld: "Any text", intFld: 1077701},
			criterias:   criterias1,
			expectation: true},
		{name: "Fail on all",
			fileringObj: struct4filter{stringFld: "Any text", intFld: 1077701},
			criterias:   criterias2,
			expectation: false},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			filter1 := repository.NewTableSubstrFilter[struct4filter](tc.criterias)
			assert.True(t, tc.expectation == filter1.Filter(tc.fileringObj), fmt.Sprintf("%s: Фильтрация созданного через New...", tc.name))

			filter2 := repository.TableSubstrFilter[struct4filter]{}
			filter2.SetCriterias(tc.criterias)
			assert.True(t, tc.expectation == filter2.Filter(tc.fileringObj), fmt.Sprintf("%s: Фильтрация созданного поэтапно", tc.name))
		})
	}
}
