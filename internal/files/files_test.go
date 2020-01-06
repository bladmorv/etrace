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

package files_test

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/anonymouse64/etrace/internal/files"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type filesTestSuite struct {
}

var _ = check.Suite(&filesTestSuite{})

func (p *filesTestSuite) SetUpTest(c *check.C) {
}

func (p *filesTestSuite) TestEnsureExistAndOpenExists(c *check.C) {

	// case 2: file doesn't exist, don't delete it
	// case 3: file exists, delete it
	// case 4: file doesn't exist, delete it

	tt := []struct {
		fExists            bool
		fIsDir             bool
		fShouldDelete      bool
		expectedDelete     bool
		expectedErrPattern string
	}{
		{
			fExists:        true,
			fShouldDelete:  true,
			expectedDelete: true,
		},
		{
			fExists:        false,
			fShouldDelete:  true,
			expectedDelete: false,
		},
		{
			fExists:        true,
			fShouldDelete:  false,
			expectedDelete: false,
		},
	}
	for _, t := range tt {
		r := files.MockOSStat(func(name string) (os.FileInfo, error) {
			var err error
			fi := &mockedOsFileInfo{
				isDir: t.fIsDir,
			}
			if !t.fExists {
				err = syscall.ENOENT
			}
			return fi, err
		})

		// do the test

		r()
	}
}

type mockedOsFileInfo struct {
	isDir bool
}

func (*mockedOsFileInfo) Name() string {
	return ""
}

func (*mockedOsFileInfo) Size() int64 {
	return 0
}

func (*mockedOsFileInfo) Mode() os.FileMode {
	return 0
}

func (m *mockedOsFileInfo) ModTime() time.Time {
	return time.Now()
}

func (m *mockedOsFileInfo) IsDir() bool {
	return m.isDir
}

func (*mockedOsFileInfo) Sys() interface{} {
	return nil
}
