package ical

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseAttendee parses a calendar line starting with ATTENDEE.
func parseAttendee(value string, params map[string]string) (*model.Attendee, error) {
	attendee := &model.Attendee{}
	for propName, propValue := range params {
		switch propName {
		case model.ParamCN:
			attendee.CommonName = propValue
		case model.ParamDir:
			parsedURI, err := parseCalAddress(propValue)
			if err != nil {
				return nil, err
			}
			attendee.Directory = parsedURI
		case model.ParamLanguage:
			attendee.Language = propValue
		case model.ParamSentBy:
			parsedURI, err := parseCalAddress(propValue)
			if err != nil {
				return nil, err
			}
			attendee.SentBy = parsedURI
		case model.ParamCUType:
			attendee.CUType = propValue
		case model.ParamRole:
			attendee.Role = propValue
		case model.ParamPartStat:
			attendee.PartStat = propValue
		case model.ParamRSVP:
			switch strings.ToUpper(propValue) {
			case model.ParamTrue:
				attendee.RSVP = true
			case model.ParamFalse:
				attendee.RSVP = false
			default:
				return nil, fmt.Errorf("%w: invalid RSVP value %q", icalerr.ErrInvalidAttendee, propValue)
			}
		case model.ParamMember:
			addresses, err := parseCalAddressList(propValue)
			if err != nil {
				return nil, err
			}
			attendee.Member = addresses
		case model.ParamDelegatedTo:
			addresses, err := parseCalAddressList(propValue)
			if err != nil {
				return nil, err
			}
			attendee.DelegatedTo = addresses
		case model.ParamDelegatedFrom:
			addresses, err := parseCalAddressList(propValue)
			if err != nil {
				return nil, err
			}
			attendee.DelegatedFrom = addresses
		default:
			if attendee.OtherParams == nil {
				attendee.OtherParams = make(map[string]string)
			}
			attendee.OtherParams[propName] = propValue
		}
	}

	parsedURI, err := parseCalAddress(value)
	if err != nil {
		return nil, err
	}
	attendee.CalAddress = parsedURI

	return attendee, nil
}

// parseCalAddress parses a calendar user address (typically a mailto URI).
func parseCalAddress(value string) (*url.URL, error) {
	parsedURI, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", icalerr.ErrInvalidAttendee, err)
	}
	return parsedURI, nil
}

// parseCalAddressList parses a comma-separated list of calendar user addresses.
func parseCalAddressList(propValue string) ([]*url.URL, error) {
	addresses := make([]*url.URL, 0, strings.Count(propValue, ",")+1)
	for part := range strings.SplitSeq(propValue, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		parsedURI, err := parseCalAddress(part)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, parsedURI)
	}
	return addresses, nil
}
