////////////////////////////////////////////////////////////////////////////
// Package: csdnblog.go
// Purpose: Unpack blog.csdn.net json file into each individual files
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type Blog struct {
	Url   string
	Title []string
	Date  []string
	Info  []string
	Blog  []string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "csdnblog"
var buildTime = "2016-10-20"

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {
	// the first command line arguments
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:\n  %s csdn.net-blog.json\n", progname) // os.Args[0])
		os.Exit(0)
	}

	fc, e := ioutil.ReadFile(os.Args[1])
	check(e)

	var blogs []Blog
	json.Unmarshal(fc, &blogs)
	fmt.Printf("Results: %v\n", blogs)
}

////////////////////////////////////////////////////////////////////////////
// check

func check(e error) {
	if e != nil {
		panic(e)
	}
}
