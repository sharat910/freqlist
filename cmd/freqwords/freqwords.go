package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/sharat910/freqlist"
)

type wordFreq struct {
	word string
	freq int
}

func getWords() []string {
	f, err := os.Open("./alice_in_wonderland.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(text))
}

func topExact(words []string, n int) (topWords []wordFreq) {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	//fmt.Printf("exact map size: %d\n\n", len(m))
	var allWords []wordFreq
	for w, c := range m {
		allWords = append(allWords, wordFreq{w, c})
	}
	sort.Slice(allWords, func(i, j int) bool { return allWords[i].freq > allWords[j].freq })
	for i := 0; i < n; i++ {
		topWords = append(topWords, allWords[i])
	}
	return
}

func main() {
	n := 50
	k := 5000
	words := getWords()
	fmt.Printf("Total words: %d\n", len(words))
	exactTop := topExact(words, n)
	freqListTop := topUsingFreqList(words, n, k)
	fmt.Printf("%10s %10s %10s %10s\n", "Exact", "ExactFreq", "Approx", "ApproxFreq")
	for i := 0; i < n; i++ {
		fmt.Printf("%10s %10d %10s %10d\n", exactTop[i].word, exactTop[i].freq, freqListTop[i].word, freqListTop[i].freq)
	}
}

func topUsingFreqList(words []string, n, k int) (topWords []wordFreq) {
	list := freqlist.New(k)
	m := make(map[string]*freqlist.Node, k)
	for _, w := range words {
		if node, ok := m[w]; ok {
			list.AccessNode(node)
		} else {
			node, deletedNodeKey := list.NewNode(w)
			m[w] = node
			if deletedNodeKey != nil {
				delete(m, deletedNodeKey.(string))
			}
		}
	}
	//fmt.Printf("approx map size: %d\n", len(m))
	//list.PrintStats()
	var allWords []wordFreq
	for w, c := range m {
		allWords = append(allWords, wordFreq{w, c.Freq()})
	}
	sort.Slice(allWords, func(i, j int) bool { return allWords[i].freq > allWords[j].freq })
	for i := 0; i < n; i++ {
		topWords = append(topWords, allWords[i])
	}
	return
}
