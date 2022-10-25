package slice

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoGroup(t *testing.T) {

}

func TestFoldByKeyFunc(t *testing.T) {

}

func TestFullJoinByKeyFunc(t *testing.T) {

}

func TestJoinByKeyFunc(t *testing.T) {

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
