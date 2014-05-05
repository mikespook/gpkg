gpkg
====

[![Build Status](https://travis-ci.org/mikespook/gpkg.png?branch=master)](https://travis-ci.org/mikespook/gpkg)

This command provides a tool for searching and managing as well as querying information about golang packages.

gpkg is using [go-search](http://go-search.org/infoapi)'s API to grab packages information witch created by [David Deng](http://daviddengcn.com/).

Usage
=====

	go command [arguments]

The commands are:

	var			Show all of Go environment variables
	upgrade		Perform an upgrade
	install		Install new packages
	remove		Remove packages
	download	Download the package only
	search		Search the package list for a regex pattern
	show		Show a readable record for the package

Use "gpkg help [command]" for more information about a command.

Authors
=======

 * Xing Xing <mikespook@gmail.com> [Blog](http://mikespook.com) 
[@Twitter](http://twitter.com/mikespook)

Open Source - MIT Software License
==================================

See LICENSE.
