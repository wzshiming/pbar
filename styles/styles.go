// Code generated by go-bindata.
// sources:
// normal.json
// DO NOT EDIT!

package styles

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _normalJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\x51\x4b\xc3\x30\x14\x85\xdf\xef\xaf\xb8\x9c\xe7\x0c\x87\xa0\x0f\xc5\x09\x6e\xb2\x17\xd9\x1c\x52\xf4\xa1\x2b\x23\xd0\xa8\xc1\xb4\x85\x90\x89\xe0\xf6\xdf\x25\x4d\xc9\x6a\x27\x53\x54\xd8\x5b\x4f\x7a\xcf\xc7\xc9\xb9\xc9\x88\xb1\x01\x89\x77\x62\xc6\xb5\xaa\xea\x12\x09\x7b\xc1\xb8\x97\x06\x09\x33\x0f\x89\x59\x60\x35\x97\xa5\x42\xc2\x48\x6b\x27\x0d\x88\x79\x4b\x2c\x30\xd5\xc6\x28\xeb\xcf\x19\x5e\xcf\xd7\x65\x23\x0f\x33\x26\x6b\x6b\x55\xe5\x22\xe5\xae\x36\x7e\xee\xcc\x7f\xa7\xea\xcd\xed\x65\xc0\xc5\xbc\xe6\x54\x3b\xa3\x2e\xd1\x8b\xe3\x0f\x23\xe8\x41\x17\xee\x19\x09\x9f\x0e\xbd\x5a\xdd\xe8\xaa\x68\x86\x3c\x93\x78\x4b\xe2\x4f\x97\x9d\xe9\x62\x26\xed\xcb\xb4\xb6\xde\x98\x79\x63\xe3\xee\xc1\x7a\xb8\xaf\x81\x0d\x92\x3b\x8d\x7d\x6b\xed\x94\x16\xcd\xb7\xaf\xa1\xfc\x11\x82\x6e\x8b\xcc\x5a\xd6\x00\x02\xcb\x25\x5a\xd2\x06\x02\x27\x41\xe4\x61\x3c\x16\xd4\xf8\xa8\xe5\xe6\xd4\x8f\xf5\xd3\x3d\x2e\x64\x51\xe8\xea\x69\x2c\x6d\xdb\x51\x70\xef\x1e\xc9\x68\x37\xaa\xaa\xc3\xa3\x83\xbd\xa5\x9e\x7f\x5e\xea\x58\xda\x7f\xd8\xe9\xaf\xee\x19\x33\x2c\x94\x9d\x84\x1f\x47\xcf\x71\xf4\x0c\x57\x8f\xae\x9b\x22\xa7\x8f\x00\x00\x00\xff\xff\x3b\x0d\x21\x97\x5c\x04\x00\x00")

func normalJsonBytes() ([]byte, error) {
	return bindataRead(
		_normalJson,
		"normal.json",
	)
}

func normalJson() (*asset, error) {
	bytes, err := normalJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "normal.json", size: 1116, mode: os.FileMode(420), modTime: time.Unix(1542101434, 0)}
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
	"normal.json": normalJson,
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
	"normal.json": &bintree{normalJson, map[string]*bintree{}},
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
