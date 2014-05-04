package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	apiSearch  = "http://go-search.org/api?action=search&q=%s"
	apiPackage = "http://go-search.org/api?action=package&id=%s"
)

type Search struct {
	Query string
	Hits  []Package
}

func search(query string) error {
	resp, err := http.Get(fmt.Sprintf(apiSearch, query))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var s Search
	if err := json.Unmarshal(body, &s); err != nil {
		return err
	}
	for _, p := range s.Hits {
		fmt.Printf("%s - %s\n", p.Package, p.Synopsis)
	}
	return nil
}

func show(id string) error {
	resp, err := http.Get(fmt.Sprintf(apiPackage, id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var p Package
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Package: %s\n", p.Package)
	fmt.Printf("Project URL: %s\n", p.ProjectURL)
	fmt.Printf("Rank: %d\n", p.StaticRank)
	fmt.Printf("Description:\n%s\n%s\n", strings.Repeat("-", 12), p.Description)
	if p.Imports != nil {
		fmt.Println("\nDepends:")
		for _, v := range p.Imports {
			fmt.Printf("\t%s\n", v)
		}
	}
	return nil
}

func gocmd(verbose bool, param ...string) error {
	cmd := exec.Command("go", param...)
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	if verbose {
		if output, err := cmd.CombinedOutput(); err != nil {
			return err
		} else {
			fmt.Printf("%s\n", output)
		}
	}
	return nil
}

func clean(verbose bool, id, goPath string) error {
	for _, v := range strings.Split(goPath, ":") {
		if err := os.RemoveAll(v + id); err != nil {
			return err
		}
		if verbose {
			fmt.Printf("Removed: %s\n", v+id)
		}
	}
	return nil
}
