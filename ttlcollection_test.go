package dsutils

import "testing" 
import "log"
import "time"
import "fmt"

func TestTTLCollection(t *testing.T) {
	log.Println("TTLCollection")
	c1 := NewTTLCollection(1.5)

	c1.Add("elem1")
	time.Sleep(500 * time.Millisecond)
	// fmt.Printf("LIST LENGTH=%d\n", len(c1.List()))
	if len(c1.List()) != 1 {
		t.Errorf("List should have len==1")
		panic(1)
	}

	c1.Add("elem2")
	time.Sleep(500 * time.Millisecond)
	if len(c1.List()) != 2 {
		t.Errorf("List should have len==2")
		panic(1)
	}

	c1.Add("elem3")
	time.Sleep(1100 * time.Millisecond)
	fmt.Printf("LIST LENGTH=%d\n", len(c1.List()))
	l := c1.List()
	if len(l) != 1 {
		t.Errorf("List should have len==1")
		panic(1)
	}
	fmt.Printf("%v", l[0])
	if l[0] != "elem3" {
		t.Errorf("elem3 was not found in list")
		panic(1)
	}

	c1.Add("elem4")
	c1.Add("elem5")

	log.Printf("Elements=%v", c1.List())
}
