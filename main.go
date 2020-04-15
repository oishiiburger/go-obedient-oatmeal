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
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

// Constants
const (
	USAGE string = "Usage: go-obedient-oatmeal filename.txt [number of band names]"
)

// Globals
var lastLen int

func main() {
	switch len(os.Args) {
	case 2:
		n, a := collectWords(os.Args[1])
		fmt.Println("\r" + generateBandName(n, a))
	case 3:
		v, err := strconv.Atoi(os.Args[2])
		if err == nil {
			n, a := collectWords(os.Args[1])
			for i := 0; i < v; i++ {
				fmt.Println("\r" + generateBandName(n, a))
			}
		} else {
			errHandler(err)
		}
	default:
		fmt.Println(USAGE)
	}
}

// Yields the band name string from slices of nouns and verbs.
func generateBandName(nouns, adjs []string) (name string) {
	length := randWithSeed(3) + 2
	var namebuf = new(bytes.Buffer)
	for i := 1; i <= length; i++ {
		if i <= length-1 {
			choice := randWithSeed(2)
			if choice == 0 {
				nidx := randWithSeed(len(nouns))
				namebuf.WriteString(nouns[nidx] + " ")
			} else {
				aidx := randWithSeed(len(adjs))
				namebuf.WriteString(adjs[aidx] + " ")
			}
		} else {
			nidx := randWithSeed(len(nouns))
			namebuf.WriteString(nouns[nidx])
		}
	}
	name = namebuf.String()
	return
}

// Opens the file and parses, tags to yield noun and adj slices.
func collectWords(filename string) ([]string, []string) {
	printWithClear("Reading file " + filename + "...")
	f, ferr := ioutil.ReadFile(filename)
	if ferr != nil {
		errHandler(ferr)
	}
	doc, derr := prose.NewDocument(string(f))
	if derr != nil {
		errHandler(derr)
	}
	tok := doc.Tokens()
	length := len(tok)
	var nouns, adjs []string
	printWithClear("") //fix this
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
	printWithClear("    ") //fix this
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

// Checks to see if a string element is a member of a slice.
// Returns a bool and the index (or -1).
func isMember(val string, slice []string) (bl bool, idx int) {
	for i, v := range slice {
		if val == v {
			return true, i
		}
	}
	return false, -1
}

// Returns a random int in range after a fresh seed from time with delay.
func randWithSeed(max int) (rnd int) {
	time.Sleep(1)
	rand.Seed(time.Now().UnixNano())
	rnd = rand.Intn(max)
	return
}

// Prints with a carriage return and rubbing out the previous line.
func printWithClear(str string) {
	fmt.Print("\r")
	for i := 0; i < lastLen; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r" + str)
	lastLen = len(str)
}

// Does some error logging and quits.
func errHandler(err error) {
	printWithClear("")
	log.Fatal(err)
	os.Exit(1)
}
