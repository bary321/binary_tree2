package binary_tree2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestInsert(t *testing.T) {
	tree := NewBtree(20)
	tmp := rand.Int() % 400000000
	min := tmp
	max := tmp
	for i := 0; i < 100000000; i++ {
		Insert(tmp, tree)
		if min > tmp {
			min = tmp
		}
		if max < tmp {
			max = tmp
		}
		tmp = rand.Int() % 40
	}
	Insert(22, tree)
	fmt.Println(tree.Element)

	b := Find(22, tree)
	if b != nil {
		fmt.Println("b:", b.Element)
	}
	b = FindMin(tree)
	fmt.Println("min: ", b.Element)
	if b.Element != min {
		t.Error("min expected", min, "found", b.Element)
	}
	b = FindMax(tree)
	fmt.Println("max", b.Element)
	if b.Element != max {
		t.Error("max expected", max, "found", b.Element)
	}
	tree, err := Delete(22, tree)
	if err != nil {
		fmt.Println(err)
	}
	b = Find(22, tree)
	if b != nil {
		t.Error("delete error")
	}

}
