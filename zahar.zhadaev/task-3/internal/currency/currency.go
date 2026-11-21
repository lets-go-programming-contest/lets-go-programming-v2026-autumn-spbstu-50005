package currency

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	folderMode = 0o755
	fileMode   = 0o644
)

var ErrEncodingNotSupported = errors.New("encoding is not supported")

type RateValue float64

func (rv *RateValue) UnmarshalText(data []byte) error {
	normalized := strings.ReplaceAll(strings.TrimSpace(string(data)), ",", ".")
	if normalized == "" {
		return fmt.Errorf("empty rate value")
	}

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("convert rate %q to float64: %w", string(data), err)
	}

	*rv = RateValue(parsed)
	return nil
}

type Rate struct {
	NumCode  int       `json:"num_code"  xml:"NumCode"`
	CharCode string    `json:"char_code" xml:"CharCode"`
	Value    RateValue `json:"value"     xml:"Value"`
}

type valCursDocument struct {
	XMLName xml.Name `xml:"ValCurs"`
	Items   []Rate   `xml:"Valute"`
}

func LoadFromXML(filePath string) ([]Rate, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open xml %q: %w", filePath, err)
	}
	defer func() {
		_ = f.Close()
	}()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("%w: %q", ErrEncodingNotSupported, charset)
	}

	var doc valCursDocument
	if err := dec.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml %q: %w", filePath, err)
	}

	if len(doc.Items) == 0 {
		return []Rate{}, nil
	}

	return doc.Items, nil
}

func OrderByValueDesc(rates []Rate) {
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}

func ExportJSON(outPath string, rates []Rate) error {
	dir := filepath.Dir(outPath)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, folderMode); err != nil {
			return fmt.Errorf("ensure output dir %q: %w", dir, err)
		}
	}

	payload, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal rates to json: %w", err)
	}

	if err := os.WriteFile(outPath, payload, fileMode); err != nil {
		return fmt.Errorf("write json %q: %w", outPath, err)
	}

	return nil
}
