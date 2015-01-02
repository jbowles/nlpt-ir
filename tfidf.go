/*
* Copyright ©2015 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* tfidf == Term Frequency Inverse Document Frequency.
 */

package nlpt_ir

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// TYPES //////////////////////////////////////////////////
type Document struct {
	Token     string
	Input     string
	Slice     []string
	Logarithm string
	//put tokenizer digest here
}

// Vector contains values for tf-idf value, document number, and index location of token/term for quicker lookup
type Vector struct {
	docNum     int
	index      int
	dotProduct float64
}

// Field contains a space of the map of the token/term to its Vectors
type VecField struct {
	Space map[string][]Vector
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value []Vector
}

// A slice of Pairs that implements sort.Interface to sort by Value of Hash Map.
type PairList []Pair

// SORT //////////////////////////////////////////////////

// Create needed Sort methods: Len(), Less(), Swap()
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value[0].dotProduct < p[j].Value[0].dotProduct }

// A function to turn a map into a PairList, then sort and return it.
func (m *VecField) SortByTfIdf() PairList {
	p := make(PairList, len(m.Space))
	i := 0
	for k, v := range m.Space {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// TF-IDF //////////////////////////////////////////////////

// TermCount gets the total count of all terms in a document
// param:
// return:
func (d *Document) TermCount() float64 {
	tokens := strings.Split(d.Input, " ") // TODO: use the tokenizer here
	return float64(len(tokens))
}

// TokenFreq gets the frequency of term in a document
// param:
// return:
func (d *Document) TokenFreq() float64 {
	total := d.TermCount()
	count := 0.0
	//TODO: replace strings.Fields with tokenizer
	for _, w := range strings.Fields(d.Input) {
		switch w {
		case d.Token:
			count += 1.0
		}
	}
	return count / total
}

// NumDocsContain calculates the number of documents that cotain one term
// param:
// return:
func (d *Document) NumDocsContain() (count float64) {
	for range d.Slice {
		if d.TokenFreq() > 0.0 {
			count += 1.0
		}
	}
	return
}

// Tf is the technical term frequency of tf-idf
// param:
// return:
func (d *Document) Tf() float64 {
	return (d.TokenFreq() / d.TermCount())
}

// Idf is the inverse document frequency of tf-idf
// param:
// return:
func (d *Document) Idf() (idf float64) {
	// set val for reuse; +1 so we don't get +Inf values
	val := float64(len(d.Slice)+1) / (d.NumDocsContain() + 1)
	switch d.Logarithm {
	case "log":
		idf = math.Log(val) //Log returns the natural logarithm of x.
	case "log10":
		idf = math.Log10(val) //Log10 returns the decimal logarithm of x.
	case "nolog":
		idf = val //no logarithm
	case "log1p":
		idf = math.Log1p(val) //Log1p natural log of 1 plus its argument x
	case "log2":
		idf = math.Log2(val) //Log2 returns the binary log of x.
	default:
		idf = math.Log(val)
	}
	return
}

// TfIdf returns the Term Frequency-Inverse Document Frequency for a word and all documents
func (d *Document) TfIdf() float64 {
	return (d.Tf() * d.Idf())
}

func (f *VecField) Compose(documents []string) {
	d := &Document{Slice: documents, Logarithm: "log"}
	f.Space = make(map[string][]Vector)
	for docNum, doc := range documents {
		for idx, word := range strings.Fields(doc) {
			v, ok := f.Space[word]
			fmt.Printf("%v %v\n", v, ok)
			if !ok {
				fmt.Println("****** aarrrggg")
				v = nil
			}
			d.Token = word
			d.Input = doc
			//tfidf_product := TfIdf(word, doc, documents, "log")
			tfidf_product := d.TfIdf()
			f.Space[word] = append(v, Vector{docNum, idx, tfidf_product})
		}
	}
}
