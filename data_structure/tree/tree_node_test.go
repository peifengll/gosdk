package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

type NodeData struct {
	Name   string
	Action func() `json:"-"` //这样，json.Marshal 在序列化 NodeData 时会跳过 Action 字段，从而避免报错
	Age    int
}

// 实现这个方法也可以
func (c NodeData) MarshalJSON() ([]byte, error) {
	// 此处只序列化 Name 和 Age 字段
	return json.Marshal(struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}{
		Name: c.Name,
		Age:  c.Age,
	})
}

func (c NodeData) String() string {
	return fmt.Sprintf(`{"Name":%s,"Age":%d}`, c.Name, c.Age)
}

func TestNode(t *testing.T) {
	cc := NodeData{
		Name: "root",
		Action: func() {
			fmt.Println("test")
		},
		Age: 1,
	}
	aa := NodeData{
		Name: "lee",
		Action: func() {
			fmt.Println("test")
		},
		Age: 1,
	}

	node := NewNode(cc)
	node.Children = append(node.Children, NewNode(aa))
	str, err := json.Marshal(node)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(str))
}
