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
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

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
				fmt.Println("\r" + strconv.Itoa(i+1) + "\t" + generateBandName(n, a))
			}
		} else {
			printUsage()
		}
	default:
		printUsage()
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
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot access file.")
		printUsage()
	}
	size, err := os.Stat(filename)
	if err != nil {
		errHandler(err)
	}
	fmt.Println(("Using file " + filename + ", " +
		strconv.FormatInt(size.Size()/1000, 10) + "kb"))
	doc, err := prose.NewDocument(string(f))
	if err != nil {
		errHandler(err)
	}
	tok := doc.Tokens()
	var nouns, adjs []string
	repunc, err := regexp.Compile("[.,;!/\\'\"\\p{Pd}]+")
	if err != nil {
		errHandler(err)
	}
	for _, word := range tok {
		tword := strings.Title(strings.ToLower(word.Text))
		tword = repunc.ReplaceAllString(tword, "")
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
	return nouns, adjs
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

// Prints the usage message and exits
func printUsage() {
	fmt.Println("Usage: go-obedient-oatmeal filename.txt [number of band names]")
	os.Exit(1)
}

// Does some general error logging and quits.
func errHandler(err error) {
	log.Fatal(err)
	os.Exit(1)
}
