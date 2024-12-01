// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedlist

import (
	"cmp"
	// "encoding/json"
	"slices"
	"strings"
	"testing"
)

func TestListNew(t *testing.T) {
	list1 := New[int]()

	if actualValue := list1.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New[int]()
	list2.Prepend(1, 2)

	if actualValue := list2.GetSize(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(2); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListAdd(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListAppendAndPrepend(t *testing.T) {
	list := New[string]()
	list.Append("b")
	list.Prepend("a")
	list.Append("c")
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListRemove(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue != "" || ok {
		t.Errorf("Got %v expected %v", actualValue, "")
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.GetSize(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != "" || ok {
		t.Errorf("Got %v expected %v", actualValue, "")
	}
	list.Remove(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListSwap(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	list.Swap(0, 1)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListSort(t *testing.T) {
	list := New[string]()
	list.Sort(cmp.Compare[string])
	list.Append("e", "f", "g", "a", "b", "c", "d")
	list.Sort(cmp.Compare[string])
	for i := 1; i < list.GetSize(); i++ {
		a, _ := list.Get(i - 1)
		b, _ := list.Get(i)
		if a > b {
			t.Errorf("Not sorted! %s > %s", a, b)
		}
	}
}

func TestListClear(t *testing.T) {
	list := New[string]()
	list.Append("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.GetSize(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListContains(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	if actualValue := list.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains(""); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Contains("b"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	list.Clear()
	if actualValue := list.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Contains("b"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestListValues(t *testing.T) {
	list := New[string]()
	list.Append("a")
	list.Append("b", "c")
	if actualValue, expectedValue := list.GetAllNode(), []string{"a", "b", "c"}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

// func TestListIndexOf(t *testing.T) {
// 	list := New[string]()
//
// 	expectedIndex := -1
// 	if index := list.IndexOf("a"); index != expectedIndex {
// 		t.Errorf("Got %v expected %v", index, expectedIndex)
// 	}
//
// 	list.Add("a")
// 	list.Add("b", "c")
//
// 	expectedIndex = 0
// 	if index := list.IndexOf("a"); index != expectedIndex {
// 		t.Errorf("Got %v expected %v", index, expectedIndex)
// 	}
//
// 	expectedIndex = 1
// 	if index := list.IndexOf("b"); index != expectedIndex {
// 		t.Errorf("Got %v expected %v", index, expectedIndex)
// 	}
//
// 	expectedIndex = 2
// 	if index := list.IndexOf("c"); index != expectedIndex {
// 		t.Errorf("Got %v expected %v", index, expectedIndex)
// 	}
// }

func TestListInsert(t *testing.T) {
	list := New[string]()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.GetSize(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := strings.Join(list.GetAllNode(), ""), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListSet(t *testing.T) {
	list := New[string]()
  list.Append("c", "d", "k")
	list.UpdateNodeValue(0, "a")
	list.UpdateNodeValue(1, "b")
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.UpdateNodeValue(2, "c") // append
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.UpdateNodeValue(4, "d")  // ignore
	list.UpdateNodeValue(1, "bb") // update
	if actualValue := list.GetSize(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := list.GetAllNode(), []string{"a", "bb", "c"}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

// func TestListEach(t *testing.T) {
// 	list := New[string]()
// 	list.Append("a", "b", "c")
// 	list.Each(func(index int, value string) {
// 		switch index {
// 		case 0:
// 			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		case 1:
// 			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		case 2:
// 			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		default:
// 			t.Errorf("Too many")
// 		}
// 	})
// }

// func TestListMap(t *testing.T) {
// 	list := New[string]()
// 	list.Append("a", "b", "c")
// 	mappedList := list.Map(func(index int, value string) string {
// 		return "mapped: " + value
// 	})
// 	if actualValue, _ := mappedList.Get(0); actualValue != "mapped: a" {
// 		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
// 	}
// 	if actualValue, _ := mappedList.Get(1); actualValue != "mapped: b" {
// 		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
// 	}
// 	if actualValue, _ := mappedList.Get(2); actualValue != "mapped: c" {
// 		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
// 	}
// 	if mappedList.Size() != 3 {
// 		t.Errorf("Got %v expected %v", mappedList.Size(), 3)
// 	}
// }

// func TestListSelect(t *testing.T) {
// 	list := New[string]()
// 	list.Append("a", "b", "c")
// 	selectedList := list.Select(func(index int, value string) bool {
// 		return value >= "a" && value <= "b"
// 	})
// 	if actualValue, _ := selectedList.Get(0); actualValue != "a" {
// 		t.Errorf("Got %v expected %v", actualValue, "value: a")
// 	}
// 	if actualValue, _ := selectedList.Get(1); actualValue != "b" {
// 		t.Errorf("Got %v expected %v", actualValue, "value: b")
// 	}
// 	if selectedList.Size() != 2 {
// 		t.Errorf("Got %v expected %v", selectedList.Size(), 3)
// 	}
// }

// func TestListAny(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
// 	any := list.Any(func(index int, value string) bool {
// 		return value == "c"
// 	})
// 	if any != true {
// 		t.Errorf("Got %v expected %v", any, true)
// 	}
// 	any = list.Any(func(index int, value string) bool {
// 		return value == "x"
// 	})
// 	if any != false {
// 		t.Errorf("Got %v expected %v", any, false)
// 	}
// }
// func TestListAll(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
// 	all := list.All(func(index int, value string) bool {
// 		return value >= "a" && value <= "c"
// 	})
// 	if all != true {
// 		t.Errorf("Got %v expected %v", all, true)
// 	}
// 	all = list.All(func(index int, value string) bool {
// 		return value >= "a" && value <= "b"
// 	})
// 	if all != false {
// 		t.Errorf("Got %v expected %v", all, false)
// 	}
// }
// func TestListFind(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
// 	foundIndex, foundValue := list.Find(func(index int, value string) bool {
// 		return value == "c"
// 	})
// 	if foundValue != "c" || foundIndex != 2 {
// 		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 2)
// 	}
// 	foundIndex, foundValue = list.Find(func(index int, value string) bool {
// 		return value == "x"
// 	})
// 	if foundValue != "" || foundIndex != -1 {
// 		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
// 	}
// }
// func TestListChaining(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
// 	chainedList := list.Select(func(index int, value string) bool {
// 		return value > "a"
// 	}).Map(func(index int, value string) string {
// 		return value + value
// 	})
// 	if chainedList.Size() != 2 {
// 		t.Errorf("Got %v expected %v", chainedList.Size(), 2)
// 	}
// 	if actualValue, ok := chainedList.Get(0); actualValue != "bb" || !ok {
// 		t.Errorf("Got %v expected %v", actualValue, "b")
// 	}
// 	if actualValue, ok := chainedList.Get(1); actualValue != "cc" || !ok {
// 		t.Errorf("Got %v expected %v", actualValue, "c")
// 	}
// }
//
// func TestListIteratorNextOnEmpty(t *testing.T) {
// 	list := New[string]()
// 	it := list.Iterator()
// 	for it.Next() {
// 		t.Errorf("Shouldn't iterate on empty list")
// 	}
// }
//
// func TestListIteratorNext(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
// 	it := list.Iterator()
// 	count := 0
// 	for it.Next() {
// 		count++
// 		index := it.Index()
// 		value := it.Value()
// 		switch index {
// 		case 0:
// 			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		case 1:
// 			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		case 2:
// 			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
// 				t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 			}
// 		default:
// 			t.Errorf("Too many")
// 		}
// 	}
// 	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
// 		t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 	}
// }
//
// func TestListIteratorBegin(t *testing.T) {
// 	list := New[string]()
// 	it := list.Iterator()
// 	it.Begin()
// 	list.Add("a", "b", "c")
// 	for it.Next() {
// 	}
// 	it.Begin()
// 	it.Next()
// 	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
// 		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
// 	}
// }
//
// func TestListIteratorFirst(t *testing.T) {
// 	list := New[string]()
// 	it := list.Iterator()
// 	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
// 		t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 	}
// 	list.Add("a", "b", "c")
// 	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
// 		t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 	}
// 	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
// 		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
// 	}
// }
//
// func TestListIteratorNextTo(t *testing.T) {
// 	// Sample seek function, i.e. string starting with "b"
// 	seek := func(index int, value string) bool {
// 		return strings.HasSuffix(value, "b")
// 	}
//
// 	// NextTo (empty)
// 	{
// 		list := New[string]()
// 		it := list.Iterator()
// 		for it.NextTo(seek) {
// 			t.Errorf("Shouldn't iterate on empty list")
// 		}
// 	}
//
// 	// NextTo (not found)
// 	{
// 		list := New[string]()
// 		list.Add("xx", "yy")
// 		it := list.Iterator()
// 		for it.NextTo(seek) {
// 			t.Errorf("Shouldn't iterate on empty list")
// 		}
// 	}
//
// 	// NextTo (found)
// 	{
// 		list := New[string]()
// 		list.Add("aa", "bb", "cc")
// 		it := list.Iterator()
// 		it.Begin()
// 		if !it.NextTo(seek) {
// 			t.Errorf("Shouldn't iterate on empty list")
// 		}
// 		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
// 			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
// 		}
// 		if !it.Next() {
// 			t.Errorf("Should go to first element")
// 		}
// 		if index, value := it.Index(), it.Value(); index != 2 || value != "cc" {
// 			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
// 		}
// 		if it.Next() {
// 			t.Errorf("Should not go past last element")
// 		}
// 	}
// }
//
// func TestListSerialization(t *testing.T) {
// 	list := New[string]()
// 	list.Add("a", "b", "c")
//
// 	var err error
// 	assert := func() {
// 		if actualValue, expectedValue := list.Values(), []string{"a", "b", "c"}; !slices.Equal(actualValue, expectedValue) {
// 			t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 		}
// 		if actualValue, expectedValue := list.Size(), 3; actualValue != expectedValue {
// 			t.Errorf("Got %v expected %v", actualValue, expectedValue)
// 		}
// 		if err != nil {
// 			t.Errorf("Got error %v", err)
// 		}
// 	}
//
// 	assert()
//
// 	bytes, err := list.ToJSON()
// 	assert()
//
// 	err = list.FromJSON(bytes)
// 	assert()
//
// 	bytes, err = json.Marshal([]any{"a", "b", "c", list})
// 	if err != nil {
// 		t.Errorf("Got error %v", err)
// 	}
//
// 	err = json.Unmarshal([]byte(`["a","b","c"]`), &list)
// 	if err != nil {
// 		t.Errorf("Got error %v", err)
// 	}
// 	assert()
// }
//
// func TestListString(t *testing.T) {
// 	c := New[int]()
// 	c.Add(1)
// 	if !strings.HasPrefix(c.String(), "SinglyLinkedList") {
// 		t.Errorf("String should start with container name")
// 	}
// }
//
func benchmarkGet(b *testing.B, list *LinkedList[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *LinkedList[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Append(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *LinkedList[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

func BenchmarkSinglyLinkedListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

