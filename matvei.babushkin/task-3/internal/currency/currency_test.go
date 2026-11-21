package currency_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Matvei-Babushkin/task-3/internal/currency"
)

const (
	dirPermissions  os.FileMode = 0o755
	filePermissions os.FileMode = 0o600
)

func TestConversion(t *testing.T) {
	t.Parallel()

	rates, err := currency.ReadXML("../../testdata/currencies.xml")
	if err != nil {
		t.Fatal(err)
	}

	currency.SortByValue(rates)

	want := []currency.Currency{{978, "EUR", 100.5}, {840, "USD", 90.1234}, {36, "AUD", 9.75}}

	if !reflect.DeepEqual(rates, want) {
		t.Fatalf("got %#v, want %#v", rates, want)
	}

	output := filepath.Join(t.TempDir(), "nested", "rates.json")

	if err := currency.WriteJSON(output, rates, dirPermissions, filePermissions); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	var rows []map[string]interface{}
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatal(err)
	}

	if len(rows) != 3 {
		t.Fatalf("got %d rows", len(rows))
	}

	for i, row := range rows {
		if len(row) != 3 ||
			row["num_code"] != float64(want[i].NumCode) ||
			row["char_code"] != want[i].CharCode ||
			row["value"] != want[i].Value {
			t.Fatalf("unexpected JSON row: %#v", row)
		}
	}
}

func TestReadXMLInvalid(t *testing.T) {
	t.Parallel()

	for name, input := range map[string]string{
		"malformed":  "<ValCurs><Valute>",
		"wrong root": "<Other/>",
		"bad number": `<ValCurs>
<Valute>
<NumCode>840</NumCode>
<CharCode>USD</CharCode>
<Value>oops</Value>
</Valute>
</ValCurs>`,
		"missing value": "<ValCurs><Valute><NumCode>840</NumCode><CharCode>USD</CharCode></Valute></ValCurs>",
		"nonfinite": `<ValCurs>
<Valute>
<NumCode>840</NumCode>
<CharCode>USD</CharCode>
<Value>NaN</Value>
</Valute>
</ValCurs>`,
		"leading text":           "garbage<ValCurs/>",
		"empty file":             "",
		"comments only":          "<!-- no root -->",
		"extra root":             "<ValCurs/><Other/>",
		"trailing text":          "<ValCurs/>garbage",
		"trailing malformed XML": "<ValCurs/><",
		"negative value": `<ValCurs>
<Valute>
<NumCode>840</NumCode>
<CharCode>USD</CharCode>
<Value>-1,2</Value>
</Valute>
</ValCurs>`,
		"encoding": `<?xml version="1.0" encoding="unknown"?><ValCurs/>`,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(t.TempDir(), "input.xml")
			if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
				t.Fatal(err)
			}

			if _, err := currency.ReadXML(path); err == nil {
				t.Fatal("expected error")
			}
		})
	}

	if _, err := currency.ReadXML(filepath.Join(t.TempDir(), "missing.xml")); err == nil {
		t.Fatal("expected read error")
	}
}

func TestEmptyXML(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "empty.xml")
	if err := os.WriteFile(path, []byte("<ValCurs/>"), 0o600); err != nil {
		t.Fatal(err)
	}

	rates, err := currency.ReadXML(path)
	if err != nil {
		t.Fatal(err)
	}

	output := filepath.Join(t.TempDir(), "empty.json")

	if err := currency.WriteJSON(output, rates, dirPermissions, filePermissions); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "[]\n" {
		t.Fatalf("got %q", data)
	}
}

func TestWriteJSONErrors(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	blocker := filepath.Join(dir, "file")

	if err := os.WriteFile(blocker, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}

	err := currency.WriteJSON(
		filepath.Join(blocker, "out.json"), []currency.Currency{}, dirPermissions, filePermissions,
	)
	if err == nil {
		t.Fatal("expected directory error")
	}

	if err := currency.WriteJSON(dir, []currency.Currency{}, dirPermissions, filePermissions); err == nil {
		t.Fatal("expected write error")
	}
}

func TestSortUsesValueNotNominal(t *testing.T) {
	t.Parallel()

	// The historical data in the example has different currency nominals.
	input := `<ValCurs>
 <Valute><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Value>43,8254</Value></Valute>
 <Valute><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Value>58,1090</Value></Valute>
 <Valute><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Value>16,0102</Value></Valute>
 </ValCurs><!-- end -->`
	path := filepath.Join(t.TempDir(), "rates.xml")

	if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
		t.Fatal(err)
	}

	rates, err := currency.ReadXML(path)
	if err != nil {
		t.Fatal(err)
	}

	currency.SortByValue(rates)

	want := []currency.Currency{{980, "UAH", 58.109}, {826, "GBP", 43.8254}, {36, "AUD", 16.0102}}

	if !reflect.DeepEqual(rates, want) {
		t.Fatalf("got %#v, want %#v", rates, want)
	}
}

func TestMissingISOCodes(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "historical.xml")
	input := "<ValCurs><Valute><Value>0,6400</Value></Valute></ValCurs>"

	if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
		t.Fatal(err)
	}

	rates, err := currency.ReadXML(path)
	if err != nil {
		t.Fatal(err)
	}

	want := []currency.Currency{{NumCode: 0, CharCode: "", Value: 0.64}}

	if !reflect.DeepEqual(rates, want) {
		t.Fatalf("got %#v, want %#v", rates, want)
	}
}

func TestWriteJSONPermissions(t *testing.T) {
	t.Parallel()

	const (
		privateDirMode   os.FileMode = 0o700
		readOnlyFileMode os.FileMode = 0o400
	)

	output := filepath.Join(t.TempDir(), "private", "rates.json")
	if err := currency.WriteJSON(output, []currency.Currency{}, privateDirMode, readOnlyFileMode); err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Stat(output)
	if err != nil {
		t.Fatal(err)
	}

	if fileInfo.Mode().Perm() != readOnlyFileMode {
		t.Fatalf("file mode: got %o, want %o", fileInfo.Mode().Perm(), readOnlyFileMode)
	}

	dirInfo, err := os.Stat(filepath.Dir(output))
	if err != nil {
		t.Fatal(err)
	}

	if dirInfo.Mode().Perm() != privateDirMode {
		t.Fatalf("directory mode: got %o, want %o", dirInfo.Mode().Perm(), privateDirMode)
	}
}

func TestDecimalValues(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		input    string
		expected float64
	}{
		{"43,8254", 43.8254},
		{"16,0102", 16.0102},
		{" 0,6400 ", 0.64},
		{"10", 10},
		{"0,0001", 0.0001},
		{"001,2500", 1.25},
	} {
		t.Run(testCase.input, func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(t.TempDir(), "input.xml")
			input := "<ValCurs><Valute><Value>" + testCase.input + "</Value></Valute></ValCurs>"

			if err := os.WriteFile(path, []byte(input), filePermissions); err != nil {
				t.Fatal(err)
			}

			rates, err := currency.ReadXML(path)
			if err != nil {
				t.Fatal(err)
			}

			if len(rates) != 1 || rates[0].Value != testCase.expected {
				t.Fatalf("got %#v, want Value %v", rates, testCase.expected)
			}
		})
	}
}

func TestInvalidDecimalValues(t *testing.T) {
	t.Parallel()

	for index, text := range []string{
		"", ",5", "5,", "1,2,3", "1.2", "1,2e3", "NaN", "Inf", "-1,2", "+1,2",
		"1_000,2", "0x1p2", "1 234,56", "1,,2", "1,2.3", strings.Repeat("9", 400),
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(t.TempDir(), "input.xml")
			input := "<ValCurs><Valute><Value>" + text + "</Value></Valute></ValCurs>"

			if err := os.WriteFile(path, []byte(input), filePermissions); err != nil {
				t.Fatal(err)
			}

			if _, err := currency.ReadXML(path); err == nil {
				t.Fatalf("expected error for Value %q", text)
			}
		})
	}
}
