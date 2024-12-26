package go_grammar

import "testing"

// "any" ==  "interface{}"

func TestMap(t *testing.T) {
	// key value 形式的数据结构
	m1 := map[int]string{1: "a"}
	m2 := map[string]string{"1": "a"}
	m3 := map[string]any{"1": 1, "2": "string", "3": struct{}{}}
	t.Logf("m1: %v\n", m1)
	t.Logf("m2: %v\n", m2)
	t.Logf("m3: %v\n", m3)
}
