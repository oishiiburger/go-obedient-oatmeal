Go Obedient Oatmeal
===================

This program generates random band names using nouns and adjectives parsed from text files.

![Go Obedient Oatmeal sample](https://raw.githubusercontent.com/oishiiburger/go-obedient-oatmeal/master/img/goo-gif.gif)

Usage
-----

```
go-obedient-oatmeal filename_to_parse [number_of_band_names_to_generate]
```

Overview
--------

The program utilizes the [prose](https://github.com/jdkato/prose) library to parse, tokenize, and tag a plaintext file, yielding a list of included nouns and adjectives. A pseudorandom length and indices are determined. Band names may consist of combinations of nouns and adjectives (the choice between them is also pseudorandom), but the last element in any band name must be a noun to ensure outputs are more-or-less grammatical.

This is the [Go](https://golang.org/) version. Another version using [Python3](https://www.python.org/downloads/) is [Obedient Oatmeal](https://github.com/oishiiburger/obedient-oatmeal).