// Copyright 2013 GDG Korea Golang.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"code.google.com/p/go.tools/present"
)

func init() {
	present.PlayEnabled = true
}

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = renderTour(f, "")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succeed.")
}

// renderTour loads tour.article and the relevant HTML templates from the given
// tour root, and renders the template to the provided writer.
func renderTour(w io.Writer, root string) error {
	// Open and parse source file.
	source := filepath.Join(root, "tour.article")
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	doc, err := present.Parse(f, source, 0)
	if err != nil {
		return err
	}

	// Set up templates.
	action := filepath.Join(root, "template", "action.tmpl")
	tour := filepath.Join(root, "template", "tour.tmpl")
	t := present.Template().Funcs(template.FuncMap{"nocode": nocode})
	_, err = t.ParseFiles(action, tour)
	if err != nil {
		return err
	}

	// Render.
	return doc.Render(w, t)
}

// nocode returns true if the provided Section contains
// no Code elements with Play enabled.
func nocode(s present.Section) bool {
	for _, e := range s.Elem {
		if c, ok := e.(present.Code); ok && c.Play {
			return false
		}
	}
	return true
}
