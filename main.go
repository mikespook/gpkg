package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func init() {
	flag.Parse()
}

func usage() {
	usage := `gpkg is a tool for searching and managing Go packages.

Usage:

	go command [arguments]

The commands are:

	path		Show $GOPATH
	upgrade		Perform an upgrade
	install		Install new packages
	remove		Remove packages
	download	Download the package only
	search		Search the package list for a regex pattern
	show		Show a readable record for the package
	depends		Show dependency information for a package

Use "gpkg help [command]" for more information about a command.`
	fmt.Println(usage)
}

func main() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		fmt.Println("You must install Go and setup the environment variable $GOPATH first.")
		return
	}
	if flag.NArg() < 1 {
		usage()
		return
	}
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Display verbose")
	switch flag.Arg(0) {
	case "path":
		fmt.Println("GOPATH:")
		for i, v := range strings.Split(goPath, ":") {
			if i == 0 {
				fmt.Printf("\t%s (default)\n", v)
			} else {
				fmt.Printf("\t%s\n", v)
			}
		}
	case "upgrade":
		id := flag.Arg(1)
		if err := gocmd(verbose, "get", "-u", id); err != nil {
			fmt.Println(err)
		}
	case "install":
		id := flag.Arg(1)
		if err := gocmd(verbose, "get", id); err != nil {
			fmt.Println(err)
		}
	case "remove":
		id := flag.Arg(1)
		if err := gocmd(verbose, "clean", id); err != nil {
			fmt.Println(err)
		}
		if err := clean(verbose, id, goPath); err != nil {
			fmt.Println(err)
		}
	case "download":
		id := flag.Arg(1)
		if err := gocmd(verbose, "get", "-d", id); err != nil {
			fmt.Println(err)
		}
	case "search":
		query := flag.Arg(1)
		if err := search(query); err != nil {
			fmt.Println(err)
		}
	case "show":
		id := flag.Arg(1)
		if err := show(id); err != nil {
			fmt.Println(err)
		}
	case "depends":
	default:
		usage()
	}
}

type Package struct {
	Package     string
	Name        string
	StarCount   int
	Synopsis    string
	Description string
	Imported    []string
	Imports     []string
	ProjectURL  string
	StaticRank  int
}
