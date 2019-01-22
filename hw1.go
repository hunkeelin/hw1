package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sync/atomic"
)

var (
	pos   = flag.String("positive", "pos.txt", "the list of positive words")
	neg   = flag.String("negative", "neg.txt", "the list of negative words")
	stop  = flag.String("stop", "stop.txt", "the list of stop words")
	data  = flag.String("dataset", "data.txt", "dataset")
	value int
)

func main() {
	flag.Parse()
	d, err := ioutil.ReadFile(*data)
	if err != nil {
		panic(err)
	}
	s, err := ioutil.ReadFile(*stop)
	if err != nil {
		panic(err)
	}
	p, err := ioutil.ReadFile(*pos)
	if err != nil {
		panic(err)
	}
	n, err := ioutil.ReadFile(*neg)
	if err != nil {
		panic(err)
	}
	lex_pos := splitnappend(p, byte(10))
	lex_neg := splitnappend(n, byte(10))
	lex_stop := splitnappend(s, byte(10))
	a, b, c := parseData(d, lex_pos, lex_neg, lex_stop)
	fmt.Println("Value for positive is", a)
	fmt.Println("Value for negative is", b)
	fmt.Println("Value for stop is", c)
}
func parseData(d []byte, p, n, s [][]byte) (int32, int32, int32) {
	var word []byte
	var pos, neg, stop int32
	for _, l := range d {
		if l != byte(32) { // if not white space
			word = append(word, l)
		} else { // when it's whitespace
			switch {
			case byteinslice(p, word):
				atomic.AddInt32(&pos, 1)
			case byteinslice(n, word):
				atomic.AddInt32(&neg, -1)
			case byteinslice(s, word):
				atomic.AddInt32(&stop, -1)
			default:
				//			fmt.Println(string(word))
			}
			word = nil
		}
	}
	return pos, neg, stop
}
func byteinslice(slice [][]byte, ele []byte) bool {
	for _, e := range slice {
		if string(e) == string(ele) {
			return true
		}
	}
	return false
}
func splitnappend(f []byte, delim byte) [][]byte {
	var toreturn [][]byte
	var tmp []byte
	for _, l := range f {
		if l != delim { // if f != new line if delim = \n
			if l != byte(13) {
				tmp = append(tmp, l)
			}
		} else {
			toreturn = append(toreturn, tmp)
			tmp = nil
		}
	}
	return toreturn
}
