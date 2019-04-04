package qr

import (
	"encoding/base64"
	"errors"

	"github.com/boombuler/barcode/utils"
)

func encodeRawBytes(content string, ecl ErrorCorrectionLevel) (*utils.BitList, *versionInfo, error) {
	data, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, nil, err
	}

	vi := findSmallestVersionInfo(ecl, byteMode, len(data)*8)
	if vi == nil {
		return nil, nil, errors.New("To much data to encode")
	}

	// It's not correct to add the unicode bytes to the result directly but most readers can't handle the
	// required ECI header...
	res := new(utils.BitList)
	res.AddBits(int(byteMode), 4)
	res.AddBits(len(content), vi.charCountBits(byteMode))
	for _, b := range data {
		res.AddByte(b)
	}
	addPaddingAndTerminator(res, vi)
	return res, vi, nil
}

func encodeUnicode(content string, ecl ErrorCorrectionLevel) (*utils.BitList, *versionInfo, error) {
	data := []byte(content)

	vi := findSmallestVersionInfo(ecl, byteMode, len(data)*8)
	if vi == nil {
		return nil, nil, errors.New("To much data to encode")
	}

	// It's not correct to add the unicode bytes to the result directly but most readers can't handle the
	// required ECI header...
	res := new(utils.BitList)
	res.AddBits(int(byteMode), 4)
	res.AddBits(len(content), vi.charCountBits(byteMode))
	for _, b := range data {
		res.AddByte(b)
	}
	addPaddingAndTerminator(res, vi)
	return res, vi, nil
}
