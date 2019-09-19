package meta

import (
	"testing"
	"fmt"
	"errors"
)

func TestUploadFileMeta(t *testing.T) {
	m := make(map[string]FileMeta)

	fm := FileMeta{FileName: "yss", FileSha1: "aaaaaaaaaa"}

	m[fm.FileSha1] = fm

	meta := m[fm.FileSha1]

	meta.FileName = "YSS"

	t.Log(fm.FileName)
}

func TestChuandi(t *testing.T) {
	a()
}

type C struct {
	p P
}

type P struct {
	v1 int
}

func a() {
	p := P{100}
	c := C{p}
	fmt.Printf("%p %p\n", &c, &c.p.v1) //0xc0000202b8 0xc0000202b8
	b(&c)
	fmt.Println(errors.New("record not found") == errors.New("record not found"))
}

func b(c *C) {
	fmt.Printf("%p %p \n", c, &c.p.v1) //0xc0000202b8 0xc0000202b8
}

func groupAnagrams(strs []string) [][]string {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	m := make(map[int]*[]string)
	temp := 1
	for _, s := range strs {
		byteArr := []byte(s)
		temp = 1
		for _, code := range byteArr {
			temp *= primes[code-97]
		}
		if l, ok := m[temp]; ok {
			*l = append(*l, s)
		} else {
			l := []string{}
			l = append(l, s)
			m[temp] = &l
		}
	}
	rt := make([][]string, 0, len(m))
	for _, v := range m {
		rt = append(rt, *v)
	}
	return rt
}

func BenchmarkAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagrams([]string{"eat", "ate", "tea", "tan", "nat", "bat", "tab"})
	}
	b.StopTimer()
}
