/*
 * Copyright (C) 2019 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package files

import "os"

func MockOSStat(mocked func(string) (os.FileInfo, error)) func() {
	old := osStat
	osStat = mocked
	return func() {
		osStat = old
	}
}

func MockOSRemove(mocked func(string) error) func() {
	old := osRemove
	osRemove = mocked
	return func() {
		osRemove = old
	}
}

func MockOSCreate(mocked func(string) (*os.File, error)) func() {
	old := osCreate
	osCreate = mocked
	return func() {
		osCreate = old
	}
}

func MockOSOpenFile(mocked func(string, int, os.FileMode) (*os.File, error)) func() {
	old := osOpenFile
	osOpenFile = mocked
	return func() {
		osOpenFile = old
	}
}
