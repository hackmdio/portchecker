package main

import (
	"testing"
)

func TestParseStringToNetPort(t *testing.T) {
	tables := []struct {
		data   string
		result []string
	}{
		{"tcp://127.0.0.0.1:1234", []string{"tcp://127.0.0.0.1:1234", "tcp", "127.0.0.0.1", "1234"}},
		{"tcp://test.123.asda.asd:1234", []string{"tcp://test.123.asda.asd:1234", "tcp", "test.123.asda.asd", "1234"}},
		{"test.123.asda.asd:1234", []string{"test.123.asda.asd:1234", "tcp", "test.123.asda.asd", "1234"}},
		{"test.123.asda.asd", []string{"test.123.asda.asd", "tcp", "test.123.asda.asd", "80"}},
		{"tcp://127.0.0.0.1", []string{"tcp://127.0.0.0.1", "tcp", "127.0.0.0.1", "80"}},
		{"postgres://asdf:asdf@asdas.asdfasf.asdf.asdf:234", []string{"postgres://asdf:asdf@asdas.asdfasf.asdf.asdf:234", "tcp", "asdas.asdfasf.asdf.asdf", "234"}},
		{"postgres://asdf@asdas.asdfasf.asdf.asdf:234", []string{"postgres://asdf@asdas.asdfasf.asdf.asdf:234", "tcp", "asdas.asdfasf.asdf.asdf", "234"}},
		{"postgres://asdas.asdfasf.asdf.asdf:234", []string{"postgres://asdas.asdfasf.asdf.asdf:234", "tcp", "asdas.asdfasf.asdf.asdf", "234"}},
		{"postgres://asdf:asdf@asdas.asdfasf.asdf.asdf", []string{"postgres://asdf:asdf@asdas.asdfasf.asdf.asdf", "tcp", "asdas.asdfasf.asdf.asdf", "5432"}},
		{"postgres://asdf@asdas.asdfasf.asdf.asdf", []string{"postgres://asdf@asdas.asdfasf.asdf.asdf", "tcp", "asdas.asdfasf.asdf.asdf", "5432"}},
		{"postgres://asdas.asdfasf.asdf.asdf", []string{"postgres://asdas.asdfasf.asdf.asdf", "tcp", "asdas.asdfasf.asdf.asdf", "5432"}},
		{"redis://localhost:1234", []string{"redis://localhost:1234", "tcp", "localhost", "1234"}},
		{"redis://localhost", []string{"redis://localhost", "tcp", "localhost", "6379"}},
		{"redis://test-cache", []string{"redis://test-cache", "tcp", "test-cache", "6379"}},
		{"", []string{"", "tcp", "localhost", "80"}},
	}

	for _, entry := range tables {
		result := ParseNetworkString(entry.data)
		if result[0] != entry.result[0] ||
			result[1] != entry.result[1] ||
			result[2] != entry.result[2] ||
			result[3] != entry.result[3] {
			t.Errorf("fail %s, %s, %s", entry.data, result, entry.result)
		}
	}
}