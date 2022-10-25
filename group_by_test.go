package slice

import (
	"testing"
)

func TestGroupByKey(t *testing.T) {
	type user struct {
		age  int
		name string
	}
	slice := make([]*user, 0)
	slice = append(slice, &user{age: 10, name: "张飞"})
	slice = append(slice, &user{age: 10, name: "关羽"})
	slice = append(slice, &user{age: 20, name: "刘备"})
	slice = append(slice, &user{age: 250, name: "曹操"})

	r := GroupByKey(slice, func(index int, item *user) int {
		return item.age
	})
	t.Logf("%#v", r)

}

func TestGroupByKeyThenCount(t *testing.T) {
	type user struct {
		age  int
		name string
	}
	slice := make([]*user, 0)
	slice = append(slice, &user{age: 10, name: "张飞"})
	slice = append(slice, &user{age: 10, name: "关羽"})
	slice = append(slice, &user{age: 20, name: "刘备"})
	slice = append(slice, &user{age: 250, name: "曹操"})

	r := GroupByKeyThenCount(slice, func(index int, item *user) int {
		return item.age
	})
	t.Logf("%#v", r)
}

func TestGroupByKeyThenOrderByCount(t *testing.T) {
	type user struct {
		age  int
		name string
	}
	slice := make([]*user, 0)
	slice = append(slice, &user{age: 10, name: "张飞"})
	slice = append(slice, &user{age: 10, name: "关羽"})
	slice = append(slice, &user{age: 20, name: "刘备"})
	slice = append(slice, &user{age: 250, name: "曹操"})

	r := GroupByKeyThenOrderByCount(slice, func(index int, item *user) int {
		return item.age
	})
	t.Logf("%#v", r)
}
