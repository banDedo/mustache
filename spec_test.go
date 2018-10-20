package mustache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

type container struct {
	enabled       bool
	acceptMissing bool
}

var enabledTests = map[string]map[string]container{
	"comments.json": map[string]container{
		"Inline": {
			true,
			false,
		},
		"Multiline": {
			true,
			false,
		},
		"Standalone": {
			true,
			false,
		},
		"Indented Standalone": {
			true,
			false,
		},
		"Standalone Line Endings": {
			true,
			false,
		},
		"Standalone Without Previous Line": {
			true,
			false,
		},
		"Standalone Without Newline": {
			true,
			false,
		},
		"Multiline Standalone": {
			true,
			false,
		},
		"Indented Multiline Standalone": {
			true,
			false,
		},
		"Indented Inline": {
			true,
			false,
		},
		"Surrounding Whitespace": {
			true,
			false,
		},
	},
	"delimiters.json": map[string]container{
		"Pair Behavior": {
			true,
			false,
		},
		"Special Characters": {
			true,
			false,
		},
		"Sections": {
			true,
			false,
		},
		"Inverted Sections": {
			true,
			false,
		},
		"Partial Inheritence": {
			true,
			false,
		},
		"Post-Partial Behavior": {
			true,
			false,
		},
		"Outlying Whitespace (Inline)": {
			true,
			false,
		},
		"Standalone Tag": {
			true,
			false,
		},
		"Indented Standalone Tag": {
			true,
			false,
		},
		"Pair with Padding": {
			true,
			false,
		},
		"Surrounding Whitespace": {
			true,
			false,
		},
		"Standalone Line Endings": {
			true,
			false,
		},
		"Standalone Without Previous Line": {
			true,
			false,
		},
		"Standalone Without Newline": {
			true,
			false,
		},
	},
	"interpolation.json": map[string]container{
		"No Interpolation": {
			true,
			false,
		},
		"Basic Interpolation": {
			true,
			false,
		},
		// disabled b/c Go uses "&#34;" in place of "&quot;"
		// both are valid escapings, and we validate the behavior in mustache_test.go
		"HTML Escaping": {
			false,
			true,
		},
		"Triple Mustache": {
			true,
			false,
		},
		"Ampersand": {
			true,
			false,
		},
		"Basic Integer Interpolation": {
			true,
			false,
		},
		"Triple Mustache Integer Interpolation": {
			true,
			false,
		},
		"Ampersand Integer Interpolation": {
			true,
			false,
		},
		"Basic Decimal Interpolation": {
			true,
			false,
		},
		"Triple Mustache Decimal Interpolation": {
			true,
			false,
		},
		"Ampersand Decimal Interpolation": {
			true,
			false,
		},
		"Basic Context Miss Interpolation": {
			true,
			true,
		},
		"Triple Mustache Context Miss Interpolation": {
			true,
			true,
		},
		"Ampersand Context Miss Interpolation": {
			true,
			true,
		},
		"Dotted Names - Basic Interpolation": {
			true,
			false,
		},
		"Dotted Names - Triple Mustache Interpolation": {
			true,
			false,
		},
		"Dotted Names - Ampersand Interpolation": {
			true,
			false,
		},
		"Dotted Names - Arbitrary Depth": {
			true,
			false,
		},
		"Dotted Names - Broken Chains": {
			true,
			true,
		},
		"Dotted Names - Broken Chain Resolution": {
			true,
			true,
		},
		"Dotted Names - Initial Resolution": {
			true,
			false,
		},
		"Interpolation - Surrounding Whitespace": {
			true,
			false,
		},
		"Triple Mustache - Surrounding Whitespace": {
			true,
			false,
		},
		"Ampersand - Surrounding Whitespace": {
			true,
			false,
		},
		"Interpolation - Standalone": {
			true,
			false,
		},
		"Triple Mustache - Standalone": {
			true,
			false,
		},
		"Ampersand - Standalone": {
			true,
			false,
		},
		"Interpolation With Padding": {
			true,
			false,
		},
		"Triple Mustache With Padding": {
			true,
			false,
		},
		"Ampersand With Padding": {
			true,
			false,
		},
	},
	"inverted.json": map[string]container{
		"Falsey": {
			true,
			false,
		},
		"Truthy": {
			true,
			false,
		},
		"Context": {
			true,
			false,
		},
		"List": {
			true,
			false,
		},
		"Empty List": {
			true,
			false,
		},
		"Doubled": {
			true,
			false,
		},
		"Nested (Falsey)": {
			true,
			false,
		},
		"Nested (Truthy)": {
			true,
			false,
		},
		"Context Misses": {
			true,
			true,
		},
		"Dotted Names - Truthy": {
			true,
			false,
		},
		"Dotted Names - Falsey": {
			true,
			false,
		},
		"Internal Whitespace": {
			true,
			false,
		},
		"Indented Inline Sections": {
			true,
			false,
		},
		"Standalone Lines": {
			true,
			false,
		},
		"Standalone Indented Lines": {
			true,
			false,
		},
		"Padding": {
			true,
			false,
		},
		"Dotted Names - Broken Chains": {
			true,
			true,
		},
		"Surrounding Whitespace": {
			true,
			false,
		},
		"Standalone Line Endings": {
			true,
			false,
		},
		"Standalone Without Previous Line": {
			true,
			false,
		},
		"Standalone Without Newline": {
			true,
			false,
		},
	},
	"partials.json": map[string]container{
		"Basic Behavior": {
			true,
			false,
		},
		"Failed Lookup": {
			true,
			false,
		},
		"Context": {
			true,
			false,
		},
		"Recursion": {
			true,
			false,
		},
		"Surrounding Whitespace": {
			true,
			false,
		},
		"Inline Indentation": {
			true,
			false,
		},
		"Standalone Line Endings": {
			true,
			false,
		},
		"Standalone Without Previous Line": {
			true,
			false,
		},
		"Standalone Without Newline": {
			true,
			false,
		},
		"Standalone Indentation": {
			true,
			false,
		},
		"Padding Whitespace": {
			true,
			false,
		},
	},
	"sections.json": map[string]container{
		"Truthy": {
			true,
			false,
		},
		"Falsey": {
			true,
			false,
		},
		"Context": {
			true,
			false,
		},
		"Deeply Nested Contexts": {
			true,
			false,
		},
		"List": {
			true,
			false,
		},
		"Empty List": {
			true,
			false,
		},
		"Doubled": {
			true,
			false,
		},
		"Nested (Truthy)": {
			true,
			false,
		},
		"Nested (Falsey)": {
			true,
			false,
		},
		"Context Misses": {
			true,
			true,
		},
		"Implicit Iterator - String": {
			true,
			false,
		},
		"Implicit Iterator - Integer": {
			true,
			false,
		},
		"Implicit Iterator - Decimal": {
			true,
			false,
		},
		"Implicit Iterator - Array": {
			true,
			false,
		},
		"Dotted Names - Truthy": {
			true,
			false,
		},
		"Dotted Names - Falsey": {
			true,
			false,
		},
		"Dotted Names - Broken Chains": {
			true,
			true,
		},
		"Surrounding Whitespace": {
			true,
			false,
		},
		"Internal Whitespace": {
			true,
			false,
		},
		"Indented Inline Sections": {
			true,
			false,
		},
		"Standalone Lines": {
			true,
			false,
		},
		"Indented Standalone Lines": {
			true,
			false,
		},
		"Standalone Line Endings": {
			true,
			false,
		},
		"Standalone Without Previous Line": {
			true,
			false,
		},
		"Standalone Without Newline": {
			true,
			false,
		},
		"Padding": {
			true,
			false,
		},
	},
	"~lambdas.json": nil, // not implemented
}

type specTest struct {
	Name        string            `json:"name"`
	Data        interface{}       `json:"data"`
	Expected    string            `json:"expected"`
	Template    string            `json:"template"`
	Description string            `json:"desc"`
	Partials    map[string]string `json:"partials"`
}

type specTestSuite struct {
	Tests []specTest `json:"tests"`
}

func TestSpec(t *testing.T) {
	root := filepath.Join(os.Getenv("PWD"), "spec", "specs")
	if _, err := os.Stat(root); err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("Could not find the specs folder at %s, ensure the submodule exists by running 'git submodule update --init'", root)
		}
		t.Fatal(err)
	}

	paths, err := filepath.Glob(root + "/*.json")
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(paths)

	for _, path := range paths {
		_, file := filepath.Split(path)
		enabled, ok := enabledTests[file]
		if !ok {
			t.Errorf("Unexpected file %s, consider adding to enabledFiles", file)
			continue
		}
		if enabled == nil {
			continue
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var suite specTestSuite
		err = json.Unmarshal(b, &suite)
		if err != nil {
			t.Fatal(err)
		}
		for _, test := range suite.Tests {
			runTest(t, file, &test)
		}
	}
}

func runTest(t *testing.T, file string, test *specTest) {
	c, ok := enabledTests[file][test.Name]
	if !ok {
		t.Errorf("[%s %s]: Unexpected test, add to enabledTests", file, test.Name)
	}
	if !c.enabled {
		t.Logf("[%s %s]: Skipped", file, test.Name)
		return
	}

	var out string
	var err error
	if len(test.Partials) > 0 {
		out, err = RenderPartials(test.Template, &StaticProvider{test.Partials}, test.Data)
	} else {
		out, err = Render(test.Template, test.Data)
	}

	skipMissing := c.acceptMissing && ErrorMatchesType(err, ErrorTypeMissingParams)
	if err != nil && !skipMissing {
		t.Errorf("[%s %s]: %s", file, test.Name, err.Error())
		return
	}
	if out != test.Expected && !skipMissing {
		t.Errorf("[%s %s]: Expected %q, got %q", file, test.Name, test.Expected, out)
		return
	}

	t.Logf("[%s %s]: Passed", file, test.Name)
}
