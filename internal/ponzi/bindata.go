// Code generated by go-bindata.
// sources:
// data/roundedCornerNW.ply
// DO NOT EDIT!

package ponzi

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

var _roundedcornernwPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xff\x6a\x83\x30\x10\xc7\xff\xf7\x29\xee\x05\x2a\xb9\x98\x1f\xe6\x69\x86\xab\xd7\x56\xb0\x51\x32\x3b\xea\x9e\x7e\x30\x52\x97\x8b\x65\x34\x4c\x10\xee\xf3\xd5\xfb\x24\xe6\x9c\xc7\xb5\x3a\x4d\xe1\xda\x2d\xd0\x7d\x1c\x87\x01\xb0\x16\x15\x8d\x74\x25\xbf\xc0\x27\x85\x85\xee\xe0\xaa\x39\x4c\x33\x85\x65\x85\xd3\x38\x75\x0b\xdc\xf3\x60\xcd\x83\xaf\x3c\xf0\xbb\x1e\xbf\x6b\xf2\x49\xd7\xed\x78\xe9\x02\x04\xea\xf3\xe8\x1c\x88\x7c\x1e\xbe\x8f\x37\xda\x76\x4d\xfd\x99\xa0\xfd\x7d\x65\xd8\xbe\x04\x9f\x85\xf2\x7f\x8b\xfa\xfe\xed\x42\x5d\x4f\xa1\x12\xb5\xf8\xb9\x60\x2b\x0e\x72\xab\xf6\x0f\xf1\x51\x48\xad\x1f\x77\x75\x10\x75\xe3\x04\xb6\x98\x48\xb0\x76\x06\xb5\xc5\x12\x89\x35\xba\x31\x96\x49\x5a\x65\xad\x76\xaf\x4b\xb0\x46\x44\x54\x82\x49\x8c\x91\xae\x29\x92\x28\x54\x12\x15\x93\xc4\xa8\x40\x12\x97\x4d\x25\x71\x73\x05\x92\x78\x00\x22\x99\x49\x3c\xa6\x02\x49\x1c\x45\x2a\x89\x03\x7b\x59\x22\xf7\xff\xc9\x93\xea\x6f\x89\x04\x64\xdc\x80\x64\xac\xa0\x61\xac\x41\x31\x36\xa0\x19\x5b\x30\x8c\x5b\xb0\x8c\x1d\xb4\x8c\xbf\x03\x00\x00\xff\xff\x42\xda\xa7\xed\x37\x04\x00\x00")

func roundedcornernwPlyBytes() ([]byte, error) {
	return bindataRead(
		_roundedcornernwPly,
		"roundedCornerNW.ply",
	)
}

func roundedcornernwPly() (*asset, error) {
	bytes, err := roundedcornernwPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "roundedCornerNW.ply", size: 1079, mode: os.FileMode(438), modTime: time.Unix(1504758091, 0)}
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
	"roundedCornerNW.ply": roundedcornernwPly,
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
	"roundedCornerNW.ply": &bintree{roundedcornernwPly, map[string]*bintree{}},
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
