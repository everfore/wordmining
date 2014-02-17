package word

import (
	"./wording"
	//"fmt"
	"sort"
)

type Slice wording.Slice

func (s Slice) Len() int           { return len(s) }
func (s Slice) Less(i, j int) bool { return s[i].Score > s[j].Score }
func (s Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *Slice) Sort() {
	sort.Sort(s)
}

// func main() {
// 	s := Slice{{}, {}, {}}
// 	s[0].Score = 100
// 	s[1].Score = 80
// 	s[2].Score = 120
// 	s.Sort()
// 	fmt.Println(s)
// }
