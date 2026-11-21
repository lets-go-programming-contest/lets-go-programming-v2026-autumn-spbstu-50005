package rates

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Haist-PM/task-3/internal/filedata"
	"golang.org/x/text/encoding/charmap"
)

var (
	errUnsupportedEncoding = errors.New("unsupported XML encoding")
	errNonFiniteValue      = errors.New("currency value must be finite")
	errMissingRoot         = errors.New("expected ValCurs as the XML root")
	errTrailingXML         = errors.New("unexpected data after ValCurs")
)

type xmlDocument struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Entries []xmlEntry `xml:"Valute"`
}

type xmlEntry struct {
	ID       string `xml:"ID,attr"`
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(label, "windows-1251") {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("%w: %q", errUnsupportedEncoding, label)
}

func nextXMLContent(decoder *xml.Decoder) (xml.Token, error) {
	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("read XML token: %w", err)
		}

		switch token := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if len(bytes.TrimSpace(token)) == 0 {
				continue
			}
		}

		return token, nil
	}
}

func decodeXML(data []byte, target any) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	firstToken, err := nextXMLContent(decoder)
	if err != nil {
		return err
	}

	root, ok := firstToken.(xml.StartElement)
	if !ok {
		return errMissingRoot
	}

	if err := decoder.DecodeElement(target, &root); err != nil {
		return fmt.Errorf("parse CBR XML: %w", err)
	}

	if _, err := nextXMLContent(decoder); !errors.Is(err, io.EOF) {
		if err != nil {
			return err
		}

		return errTrailingXML
	}

	return nil
}

func parseValue(text string) (float64, error) {
	numberText := strings.TrimSpace(text)
	whole, fraction, hasComma := strings.Cut(numberText, ",")

	if hasComma {
		numberText = whole + "." + fraction
	}

	const bits = 64

	value, err := strconv.ParseFloat(numberText, bits)
	if err != nil {
		return 0, fmt.Errorf("parse Value %q: %w", text, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("%w: %q", errNonFiniteValue, text)
	}

	return value, nil
}

func Load(path string) ([]Currency, error) {
	document, err := filedata.Read[xmlDocument](path, decodeXML)
	if err != nil {
		return nil, fmt.Errorf("read CBR XML: %w", err)
	}

	currencies := make([]Currency, 0, len(document.Entries))

	for _, entry := range document.Entries {
		value, err := parseValue(entry.Value)
		if err != nil {
			return nil, fmt.Errorf("currency %q: %w", entry.ID, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  entry.NumCode,
			CharCode: entry.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func SortDescending(currencies []Currency) {
	sort.SliceStable(currencies, func(first, second int) bool {
		return currencies[first].Value > currencies[second].Value
	})
}
