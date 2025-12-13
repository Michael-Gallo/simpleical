package parse

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
)

func TestParseAttachment(t *testing.T) {
	testCases := []struct {
		name               string
		value              string
		params             map[string]string
		expectedAttachment *model.Attachment
		expectedError      error
	}{
		{
			name:  "valid attachment, CID URI",
			value: "CID:part1.970409T232345@example.com",
			expectedAttachment: &model.Attachment{
				Value: "URI",
				URI: &url.URL{
					Scheme: "cid",
					Opaque: "part1.970409T232345@example.com",
				},
			},
			expectedError: nil,
		},
		{
			name:  "valid attachment, second ical official example",
			value: "ftp://example.com/pub/reports/r-960812.ps",
			params: map[string]string{
				"FMTTYPE": "application/postscript",
			},
			expectedAttachment: &model.Attachment{
				Value: "URI",
				URI: &url.URL{
					Scheme: "ftp",
					Host:   "example.com",
					Path:   "/pub/reports/r-960812.ps",
				},
				FormatType: "application/postscript",
			},
			expectedError: nil,
		},
		{
			name:  "attachment with base64 binary data",
			value: "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4",
			params: map[string]string{
				"FMTTYPE":  "text/plain",
				"ENCODING": "BASE64",
				"VALUE":    "BINARY",
			},
			expectedAttachment: &model.Attachment{
				Value:      "BINARY",
				Encoding:   "BASE64",
				Binary:     "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4",
				FormatType: "text/plain",
			},
		},
		{
			name:  "invalid base64 encoded data",
			value: "hello",
			params: map[string]string{
				"FMTTYPE":  "text/plain",
				"ENCODING": "BASE64",
				"VALUE":    "BINARY",
			},
			expectedAttachment: nil,
			expectedError:      fmt.Errorf("%w: invalid base64 encoded data", errParseErrorInComponent),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			attachment, err := parseAttachment(testCase.value, testCase.params)
			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedAttachment, attachment)
		})
	}
}
