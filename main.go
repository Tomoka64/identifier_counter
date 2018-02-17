package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/fatih/color"
)

type File struct {
	File *token.File
}
type Slice struct {
	S string
	N int
}

type Slices []Slice

func (f Slices) Len() int {
	return len(f)
}

func (f Slices) Less(i, j int) bool { //Greater Greaterというの名前のがなかったのでLessの＞の向き変えました
	return f[i].N > f[j].N
}

func (f Slices) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// var num = flag.Int("num", 5, "put a number")

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "<usage> %s <directory>", os.Args[0])
		os.Exit(1)
	}
	f := token.NewFileSet()

	for _, arg := range os.Args[1:] {
		b, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatalln("could not read the file. Error:", err)
		}
		ff := f.AddFile(arg, f.Base(), len(b))
		color.Yellow("Directory: %v", ff.Name())
		color.Red("Size: %d\n", ff.Size())
		color.Blue("Base: %d\n", ff.Base())

		file := File{
			File: ff,
		}

		file.WordCount(b)

	}
}

func (f *File) WordCount(b []byte) {
	counts := make(map[string]int)
	var s scanner.Scanner
	s.Init(f.File, b, nil, scanner.ScanComments)
	for {
		_, t, l := s.Scan()
		if t == token.EOF {
			break
		}
		if t == token.IDENT {
			counts[l]++
		}
	}
	result := make([]Slice, 0, len(counts))
	for s, n := range counts {
		result = append(result, Slice{s, n})
	}
	sort.Sort(Slices(result))
	flag.Parse()
	for i := 0; i < len(result) && i < 5; i++ {
		fmt.Printf("%s\t%d\n", result[i].S, result[i].N)
	}
}
