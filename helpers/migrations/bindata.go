// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/001_initial_schema_down.sql
// migrations/001_initial_schema_up.sql
package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __001_initial_schema_downSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x56\x28\xc9\x4f\xc9\x2f\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\x06\x04\x00\x00\xff\xff\x2d\xcf\xfe\x3f\x29\x00\x00\x00")

func _001_initial_schema_downSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initial_schema_downSql,
		"001_initial_schema_down.sql",
	)
}

func _001_initial_schema_downSql() (*asset, error) {
	bytes, err := _001_initial_schema_downSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_initial_schema_down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initial_schema_upSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\xc1\x8a\x83\x30\x14\x45\xf7\xf9\x8a\xbb\x54\x98\x3f\x98\x55\x74\x9e\x20\x93\xc6\x12\x23\xd4\x95\x04\x92\x96\xd0\x9a\x80\xd1\xff\x2f\xad\x6d\x69\xc1\x5d\xd7\x97\xc3\x3d\xa7\x54\xc4\x35\x41\xf3\x42\x10\x96\xe4\xa6\x84\x8c\x01\x80\xb7\x68\x49\xd5\x5c\x60\xaf\xea\x1d\x57\x3d\xfe\xa9\xc7\x1f\x55\xbc\x13\x1a\x27\x17\x86\xc9\x04\x1b\xc7\x61\x59\xbc\xcd\xf2\x9f\x3b\x74\xf4\x53\x9a\x83\x19\x1d\x34\x1d\x34\x64\xa3\x21\x3b\x21\xd6\xf1\x62\xb6\x36\x96\xff\x32\xf6\x61\x31\x47\x1b\xbf\xb2\xb8\x65\xd4\x2f\xf0\x79\x04\x45\x15\x29\x92\x25\xb5\x6b\x68\xe6\xed\x03\x98\x4d\x3a\x6f\x19\xdb\x18\x1c\x8a\xa6\x11\xc4\xe5\xbb\xf0\x35\x00\x00\xff\xff\xb3\x72\x0e\xb6\x35\x01\x00\x00")

func _001_initial_schema_upSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initial_schema_upSql,
		"001_initial_schema_up.sql",
	)
}

func _001_initial_schema_upSql() (*asset, error) {
	bytes, err := _001_initial_schema_upSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_initial_schema_up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"001_initial_schema_down.sql": _001_initial_schema_downSql,
	"001_initial_schema_up.sql":   _001_initial_schema_upSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"001_initial_schema_down.sql": &bintree{_001_initial_schema_downSql, map[string]*bintree{}},
	"001_initial_schema_up.sql":   &bintree{_001_initial_schema_upSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
