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
	"regexp"
	"sort"
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

// ByDate implements sort.Interface for []Blog based on the Date field.
type ByDate []Blog

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date[0] < a[j].Date[0] }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "csdnblog"
var buildTime = "2016-10-20"

var resultFolder = "/tmp/r/"

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

	os.Mkdir(resultFolder, os.ModePerm)
	fndx, err := os.Create(resultFolder + "00Index.html")
	check(err)
	defer fndx.Close()

	fndx.WriteString("<!DOCTYPE html>\r\n<html>\r\n<head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />\r\n<style>\r\nbody {background-color: LightCyan;}\r\nh1   {color: blue;}\r\n</style>\r\n</head>\r\n<body>\r\n<h1>Index</h1>\r\n")

	var blogs []Blog
	json.Unmarshal(fc, &blogs)
	//fmt.Printf("Results: %v\n", blogs)

	sort.Sort(ByDate(blogs))
	for i, b := range blogs {
		_ = i
		// if i > 5 {
		// 	break
		// }

		fname := b.Title[0] + ".html"
		fname = regexp.MustCompile("/").ReplaceAllString(fname, ",")
		fileo, err := os.Create(resultFolder + fname) // b.Date[0][:10] + "_" +
		check(err)

		fmt.Fprintf(fndx, "\r\n<a href='%s'><h2>%s %s</h2></a>\r\n",
			fname, b.Date[0][:10], b.Title[0])
		fileo.WriteString("<!DOCTYPE html>\r\n<html>\r\n<head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />\r\n<style>\r\nbody {background-color: LightCyan;}\r\nh1   {color: blue;}\r\n</style>\r\n</head>\r\n<body>\r\n")
		fmt.Fprintf(fileo, "<h1>%s</h1><p>\r\nFrom %s<p>\r\n"+
			"<address>&nbsp;&nbsp;%s</address><p>\r\n",
			b.Title[0], b.Url, b.Info[0])
		fileo.WriteString(b.Blog[0])
		fileo.WriteString("\r\n</body>\r\n</html>\r\n")

		fileo.Close()
	}

	fndx.WriteString("\r\n\r\n</body>\r\n</html>\r\n")

}

////////////////////////////////////////////////////////////////////////////
// check

func check(e error) {
	if e != nil {
		panic(e)
	}
}
