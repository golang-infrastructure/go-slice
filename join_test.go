package slice

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoGroup(t *testing.T) {
	type User struct {
		Name string
		IP   string
	}
	sliceA := make([]*User, 0)
	sliceA = append(sliceA, &User{Name: "李四", IP: "2.2.2.2"})
	sliceA = append(sliceA, &User{Name: "Tom", IP: "3.3.3.3"})
	sliceA = append(sliceA, &User{Name: "Sum", IP: "3.3.3.3"})

	type IpLocation struct {
		IP       string
		Location string
	}
	sliceB := make([]*IpLocation, 0)
	sliceB = append(sliceB, &IpLocation{Location: "中国台湾", IP: "1.1.1.1"})
	sliceB = append(sliceB, &IpLocation{Location: "美国硅谷", IP: "3.3.3.3"})
	sliceB = append(sliceB, &IpLocation{Location: "美国加州", IP: "3.3.3.3"})

	r := CoGroup(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))
}

func TestFoldByKeyFunc(t *testing.T) {
	type User struct {
		Name  string
		Money int
	}
	slice := make([]*User, 0)
	slice = append(slice, &User{Name: "李四", Money: 100})
	slice = append(slice, &User{Name: "李四", Money: 900})
	slice = append(slice, &User{Name: "Tom", Money: 200})
	slice = append(slice, &User{Name: "Tom", Money: 600})
	slice = append(slice, &User{Name: "Sum", Money: 300})
	slice = append(slice, &User{Name: "Sum", Money: 200})

	r := FoldByKeyFunc(slice, func(index int, item *User) string {
		return item.Name
	}, func(source *User, destination *User) *User {
		destination.Money += source.Money
		return destination
	})
	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))
}

func TestFullJoinByKeyFunc(t *testing.T) {
	type User struct {
		Name string
		IP   string
	}
	sliceA := make([]*User, 0)
	sliceA = append(sliceA, &User{Name: "李四", IP: "2.2.2.2"})
	sliceA = append(sliceA, &User{Name: "Tom", IP: "3.3.3.3"})
	sliceA = append(sliceA, &User{Name: "Sum", IP: "3.3.3.3"})

	type IpLocation struct {
		IP       string
		Location string
	}
	sliceB := make([]*IpLocation, 0)
	sliceB = append(sliceB, &IpLocation{Location: "中国台湾", IP: "1.1.1.1"})
	sliceB = append(sliceB, &IpLocation{Location: "美国硅谷", IP: "3.3.3.3"})
	sliceB = append(sliceB, &IpLocation{Location: "美国加州", IP: "3.3.3.3"})

	r := FullJoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))

	flat := FromMapValueFlat(r)
	t.Log(flat)
}

func TestJoinByKeyFunc(t *testing.T) {
	type User struct {
		Name string
		IP   string
	}
	sliceA := make([]*User, 0)
	sliceA = append(sliceA, &User{Name: "张三", IP: "1.1.1.1"})
	sliceA = append(sliceA, &User{Name: "李四", IP: "2.2.2.2"})
	sliceA = append(sliceA, &User{Name: "Tom", IP: "3.3.3.3"})
	sliceA = append(sliceA, &User{Name: "Sum", IP: "3.3.3.3"})

	type IpLocation struct {
		IP       string
		Location string
	}
	sliceB := make([]*IpLocation, 0)
	sliceB = append(sliceB, &IpLocation{Location: "中国北京", IP: "2.2.2.2"})
	sliceB = append(sliceB, &IpLocation{Location: "中国台湾", IP: "1.1.1.1"})
	sliceB = append(sliceB, &IpLocation{Location: "美国硅谷", IP: "3.3.3.3"})
	sliceB = append(sliceB, &IpLocation{Location: "美国加州", IP: "3.3.3.3"})

	r := JoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))

	flat := FromMapValueFlat(r)
	t.Log(flat)

	// ------------------------------------------------ ---------------------------------------------------------------------

	r = JoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err = json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))

	flat = FromMapValueFlat(r)
	t.Log(flat)
}

func TestLeftJoinByKeyFunc(t *testing.T) {
	type User struct {
		Name string
		IP   string
	}
	sliceA := make([]*User, 0)
	sliceA = append(sliceA, &User{Name: "张三", IP: "1.1.1.1"})
	sliceA = append(sliceA, &User{Name: "李四", IP: "2.2.2.2"})
	sliceA = append(sliceA, &User{Name: "Tom", IP: "3.3.3.3"})
	sliceA = append(sliceA, &User{Name: "Sum", IP: "3.3.3.3"})

	type IpLocation struct {
		IP       string
		Location string
	}
	sliceB := make([]*IpLocation, 0)
	sliceB = append(sliceB, &IpLocation{Location: "中国北京", IP: "2.2.2.2"})
	sliceB = append(sliceB, &IpLocation{Location: "美国硅谷", IP: "3.3.3.3"})

	r := LeftJoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Log(string(marshal))

	flat := FromMapValueFlat(r)
	t.Log(flat)

	// ------------------------------------------------ ---------------------------------------------------------------------

	r = LeftJoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err = json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))

	flat = FromMapValueFlat(r)
	t.Log(flat)
}

func TestRightJoinByKeyFunc(t *testing.T) {
	type User struct {
		Name string
		IP   string
	}
	sliceA := make([]*User, 0)
	sliceA = append(sliceA, &User{Name: "李四", IP: "2.2.2.2"})
	sliceA = append(sliceA, &User{Name: "Tom", IP: "3.3.3.3"})
	sliceA = append(sliceA, &User{Name: "Sum", IP: "3.3.3.3"})

	type IpLocation struct {
		IP       string
		Location string
	}
	sliceB := make([]*IpLocation, 0)
	sliceB = append(sliceB, &IpLocation{Location: "中国台湾", IP: "1.1.1.1"})
	sliceB = append(sliceB, &IpLocation{Location: "中国北京", IP: "2.2.2.2"})
	sliceB = append(sliceB, &IpLocation{Location: "美国硅谷", IP: "3.3.3.3"})

	r := RightJoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err := json.Marshal(r)
	assert.Nil(t, err)
	t.Log(string(marshal))

	flat := FromMapValueFlat(r)
	t.Log(flat)

	// ------------------------------------------------ ---------------------------------------------------------------------

	r = RightJoinByKeyFunc(sliceA, func(indexA int, itemA *User) string {
		return itemA.IP
	}, sliceB, func(indexB int, itemB *IpLocation) string {
		return itemB.IP
	})

	marshal, err = json.Marshal(r)
	assert.Nil(t, err)
	t.Logf(string(marshal))

	flat = FromMapValueFlat(r)
	t.Log(flat)

}
