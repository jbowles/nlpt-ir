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
	"strings"
)

// TYPES //////////////////////////////////////////////////
type DocumentVector struct {
	V []Document
}

type Document struct {
	Token     string
	Input     string
	Slice     []string
	Logarithm string
	//put tokenizer digest here
}

// TF-IDF //////////////////////////////////////////////////

func NewDocumentVector(input []string) (dv *DocumentVector) {
	dv = &DocumentVector{}
	for _, document := range input {
		d := &Document{Slice: []string{document}}
		dv.V = append(dv.V, *d)
	}
	return
}

// TermCount gets the total count of all terms in a document
// param:
// return:
func (d *Document) TermCount() float64 {
	words := strings.Split(d.Input, " ") // TODO: use the tokenizer here
	return float64(len(words))
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

func (f *VecField) Build(dv *DocumentVector) {
	for _, d := range dv.V {
		fmt.Printf("%T,  %v\n\n", d, d)
		//initialize Space map
		f.Space = make(map[string][]Vector)
		for docNum, doc := range d.Slice {
			fmt.Printf("%v\n\n", doc)
			for idx, word := range strings.Fields(doc) {
				v, ok := f.Space[word]
				if !ok {
					v = nil
				}
				d.Token = word
				d.Input = doc
				d.Logarithm = "log"
				//tfidf_product := d.TfIdf()
				f.Space[word] = append(v, Vector{docNum, idx, d.TfIdf()})
			}
		}
		fmt.Printf("END: %v\n\n", d)
	}
}

//d := NewDocument(list []string)
//vf := &VecField{}
//vf.Compose(d)
