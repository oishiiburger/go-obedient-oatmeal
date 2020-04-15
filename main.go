// Go Obedient Oatmeal
// Chris Graham, 2020
// https://github.com/oishiiburger/go-obedient-oatmeal
//
// This is the Go version of Obedient Oatmeal, a band name generator which parses
// an input text for nouns and adjectives to use in the naming.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

func main() {
	n, a := collectWords("./.data/moby.txt")
	fmt.Println("\r" + generateBandName(n, a))
}

// Yields the band name string from slices of nouns and verbs
func generateBandName(nouns, adjs []string) (name string) {
	sd := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(sd)
	length := rd.Intn(3) + 2
	var namebuf bytes.Buffer
	for i := 1; i <= length; i++ {
		if i <= length-1 {
			choice := rd.Intn(2)
			if choice == 0 {
				nidx := rd.Intn(len(nouns))
				namebuf.WriteString(nouns[nidx] + " ")
			} else {
				aidx := rd.Intn(len(adjs))
				namebuf.WriteString(adjs[aidx] + " ")
			}
		} else {
			nidx := rd.Intn(len(nouns))
			namebuf.WriteString(nouns[nidx])
		}
	}
	name = namebuf.String()
	return
}

// Opens the file and parses, tags to yield noun and adj slices
func collectWords(filename string) ([]string, []string) {
	fmt.Print("Reading file...")
	f, _ := ioutil.ReadFile(filename)
	doc, _ := prose.NewDocument(string(f))
	tok := doc.Tokens()
	length := len(tok)
	var nouns, adjs []string
	fmt.Print("\r               ")
	for i, word := range tok {
		updateProgress(i, length)
		tword := strings.Title(strings.ToLower(word.Text))
		if len(tword) > 1 {
			if word.Tag == "NN" || word.Tag == "NNS" {
				member, _ := isMember(tword, nouns)
				if !member {
					nouns = append(nouns, tword)
				}
			} else if word.Tag == "JJ" {
				member, _ := isMember(tword, adjs)
				if !member {
					adjs = append(adjs, tword)
				}
			}
		}
	}
	fmt.Print("\r    ")
	return nouns, adjs
}

// Gives user feedback. Used while the parse and slice actions are in process.
func updateProgress(iter, tot int) {
	var percent = float32(iter) / float32(tot) * 100
	if percent > 99 {
		percent = 100
	}
	fmt.Printf("\r%d%%", int(percent))
}

// Checks to see if a string element is a member of a slice
// Returns a bool and the index (or -1)
func isMember(val string, slice []string) (bl bool, idx int) {
	for i, v := range slice {
		if val == v {
			return true, i
		}
	}
	return false, -1
}
