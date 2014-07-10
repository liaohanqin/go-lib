/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

import (
	"net/url"
	"regexp"
	"strings"
)

const (
	SCHEME_FILE  = "file://"
	SCHEME_FTP   = "ftp://"
	SCHEME_HTTP  = "http://"
	SCHEME_HTTPS = "https://"
	SCHEME_SMB   = "smb://"
)

func EncodeURI(uri, scheme string) string {
	if len(uri) < 1 {
		return ""
	}

	filepath := URIToPath(uri)
	u := url.URL{}
	u.Path = filepath
	return scheme + u.String()
}

func DecodeURI(uri string) string {
	if len(uri) < 1 {
		return ""
	}

	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}

	return u.Scheme + "://" + u.Path
}

func URIToPath(uri string) string {
	filepath := deleteStartSpace(uri)

	if isBeginWithStr(filepath, SCHEME_FILE) {
		return filepath[7:]
	} else if isBeginWithStr(filepath, SCHEME_FTP) {
		return filepath[6:]
	} else if isBeginWithStr(filepath, SCHEME_HTTP) {
		return filepath[7:]
	} else if isBeginWithStr(filepath, SCHEME_HTTPS) {
		return filepath[8:]
	} else if isBeginWithStr(filepath, SCHEME_SMB) {
		return filepath[6:]
	} else if isBeginWithStr(filepath, "/") {
		return filepath
	}

	return ""
}

func PathToURI(filepath, scheme string) string {
	if len(filepath) < 1 || len(scheme) < 1 {
		return ""
	}

	switch scheme {
	case SCHEME_FILE:
		return pathToFileURI(filepath)
	case SCHEME_FTP:
		return pathToFtpURI(filepath)
	case SCHEME_HTTP:
		return pathToHttpURI(filepath)
	case SCHEME_HTTPS:
		return pathToHttpsURI(filepath)
	case SCHEME_SMB:
		return pathToSmbURI(filepath)
	}

	return ""
}

func pathToFileURI(filepath string) string {
	filepath = deleteStartSpace(filepath)

	if isBeginWithStr(filepath, "/") {
		return SCHEME_FILE + filepath
	} else if isBeginWithStr(filepath, SCHEME_FILE) {
		return filepath
	}

	return ""
}

func pathToFtpURI(filepath string) string {
	filepath = deleteStartSpace(filepath)

	if isBeginWithStr(filepath, "/") {
		return SCHEME_FTP + filepath
	} else if isBeginWithStr(filepath, SCHEME_FTP) {
		return filepath
	}

	return ""
}

func pathToHttpURI(filepath string) string {
	filepath = deleteStartSpace(filepath)

	if isBeginWithStr(filepath, "/") {
		return SCHEME_HTTP + filepath
	} else if isBeginWithStr(filepath, SCHEME_HTTP) {
		return filepath
	}

	return ""
}

func pathToHttpsURI(filepath string) string {
	filepath = deleteStartSpace(filepath)

	if isBeginWithStr(filepath, "/") {
		return SCHEME_HTTPS + filepath
	} else if isBeginWithStr(filepath, SCHEME_HTTPS) {
		return filepath
	}

	return ""
}

func pathToSmbURI(filepath string) string {
	filepath = deleteStartSpace(filepath)

	if isBeginWithStr(filepath, "/") {
		return SCHEME_SMB + filepath
	} else if isBeginWithStr(filepath, SCHEME_SMB) {
		return filepath
	}

	return ""
}

func deleteStartSpace(str string) string {
	if len(str) <= 0 {
		return ""
	}

	tmp := strings.TrimLeft(str, " ")

	return tmp
}

func isBeginWithStr(str, substr string) bool {
	// TODO could use strings.HasPrefix(), :-)
	ok, _ := regexp.MatchString("^"+substr, str)

	return ok
}
