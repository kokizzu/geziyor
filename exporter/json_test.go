package exporter

import "testing"

func TestJSONExporter_Export(t *testing.T) {
	ch := make(chan interface{})
	defer close(ch)

	exporter := &JSONExporter{
		FileName: "test.json",
		Indent:   " ",
	}
	go exporter.Export(ch)

	ch <- map[string]string{"key": "value"}
}
