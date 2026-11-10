package utils

import (
	"bytes"
	"os"
	"testing"
)

func TestBundledCSSLoading(t *testing.T) {
	for _, tc := range []struct {
		name    string
		mounted bool
		css     string
	}{
		{name: "no mount"},
		{name: "empty mount", mounted: true},
		{
			name:    "user overrides",
			mounted: true,
			css:     ".clock--time { font-size: 60px; }\n.frame { background: red !important; }",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			if tc.mounted {
				if err := os.WriteFile("custom.css", []byte(tc.css), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			got, err := LoadCustomCSS()
			if err != nil {
				t.Fatal(err)
			}
			want := bundledCSS
			if tc.mounted {
				want = bytes.Join([][]byte{bundledCSS, []byte(tc.css)}, []byte("\n"))
			}
			if !bytes.Equal(got, want) {
				t.Fatal("expected bundled CSS followed by mounted CSS")
			}
		})
	}
}

func TestBundledCSSWeatherBesideClock(t *testing.T) {
	for _, expected := range [][]byte{
		[]byte("#clock-weather-container {"),
		[]byte("left: 0 !important;"),
		[]byte("width: auto !important;"),
		[]byte("justify-content: flex-start;"),
		[]byte("gap: 0.5rem;"),
		[]byte(".weather {"),
		[]byte("#weather-container::before {"),
		[]byte("background: none;"),
		[]byte(".weather--location,\n.weather--temp,\n.weather--stats {\n\tfont-family: -apple-system, BlinkMacSystemFont, sans-serif;\n\ttext-shadow:\n\t\t0 0 2px rgba(0, 0, 0, 0.9),\n\t\t0 0 6px rgba(0, 0, 0, 0.6);\n}"),
		[]byte(".weather--stats {"),
	} {
		if !bytes.Contains(bundledCSS, expected) {
			t.Errorf("bundled CSS is missing weather layout rule %q", expected)
		}
	}
}

func TestBundledCSSOnCustomReadError(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.Mkdir("custom.css", 0o700); err != nil {
		t.Fatal(err)
	}
	got, err := LoadCustomCSS()
	if err == nil {
		t.Fatal("expected an error when custom.css is a directory")
	}
	if !bytes.Equal(got, bundledCSS) {
		t.Fatal("expected bundled CSS to remain available on a custom CSS read error")
	}
}
