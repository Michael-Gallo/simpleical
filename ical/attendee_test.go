package ical

import (
	"net/url"
	"testing"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
)

func BenchmarkParseAttendee(b *testing.B) {
	params := map[string]string{model.ParamCN: "Jane Doe", model.ParamRole: "REQ-PARTICIPANT"}
	value := "mailto:jdoe@example.com"
	for b.Loop() {
		_, _ = parseAttendee(value, params)
	}
}

func TestParseAttendee(t *testing.T) {
	testCases := []struct {
		name             string
		value            string
		params           map[string]string
		expectedAttendee *model.Attendee
		expectedError    error
	}{
		{
			name:   "Valid attendee with common name",
			value:  "mailto:jdoe@example.com",
			params: map[string]string{model.ParamCN: "Jane Doe"},
			expectedAttendee: &model.Attendee{
				CommonName: "Jane Doe",
				CalAddress: &url.URL{Scheme: "mailto", Opaque: "jdoe@example.com"},
			},
		},
		{
			name:  "Valid attendee with scheduling params",
			value: "mailto:jdoe@example.com",
			params: map[string]string{
				model.ParamRole:     "REQ-PARTICIPANT",
				model.ParamPartStat: "ACCEPTED",
				model.ParamRSVP:     model.ParamTrue,
				model.ParamCUType:   "INDIVIDUAL",
				model.ParamCN:       "Jane Doe",
			},
			expectedAttendee: &model.Attendee{
				CommonName: "Jane Doe",
				CalAddress: &url.URL{Scheme: "mailto", Opaque: "jdoe@example.com"},
				Role:       "REQ-PARTICIPANT",
				PartStat:   "ACCEPTED",
				RSVP:       true,
				CUType:     "INDIVIDUAL",
			},
		},
		{
			name:  "Valid attendee with delegation and directory",
			value: "mailto:jdoe@example.com",
			params: map[string]string{
				model.ParamCN:            "Jane Doe",
				model.ParamDir:           "ldap://example.com:6666/o=DC%20Associates,c=US",
				model.ParamSentBy:        "mailto:jan_doe@example.com",
				model.ParamLanguage:      "en-us",
				model.ParamDelegatedFrom: "mailto:bob@example.com",
				model.ParamDelegatedTo:   "mailto:hcabot@example.com,mailto:jqpublic@example.com",
				model.ParamMember:        "mailto:projectA@example.com,mailto:projectB@example.com",
			},
			expectedAttendee: &model.Attendee{
				CommonName: "Jane Doe",
				CalAddress: &url.URL{Scheme: "mailto", Opaque: "jdoe@example.com"},
				Directory:  &url.URL{Scheme: "ldap", Host: "example.com:6666", Path: "/o=DC Associates,c=US"},
				SentBy:     &url.URL{Scheme: "mailto", Opaque: "jan_doe@example.com"},
				Language:   "en-us",
				DelegatedFrom: []*url.URL{
					{Scheme: "mailto", Opaque: "bob@example.com"},
				},
				DelegatedTo: []*url.URL{
					{Scheme: "mailto", Opaque: "hcabot@example.com"},
					{Scheme: "mailto", Opaque: "jqpublic@example.com"},
				},
				Member: []*url.URL{
					{Scheme: "mailto", Opaque: "projectA@example.com"},
					{Scheme: "mailto", Opaque: "projectB@example.com"},
				},
			},
		},
		{
			name:  "Valid attendee with other params",
			value: "mailto:jdoe@example.com",
			params: map[string]string{
				model.ParamCN: "Jane Doe",
				"MISCFIELD":   "TEST",
			},
			expectedAttendee: &model.Attendee{
				CommonName: "Jane Doe",
				CalAddress: &url.URL{Scheme: "mailto", Opaque: "jdoe@example.com"},
				OtherParams: map[string]string{
					"MISCFIELD": "TEST",
				},
			},
		},
		{
			name:          "Invalid RSVP value",
			value:         "mailto:jdoe@example.com",
			params:        map[string]string{model.ParamRSVP: "MAYBE"},
			expectedError: icalerr.ErrInvalidAttendee,
		},
		{
			name:          "Invalid cal-address value",
			value:         "%",
			params:        map[string]string{model.ParamCN: "Jane Doe"},
			expectedError: icalerr.ErrInvalidAttendee,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			attendee, err := parseAttendee(testCase.value, testCase.params)
			if testCase.expectedError != nil {
				assert.ErrorIs(t, err, testCase.expectedError)
				assert.Nil(t, attendee)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedAttendee, attendee)
		})
	}
}
