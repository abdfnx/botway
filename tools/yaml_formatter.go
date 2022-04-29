package tools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func YamlFMT(file string) {
	go FormatFile(file, true)
	FormatStream(os.Stdin, os.Stdout)

	return
}

func FormatFile(file string, overwrite bool) {
	r, err := os.Open(file)

	defer r.Close()

	if err != nil {
		log.Fatalf("Error while reading the file %s: '%s'", file, err)
	}

	var out bytes.Buffer

	if err := FormatStream(r, &out); err != nil {
		log.Fatalf("Failed formatting YAML stream: '%s'", err)
	}

	if e := DumpStream(&out, file, overwrite); e != nil {
		log.Fatalf("Cannot overwrite %s: '%s'", file, e)
	}
}

func FormatStream(r io.Reader, out io.Writer) error {
	d := yaml.NewDecoder(r)

	in := yaml.Node{
		Kind:        0,
		Style:       0,
		Tag:         "",
		Value:       "",
		Anchor:      "",
		Alias:       &yaml.Node{},
		Content:     []*yaml.Node{},
		HeadComment: "",
		LineComment: "",
		FootComment: "",
		Line:        0,
		Column:      0,
	}

	err := d.Decode(&in)

	for err == nil {
		e := yaml.NewEncoder(out)
		e.SetIndent(2)

		if err := e.Encode(&in); err != nil {
			log.Fatal(err)
		}

		e.Close()

		if err = d.Decode(&in); err == nil {
			fmt.Fprintln(out, "---")
		}
	}

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func DumpStream(out *bytes.Buffer, f string, overwrite bool) error {
	if overwrite {
		return os.WriteFile(f, out.Bytes(), 0744)
	}

	_, err := io.Copy(os.Stdout, out)

	return err
}
