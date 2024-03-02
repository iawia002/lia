package object

import (
	"bufio"
	"bytes"
	"errors"
	"io"

	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// ParseYAMLDocuments parses YAML documents in a file into separate ones.
// eg:
// input:
//
//	apiVersion: v1
//	kind: Namespace
//	metadata:
//		name: test
//
//	---
//	apiVersion: v1
//	kind: Namespace
//	metadata:
//		name: test2
//
// output:
//
//	apiVersion: v1
//	kind: Namespace
//	metadata:
//		name: test
//
//	apiVersion: v1
//	kind: Namespace
//	metadata:
//		name: test2
func ParseYAMLDocuments(contents []byte) ([][]byte, error) {
	docs := make([][]byte, 0)
	reader := k8syaml.NewYAMLReader(bufio.NewReader(bytes.NewReader(contents)))
	for {
		doc, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}
	return docs, nil
}
