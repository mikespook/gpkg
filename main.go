package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Parse()
}

func help(cmd string) {
	switch cmd {
	case "var":
		usage := `usage: gpkg var

Display all Go environment variables.
`
		fmt.Println(usage)
	case "upgrade":
		usage := `usage: gpkg upgrade [package]

Update the named packages and their dependencies.
`
		fmt.Println(usage)
	case "install":
		usage := `usage: gpkg install [package]

Install compiles and installs the packages named by the import paths,
along with their dependencies.
`
		fmt.Println(usage)
	case "remove":
		usage := `usage: gpkg remove [package]

Remove object files and package source directories.
`
		fmt.Println(usage)
	case "download":
		usage := `usage: gpkg download [package]

Download the package only; that is, it instructs get not to install
the packages.
`
		fmt.Println(usage)
	case "search":
		usage := `usage: gpkg search [key words]

Search packages by key words.

If searching for multi key words with spaces, please add a pair of
quoters:

	> go search "rbac web"
`
		fmt.Println(usage)
	case "show":
		usage := `usage: gpkg show [package]

Show a readable record for the package.
`
		fmt.Println(usage)
	default:
		usage()
	}
}

func usage() {
	usage := `gpkg is a tool for searching and managing Go packages.

Usage:

	go command [arguments]

The commands are:

	var			Show all of Go environment variables
	upgrade		Perform an upgrade
	install		Install new packages
	remove		Remove packages
	download	Download the package only
	search		Search the package list for a regex pattern
	show		Show a readable record for the package

Use "gpkg help [command]" for more information about a command.`
	fmt.Println(usage)
}

func main() {
	if os.Getenv("GOPATH") == "" {
		fmt.Println("You must install Go and setup the environment variable $GOPATH first.")
		return
	}
	if flag.NArg() < 1 {
		usage()
		return
	}
	switch flag.Arg(0) {
	case "var":
		variable()
	case "upgrade":
		id := flag.Arg(1)
		if err := gocmd("get", "-u", id); err != nil {
			fmt.Println(err)
		}
	case "install":
		id := flag.Arg(1)
		if err := gocmd("get", id); err != nil {
			fmt.Println(err)
		}
	case "remove":
		id := flag.Arg(1)
		if err := gocmd("clean", id); err != nil {
			fmt.Println(err)
		}
		if err := clean(id); err != nil {
			fmt.Println(err)
		}
	case "download":
		id := flag.Arg(1)
		if err := gocmd("get", "-d", id); err != nil {
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
	case "help":
		help(flag.Arg(1))
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
