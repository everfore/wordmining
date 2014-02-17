package wording

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Words struct {
	Word  []string
	ID    string
	Score int32
}

type Slice []Words

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
func ReadJson(infile string) *Words {
	buf, err := ioutil.ReadFile(infile)
	CheckErr(err)
	var words Words
	err = json.Unmarshal(buf, &words)
	CheckErr(err)
	return &words
}

func WriteJson(outfile string, words Words) {
	file, _ := os.Create(outfile)
	defer file.Close()
	buf, err := json.Marshal(words)
	CheckErr(err)
	err = ioutil.WriteFile(outfile, buf, 0644)
	CheckErr(err)
}

func (words *Words) ScoreCmp() int32 {
	var s int32
	// if words == nil {
	// 	return 0
	// }
	for _, i := range words.Word {
		l := len(i)
		s += (int32)(l * 10)
	}
	//words.Score = s
	//fmt.Println(s)
	return s
}
func (w *Words) Init() {
	w.Score = w.ScoreCmp()
}

// func main1() {
// 	words := Words{ID: "shaalx"}
// 	words.Word = make([]string, 10, 100)
// 	words.Word = []string{"cat", "beef"}
// 	filename := "words.json"
// 	WriteJson(filename, words)
// 	wo := ReadJson(filename)
// 	fmt.Println(*wo)
// }
