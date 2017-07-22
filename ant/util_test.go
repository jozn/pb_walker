package ant

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	for i := 1; i < 100000; i++ {
		r := Hash("sdds" + strconv.Itoa(i))
		fmt.Println(r)
		if r < 0 {
			t.Error("r is negative")
		}
	}
}

func TestHashEqual1Miliion(t *testing.T) {
	N := 1000000
	mp := make(map[int]bool, N)
	for i := 1; i < N; i++ {
		r := Hash("sddshkjhkjhkkjkjkhjjh" + strconv.Itoa(i))
		if b := mp[r]; b {
			t.Error("r is already exists: ", r)
		}
		mp[r] = true
	}
}

func TestMyHashHashEqual1Miliion(t *testing.T) {
	N := 100
	mp := make(map[int]bool, N)
	for i := 1; i < N; i++ {
		r := MyHash("sddshkjhkjhkkjkjkhjjh" + strconv.Itoa(i))
		if b := mp[r]; b {
			t.Error("r is already exists: ", r)
		}
		fmt.Println(r)
		mp[r] = true
	}
}

func TestMyHash2HashEqual1Miliion(t *testing.T) {
	N := 100
	mp := make(map[int]bool, N)
	for i := 1; i < N; i++ {
		r := MyHash2("sddshkjhkjhkkjkjkhjjh." + strconv.Itoa(i))
		if b := mp[r]; b {
			t.Error("r is already exists: ", r)
		}
		fmt.Println(r)
		mp[r] = true
	}
}
