// Code generated by go-bindata.
// sources:
// ../schemas/jsonschema-draft-04.json
// ../schemas/v2/schema.json
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _jsonschemaDraft04JSON = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x3b\x6f\xdb\x30\x10\xde\xfd\x2b\x04\xa5\x63\x52\xb9\x40\xa7\x6c\x45\xbb\x18\x68\xd1\x0c\xdd\x0c\x0f\xb4\x75\xb2\x19\x50\xa4\x42\x51\x85\x0d\x43\xff\xbd\xa4\xa8\x07\x29\x91\x92\x2d\xbb\x48\xb4\xc4\xe1\xbd\xbe\x3b\xde\x8b\xe7\x45\x20\xbf\x10\xc7\xe1\x73\x10\x1e\x84\xc8\x9e\xa3\xe8\x35\x67\xf4\x29\xdf\x1d\x20\x45\x9f\x19\xdf\x47\x31\x47\x89\x78\x5a\x7e\x8d\xf4\xd9\x43\xf8\xa8\x85\x3e\xe9\xff\x67\x48\xc6\x90\xef\x38\xce\x04\x66\x54\x49\x7f\x67\x1c\x02\xcd\x12\xa4\x20\x50\xad\xa2\xe3\x4e\x30\xc5\x8a\x39\x97\xdc\x1a\x71\x45\xd0\x6c\xdf\x38\x47\x27\x8b\x50\x11\xc5\x29\x03\xa5\x1c\x55\xe4\x47\x9b\x98\x62\xba\x12\x90\x2a\x7d\x5f\x7a\x24\x5c\x9f\x9f\xa5\x83\x1c\x12\xa5\xe2\x21\x0c\xca\x96\xa9\xec\xf8\xc3\x8c\xe5\x12\xd7\x5f\x58\x51\x01\x7b\xe0\x7e\x10\xb8\x66\x18\xc2\xc0\x69\x91\x4a\x8e\xe5\x25\xfa\x7f\x40\x82\x0a\x22\x96\x43\x3b\x88\x90\xdf\x0a\xea\xda\x82\x1d\x19\x91\x8b\xfa\x58\xa5\x21\xc5\x1c\x6b\x9d\x0a\x42\x50\x06\x1b\x27\x8c\x1c\xa7\x19\x81\x3f\xd2\x97\x7c\x68\x1a\x68\xe5\xc0\xba\x8d\x74\x10\x6e\x19\x23\x80\xa8\xfa\xd9\x3a\x1e\x84\xb4\x20\x44\xff\x4d\xb7\xfa\x84\x6d\x5f\x61\x27\xd4\xaf\x5c\x70\x4c\xf7\xa1\xcf\x7e\x45\x9d\x73\xcf\xc6\x65\x36\x7c\x8d\xa9\xf2\xf2\x94\x28\x28\x7e\x2b\xa0\xa1\x0a\x5e\x40\x07\x73\x61\x80\x6d\x6d\x34\x8e\xe9\xd3\x8c\xb3\x0c\xb8\xc0\xbd\xe8\xe9\xa2\xf3\x78\x53\xa3\xec\x01\x49\x18\x4f\x91\xba\xab\xb0\xe0\x38\x74\xc6\xaa\x2b\xca\x7b\x6b\x16\x58\x10\x98\xd4\xeb\x14\xb5\xeb\x7d\x96\x82\x26\x4b\xcf\xe6\x71\x2a\xcf\xb0\x4c\xcd\x2a\xf7\x3d\x6a\x9b\x74\xf3\x56\x5e\x8f\x02\xc7\x1d\x29\x72\x59\x28\xbf\x5a\x16\xfb\xc6\x4d\xfb\xe8\x58\xb3\x8c\x1b\x77\x0a\x77\x86\xa6\xb4\xb4\xf5\x64\x93\xbb\xa0\x24\x88\xe4\x1e\x84\xad\x13\x37\x21\x9c\xd2\x72\x0b\x42\x74\xfc\x09\x74\x2f\x0e\xbd\x9e\x3b\xd5\xbc\x2c\x1f\xaf\xd6\xd0\xb6\x52\xbb\xdf\x22\x21\x80\x4f\xe7\xa8\xb7\x78\xb8\xd4\x7d\x74\x07\x13\xc5\x71\x05\x05\x91\xa6\x91\xf4\x7b\x38\x3d\xe9\x1e\x6e\x1d\xab\xef\x3c\x0c\x74\xbf\x7d\xd5\x6c\xce\x89\xa5\xbe\x8d\xf7\x66\xce\xee\xd1\x86\x67\x80\x34\xad\x8f\xc3\xb3\xae\xc6\x1c\xe3\xb7\xc2\x96\xd9\xb4\x72\x0c\xf0\xab\x92\xe9\x5a\x05\xee\x5c\xb2\x87\xc6\x7f\xa9\x9b\x17\x6b\xb0\xcc\x75\x77\x96\x16\xb7\xcf\x1c\xde\x0a\xcc\x21\x1e\x53\x64\x0e\x73\x4f\x81\xbc\xb8\x07\xa6\xe6\xfa\x50\x55\xe2\x5b\x4d\xad\x4b\xb6\xb6\x81\x49\x77\xc7\xca\x68\x1a\x90\x67\xd7\x78\x3f\x3c\xba\xa3\x8e\xdd\xe8\x7b\xc0\x8a\x21\x03\x1a\x03\xdd\xdd\x11\xd1\x20\xd3\x46\x72\x55\x7d\x93\x0d\xb3\xcf\x34\x52\x46\x03\xd9\x8d\x75\xe2\x0e\x42\xbd\xb9\xdf\xe9\xdd\x34\xb6\x24\x9b\x5b\xa4\x56\x3f\x6b\xac\xd8\x01\x30\x1e\x25\xce\x3a\x77\xc6\x73\xd4\xbd\x96\xc9\xf5\x06\xbc\xca\xf8\x44\xb0\x2e\x09\x5a\xf3\xf5\x3a\x94\x7b\xb7\xa8\x9f\x7f\x17\x8e\x58\x53\xb2\x0e\xfc\xf5\x92\x8c\xc2\x4c\x49\xca\x84\xe7\x7d\x5d\xb6\x2f\x7e\x4f\x79\xba\x96\xe6\x75\xb7\x87\x9b\x0d\xdc\xb5\xbd\xae\xbb\x85\xb8\x8e\x64\x67\xd1\xe8\x18\xe5\xe2\x5f\x00\x00\x00\xff\xff\x4e\x9b\x8d\xdf\x17\x11\x00\x00")

func jsonschemaDraft04JSONBytes() ([]byte, error) {
	return bindataRead(
		_jsonschemaDraft04JSON,
		"jsonschema-draft-04.json",
	)
}

func jsonschemaDraft04JSON() (*asset, error) {
	bytes, err := jsonschemaDraft04JSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonschema-draft-04.json", size: 4375, mode: os.FileMode(436), modTime: time.Unix(1427662469, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _v2SchemaJSON = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x6d\x73\xdb\xb8\xf1\x7f\x9f\x4f\x81\xd1\x65\xc6\xc9\xc4\x96\x72\xf9\xff\xdf\x34\x9d\xce\x8d\x7b\x4e\xaf\x6e\x93\xda\x13\x27\xed\x8b\x58\x99\x81\x48\xc8\xc2\x1d\x45\x2a\x04\x69\x5b\x97\xfa\xbb\x77\xf1\x40\x8a\x20\x01\x12\x7c\x90\x63\x5f\xd8\x99\x26\x39\x11\x58\xec\x2e\x16\xbb\xbf\x5d\x80\xe0\xd7\x27\x08\x4d\x12\x9a\x04\x64\xf2\x1a\x4d\x8e\xd1\x3f\x2e\xce\xfe\x85\x2e\xbc\x15\x59\x63\xb4\x8c\x62\x74\x71\x83\xaf\xae\x48\x8c\x5e\x4d\x5f\xa2\xe3\xf3\xd3\xe9\xe4\x90\x77\xa0\x3e\x6f\xbd\x4a\x92\xcd\xeb\xd9\x8c\xc9\x26\x53\x1a\xcd\xae\x5f\xcd\x98\xe8\x3b\xfd\x95\x45\xe1\x0f\xb2\xf1\x53\xf9\x53\xa1\x07\x7f\x78\xa4\x1a\x46\xf1\xd5\xcc\x8f\xf1\x32\x39\x7a\xf9\xff\xaa\xb3\xea\x97\x6c\x37\x82\xa9\x68\xf1\x2b\xf1\x12\xf9\x5b\x4c\xbe\xa4\x34\x26\x7c\xf8\x4f\x13\x35\xf0\xe4\x10\x18\x0a\x97\x11\xff\x7b\x83\x93\x15\x9b\xcc\x45\x5b\xec\xfb\x34\xa1\x51\x88\x83\xf3\x38\xda\x90\x38\xa1\x84\x41\xbf\x25\x0e\x18\x11\x0d\xa0\x71\x42\xe2\x50\x7b\xca\xf5\x01\x8f\x3e\xdf\x1e\xe5\xff\xc1\x45\x88\xc9\x92\xb3\xf2\xc3\xcc\x27\x4b\x1a\x0a\xb2\x6c\x76\x4d\x42\x3f\x8a\xdf\xdc\x26\x24\x64\xf0\xc3\x44\xb4\xbe\x83\x3f\xef\x24\x79\x03\xdd\x8c\xe5\x02\xed\x4c\x4c\x96\xc4\x34\xbc\x12\x62\x8a\xdf\x49\x98\xae\x85\x98\xa0\x7a\x29\x91\xf8\xd9\x27\xcc\x8b\xe9\x86\x73\xc0\x7b\x7d\x58\x91\x7c\x8e\xae\x49\xcc\xf9\x40\xd1\x12\x25\x2b\xca\x90\x1f\x79\xe9\x9a\x84\xc9\x54\x71\x26\x69\x48\x5d\x35\x0a\x27\x5a\x69\xfd\x56\x11\x4b\x5c\x18\x57\x6a\xe5\x8f\x3e\x7f\xfa\xfc\xf5\x6e\x86\x5e\x5f\xc2\xff\xe6\x2f\x9e\xfd\xf4\x1a\xfe\xe5\xbf\x78\xfe\xd3\xd3\x49\x9d\x3c\xcb\x34\x08\xb6\xe8\x4b\x8a\x03\xba\xa4\xc4\x47\x1f\xdf\x9f\xa2\x24\x02\x99\x08\xe2\x3c\x48\xf9\x88\x34\x47\x8d\xc3\x05\x66\xe4\x1c\x2c\xa0\x2d\x97\xb3\x5a\x76\x38\x55\xc4\x0d\x2b\x63\x82\x0f\x8c\xde\xdc\xe2\xf5\x26\x20\xaf\xd1\xc1\x0c\x6f\xe8\x41\x89\x13\x61\xc8\x85\x89\xb7\xaa\x59\x35\x7c\x4b\x41\xb7\x1a\x05\x0f\x9e\xa6\x25\x12\x25\xe6\x8e\x51\x40\xa5\x3a\xde\x9d\xbe\x7b\x83\xb8\xa4\x0c\x61\xcf\x23\x9b\x04\xb4\xb6\xd8\xee\xb4\x74\x58\xcf\xc4\x9a\xf8\x14\x7f\x80\xee\x55\x36\xc0\x88\xfd\xd4\x6b\xcf\x86\x1a\x1a\x79\x38\x44\x8a\x46\x2f\x36\xc4\xba\x6e\xd4\xa6\x6c\xa6\xf5\x2c\x3c\x6e\xee\x5f\x6c\x5c\x1a\x3f\xc6\x6b\x02\x06\xe3\xc4\x84\x6a\x7b\x62\xa3\x16\x13\xb6\x81\x1f\x5d\xec\x23\x6b\x6a\xa5\xc5\x88\x97\xc6\x34\xd9\x3a\x98\x5a\xd6\xd2\xd8\xff\xa4\x8d\x9e\x4c\x9d\x34\xaa\x09\xbe\x62\xa6\x55\x88\xe3\x18\x6f\x77\x76\x40\x13\xb2\x2e\xb6\xb3\x0e\x08\xf4\x26\xaa\xcd\x5d\xde\x3b\x0d\xe9\x97\x94\x9c\x2a\x1a\x49\x9c\x12\x8d\x07\x72\xcb\x17\x38\x0e\x4e\x22\xcf\x41\x24\xad\x75\xc9\x93\x9b\x6c\xa8\xe2\x46\x0d\xe1\xca\xb4\x5a\x7e\x21\x21\x89\x71\x80\x78\xf7\x78\x8d\xf9\xcf\x08\x2f\xa2\x34\x31\xac\x56\x2d\xda\x29\xef\xce\xa3\x9c\x8c\xd7\xbb\x98\xd0\x14\xe9\xb2\x25\x64\x89\x76\xe2\xb1\x1e\xf1\x6a\x14\x65\x8c\x7a\x99\xbe\xf4\x09\x32\x44\x40\xa9\x2a\x85\x37\xb4\xd1\x2c\x8e\xda\xa4\xc3\x63\x24\xa7\x1e\xe1\xd0\x07\xef\x42\x3c\x0a\x1e\x5a\x10\xad\xc6\x06\x8d\x23\x20\x95\xa9\xb1\xcf\xe8\x0c\x40\x4a\x98\x50\x2f\x8f\xb8\x10\xaa\x17\x10\x80\x1b\x07\xd7\x29\x75\x67\x20\x88\x42\x1e\xf0\x0b\xbf\x17\xc7\x46\x17\xab\x28\x0d\x20\x02\x10\xe4\xd3\xe5\x92\xc4\x80\x01\xd0\x32\x8e\xd6\xa2\x85\xd0\xd3\x14\xa1\x5f\x68\xb2\x4a\x17\x47\xcb\x00\x5f\x47\x60\x63\x68\x8d\xe3\xdf\xfc\xe8\x26\x44\x80\x1c\x70\x10\x44\x37\xc4\xb7\x48\x01\x66\xb4\x66\x67\xcb\x0b\x12\x5f\x53\xaf\xcf\x3c\xf2\xe8\x2a\x88\x71\xee\x99\x24\x27\x50\x67\xbd\x16\x21\x34\x26\xd8\x4b\xdc\xcc\x35\x6b\x6c\xa4\x14\xc0\x80\xe0\x5c\xdd\x28\x65\x8d\xab\x06\x5f\x0e\xdc\x25\xee\x5c\x5d\xc3\xcf\xb2\xa7\xe6\x1a\x32\x6d\xc0\xc4\x80\xad\x69\x16\xd6\x72\xf9\x5b\xd6\x62\x08\xb1\xaa\xe7\x14\x52\x1f\x0c\x8c\x2e\xb7\xd0\x16\x71\x72\x19\x97\x4a\x13\x08\xc6\x05\xc0\x3f\x03\xa4\x8f\x43\xfa\xbb\x90\xcb\x32\xb3\x69\x1c\xf4\xe4\xe5\xe3\xfb\xb7\x68\x13\x51\xe0\x07\x98\x51\x78\xcd\xab\xea\x75\xaa\x13\x92\xbf\x73\x1a\x10\xd6\xcc\xac\xc1\x92\xa7\x7d\x99\x13\x34\x10\x4c\x17\x44\x75\xe6\xa4\x25\x0b\x97\x92\x99\x06\x43\xac\x1a\xb7\xd5\x10\xb5\x40\x23\x2c\xa2\x6d\x70\xd9\x9f\x75\x15\x2d\x4a\x89\x24\x30\xe6\x14\x9d\x26\x07\x0c\x91\xd0\x8b\xd2\x18\x5f\x81\x07\x83\xe9\x4e\x19\x0f\x0a\xe8\xec\x02\x90\x67\xb4\x86\x90\x47\x17\x41\xde\xed\x5e\x8d\x2e\x1f\xd3\xc9\xd0\x8c\x13\x58\x01\xbd\x8e\x7e\xe4\x3d\x09\x40\xf0\x6b\x99\xb4\xb0\x8c\x21\x1a\xfa\xf4\x9a\xfa\x90\x55\x81\xce\x7c\xc1\x2e\x9b\x22\x60\x7f\x8b\xd6\x29\xe0\x77\x88\x16\x71\xd6\x51\x75\x39\xc8\x12\xaa\x83\x69\x25\x71\xda\x23\x8e\x28\xcc\x0d\xa4\x66\x4e\xc4\xb8\xa4\x1c\x08\xd6\xa1\x91\x3a\x43\x76\x49\x19\x6c\xda\xb7\xd0\x6d\xc4\xb4\xaa\x2c\x52\xe1\xb3\x34\x9b\x67\x21\x98\x7f\x8c\xd6\x10\xa4\x65\x75\x46\x8e\xcf\x54\xf4\x5f\x08\x9b\x83\xc9\x92\xe4\x18\xcc\x23\xff\x45\xa5\x8f\xbe\x82\x48\x22\x01\xd3\x73\x42\x73\x7e\x63\x49\x03\x86\x96\x3d\x1f\xaf\xbd\xf8\x31\x01\xc4\xc7\x20\xe2\x08\x17\xc9\x44\x84\x2c\xa4\x67\xc6\x44\xeb\x9e\xa4\xca\x86\xdb\xaf\x50\xb6\xbc\xa6\xa5\x34\xba\x03\x2f\x31\x58\xcd\x4e\xb2\x51\xf3\xaa\x92\x78\x68\x09\x21\xdc\xa5\xce\x1b\x83\x83\x3b\x1a\x1e\xc2\x6f\xb7\x72\xbd\x44\x56\x77\x9c\xf4\xdb\xe4\x10\x3f\xe1\xa3\xdf\x5f\x1e\xfd\xe9\x68\xfe\x62\xa6\xfe\x79\x79\x79\xf4\x62\xfe\xe2\x29\x6f\xd7\xc3\x4d\xad\xe9\x9a\x7c\x90\x3c\xb5\xab\xc9\x5d\x5e\xb2\x1d\x1f\x7f\xbe\xbc\x9c\xfe\xe5\xf2\x72\xc6\xf9\xa9\xab\x82\xe5\xb5\x9d\x2c\x16\xff\xfd\xc3\x87\x73\xb4\x06\x20\x03\xb1\xb7\xe4\x4d\x38\xdb\xb8\x34\xad\x6e\xd8\x63\x57\x19\x79\xc4\xd9\xad\x5e\xfb\xd0\x84\xd7\xeb\x1f\xe2\x51\xb5\x06\x62\x9a\xc9\xc2\xc3\x3b\xad\xbb\xa5\x08\x52\x6a\x38\x81\x60\x00\x09\xde\xb6\x57\xd6\xb9\x88\x29\x81\x3c\x4d\x52\xca\xcc\x20\x9f\xeb\x6f\x96\xf2\xe6\x1c\x1c\xa2\xab\xce\x39\xad\xc5\xa9\x8a\x67\xae\x05\xa3\x32\xcd\x9c\xb1\x53\xbf\x97\xe8\x4b\xd0\x7b\xe8\x07\x5b\x0d\x0a\xef\xd6\x98\x71\x6c\x43\xdd\xd6\xa2\xd7\x96\xb5\xdb\x1a\x8d\x18\xea\xb7\x65\xb6\x0c\x55\xed\x4e\x6c\x29\x3a\x43\xb1\x65\x2c\xef\xd6\x90\xdc\xb5\xb7\xd3\x34\x15\x79\x6b\x48\xee\x9a\x9b\xd7\x6f\x65\x43\xa1\x86\x56\x65\x53\xa1\x4c\xcd\xe7\x70\xc3\xc3\x09\xb1\x1a\xe6\x22\x8a\x02\x82\xc3\xb2\x65\x2e\x71\x1a\x24\x5a\x34\xaa\x30\x5a\xad\x47\xd7\x71\xaa\xd5\xa4\x05\x2d\x6b\x2a\x24\xf0\xfd\x50\x78\xe7\x01\xc5\x0b\x45\xb8\x35\xfe\xb9\x22\x8e\x25\xb0\x26\x57\x91\x0e\x44\x47\xdf\x1f\xec\x4e\xc8\x27\x01\xac\xad\x41\x48\x45\x9b\x32\xe8\xef\x4e\x6b\x45\x70\x65\xb9\x74\x53\x14\x4e\xbc\xd5\x40\x94\x06\xf2\x5b\xc6\x45\x67\xdc\xa6\x72\xae\x41\xc8\xbe\x79\xb6\xca\x63\x17\x13\xbe\x9b\x40\x90\x86\x48\xbe\xe0\x95\x9a\x2d\xba\xc6\x01\xf5\x25\x98\x64\x90\x53\xa4\xd0\x26\xf2\x45\x76\x74\xa0\xdc\x4d\xb1\xf8\xb0\xa6\xfa\x92\xfd\x71\xd8\x55\xff\xec\x13\xa0\xe2\xf9\xd7\xff\xbb\x7b\xfe\xf4\xbf\x9f\x9f\xa9\xf1\x9f\x3f\x6d\xe7\xc1\xff\x8d\x83\x94\x58\xca\x19\x7b\x70\x2b\x61\x94\x94\xf0\xa7\x79\x86\x1c\x75\xd4\xa8\x25\xa3\x18\xed\x05\xd9\x89\xd2\x64\x7e\x52\x9f\x05\x13\x8c\x42\x72\xc6\x87\xfa\xd4\x21\x21\x6f\x4e\xe2\xf9\xe1\x94\xf7\x44\xec\x9a\x78\xbb\x8e\x73\x23\x6b\xad\xf3\x9b\xe2\x2a\xd9\x7b\x82\x9c\x9f\xbb\x71\x85\x0c\xd8\xee\xf8\x9c\x5d\x4c\xd6\xd8\x82\xb2\x2b\x69\x75\x0d\xa9\xbc\x75\x8d\xf5\x3b\x67\xcb\x55\x21\x06\x2f\xff\xc8\x21\x26\x46\x63\x56\xcf\x86\x82\x2f\x9a\x51\x09\x4a\xcd\xd6\x94\xe8\xb5\x02\x8d\x09\x53\x1e\x92\x1f\x40\xca\x1e\x82\xaf\x11\xfb\x9c\xf2\xc0\x55\x42\xd4\xd9\xab\x1c\x2a\x66\xe9\xed\xdc\x38\xf7\x79\xfd\xa5\xb5\x1d\x9b\x32\x64\x5b\x80\x8b\xe9\x9a\xf2\x12\x36\x93\x19\xb1\x91\x9e\x17\x05\x01\xa8\x1c\x3a\xfc\xcd\xc8\x93\x6d\x23\xb1\xd4\xcb\x82\x5d\x32\x88\xec\x40\x32\x6b\x6c\xa4\xb4\xc6\xb7\x74\x9d\xae\xdd\x28\x65\x8d\x2d\xab\xce\x0b\x52\x06\x4a\x79\xd7\x86\x64\xa5\x97\x99\x4b\x68\xef\xce\xa5\x6a\xdc\xc0\x65\x1b\x92\x95\x5e\x36\x5d\xbe\x25\xe1\x55\xe2\x88\xba\x76\xcd\x6d\x32\xb7\xa2\x96\x37\xb7\xa1\x41\x55\x9a\x73\xdb\xe7\x10\x8d\x6d\x52\x9e\xba\x2f\x95\xbc\xb5\x4d\xc6\x36\xb4\xb2\xd6\x46\x5a\x7a\x89\xca\x81\x5c\xb1\x83\xd9\x56\x42\x67\xfb\x08\xad\x36\x01\x2b\x8f\x42\x78\x39\xab\x24\x5f\x16\x19\x77\xed\x2d\x2b\xbf\x7d\xb0\x36\x86\x8a\x32\x64\xaa\x39\x79\x07\xf8\x59\x79\xfb\x2d\x47\xcf\xb1\xd8\x80\xbc\x01\x7c\x8d\x6e\x8f\x78\xcd\x4b\x80\xeb\xe6\x93\x02\xbc\x6a\x68\x68\x63\x3d\x59\xb5\x88\xfc\xed\x79\xbe\x83\xd3\x69\x7f\x59\x84\x10\xfe\xa7\x02\x1e\xf3\x87\x98\x99\x0f\x55\xbb\x94\x85\x53\x43\xe9\x32\xcf\xc7\xf8\x46\x2c\xe5\xa9\x0f\x3f\xb7\x23\x4e\x04\x50\x48\x94\x14\xfe\xe1\xad\x53\xd6\xef\xcc\xce\xc0\x1b\xf1\x3b\xc6\x2d\x21\xbb\x8f\xc2\x4e\x38\x61\x70\x29\xaa\xe6\x17\x44\x1e\x36\x2b\xcd\x82\x57\xb8\x79\x5a\x40\x48\xc1\x10\xdb\x94\xbd\x6c\xec\xdd\xac\x88\xc8\x65\x21\x5b\x85\x34\x4c\x9e\xba\xce\xd9\xe3\x93\x92\x8d\xc7\x5b\xc8\x5a\x04\x0e\xa6\x1d\x8a\x6a\x7d\x90\x7c\x5f\xcc\x9c\xaf\xf4\x8b\x74\x71\x51\x66\xe4\xd1\x55\xdb\x1e\xa9\x05\x7c\xd3\x05\xa5\x92\x17\xf3\x92\x1a\x9d\x64\x57\x27\x39\x68\x42\x56\xc8\xc2\x0a\xb9\xd9\x98\x90\x8d\x09\x59\x1d\x97\x63\x42\x56\xa5\x36\x26\x64\x8f\x38\x21\x33\x66\x54\xc0\x7b\xbc\x1d\x61\xcc\xf7\x0e\x63\x84\x19\x8c\x28\x66\x44\x31\xb2\xcd\xc3\x44\x31\xff\x81\x79\x7f\xc7\x3d\xdb\x08\x67\xd0\x08\x67\x46\x38\x33\xc2\x19\x54\x86\x33\xdc\xe5\x9d\xe0\x04\x8f\x88\xe6\x7b\x47\x34\x99\x25\x8c\xa0\xe6\xb1\x82\x1a\xf8\xc7\x92\xf2\x97\xb9\x47\x70\x33\x82\x9b\x7a\x2e\x47\x70\x53\xa5\x36\x82\x9b\x3f\x1a\xb8\xe1\x27\xbd\x47\x60\x93\x07\x94\xc2\x6f\x7c\x1b\x7e\xbe\x5f\xf4\xf3\xf0\x10\x0e\x37\x87\x11\xdd\x3c\x56\x74\xf3\x7d\xa0\x9a\x11\xcc\xa0\x11\xcc\x8c\x60\x66\x04\x33\xa8\x0c\x66\xc2\x28\xfc\xeb\x80\x67\xe5\x4a\x07\xbd\xdd\xdf\x83\xb0\x1e\xe4\xc9\x41\x0c\x6a\xa4\x51\x53\x75\x6a\x41\xc5\xb6\x15\xd7\x82\x84\x05\x21\x66\x14\xe6\x25\x38\x69\x50\xbd\xbb\xe2\xf4\x93\x8e\x2d\x98\xac\xcc\xbc\x99\xbd\xca\xc1\x2e\xd7\xf7\xab\x8e\x01\xb0\x48\x64\x45\x59\xf1\x26\x32\x40\x00\x58\xbb\xc4\x55\xd2\xb9\xd7\xfb\x5b\x9c\x91\x71\xd7\x37\x1f\xbb\x03\x08\xe3\xcd\x73\x99\x7c\x8e\xb7\xd3\xce\x76\xe2\xcc\x24\xbd\xb6\xf0\xb4\xc7\x80\x45\xaa\x5d\x60\x47\xa7\x21\xeb\xb0\x49\xa3\xef\xec\x30\x62\xd3\x51\xeb\x06\x40\xd4\x65\xc4\x21\x50\x53\x87\x71\x07\x81\x56\x5d\xe4\x1d\x02\x7f\xf5\x92\xb7\x17\x48\x73\x1d\x59\x0b\x1b\x11\x13\xf9\xc0\xa9\x4a\x52\xba\x01\xba\x01\x46\x3e\x91\xeb\xe9\x65\x27\x10\xd8\x41\xe7\xbd\x90\xe2\x3e\x35\xbd\xef\x81\xeb\x15\xed\x80\x45\x3b\x28\xbb\x11\xb0\x82\xc2\xed\x2f\xd6\xde\x87\xd6\xef\x65\xf4\x7a\xd5\xdb\x2a\x62\x7d\x18\x90\x51\xff\x58\x94\x1d\x5a\x67\x0a\x5d\x3c\x99\x35\x9d\x30\x55\x57\xfa\x40\x0c\x4e\xce\x38\x90\xb1\x1a\x82\xc3\x6d\x19\xdb\x6a\x1c\xd4\xbf\x2c\x50\x02\xb7\x45\x79\x0c\x77\x17\xe9\xcb\xf8\xc7\xd2\x23\xf3\xc5\x46\x2d\x58\x41\x85\x57\xc4\x39\x72\xb6\xec\xcd\x7e\xbd\x33\x6a\x07\x07\x41\x15\x94\xd4\xdc\xc3\x64\x93\xc4\x76\x41\x93\xb3\x42\xcd\x6e\xde\xba\x0a\xad\xaf\xef\x23\x87\x37\x93\xdb\x71\xd6\x4e\xa1\x3e\xe5\xd0\x13\xd4\x84\x93\x28\xee\x02\xbe\x63\xc8\x42\xcf\xc2\xc0\x7a\x09\x55\xe7\x3b\x67\x6e\xd7\x95\x9b\xe0\xcc\x3a\xe0\x0d\x2d\x90\x67\xf8\xeb\x9f\x54\x81\xb8\xa8\xd0\xd2\x56\x47\xa9\x44\xe9\x90\x05\xf6\xba\x8c\x75\x7c\x01\x5c\x75\x1d\xcb\xbe\x63\xd9\xb7\x49\x9b\x63\xd9\x77\x2c\xfb\x76\x2d\xfb\x9a\xbe\xc9\xd1\xf7\x0b\x18\x19\xcd\xf7\x12\x3b\xf3\x1b\x58\x77\xf5\xc8\x27\x46\x45\x57\xde\xdb\x36\xd1\xe8\x1e\x74\xaa\x91\xa5\x02\xab\xcc\x9e\xdb\xe9\x72\x57\xdb\xb5\x96\xba\xa6\xf5\xc8\xbf\x9f\xb0\xe9\xb0\x39\x6b\x14\x41\xdc\x78\xb5\xc1\x8d\x5f\x48\xb0\xe0\x43\x98\xfb\xdb\x2e\x3d\xc1\x41\xc4\x74\x91\x56\x2f\x50\xeb\x0d\xb3\x6e\x62\xbc\xd9\x0c\x75\x65\xa0\x71\x3a\xf9\x67\x5e\x86\x9a\xce\xba\x3b\xed\x87\x9e\xea\x9e\x17\x35\x0d\x86\x3e\x2b\x0e\xe1\xfe\x8a\xfd\x46\x47\x73\xb2\x23\x31\x8c\xa3\x31\xec\xdc\x58\x79\x5e\x60\x46\xbd\xe3\x34\x59\xf1\xaf\x53\xc8\x63\x27\x17\xd5\xeb\x26\x0f\x5d\xc4\xc7\x1b\xfa\x4f\xb2\xed\xda\x3b\xc2\xc0\xc4\xab\x53\x48\x06\xa8\x47\x93\x7e\x54\xce\x31\x63\x37\x51\xec\xf7\xa3\x72\xbc\xe1\xbc\xf4\x52\x89\x22\xe4\x79\x84\xb1\x9f\x23\x9f\x18\xe8\xcc\x8d\xb6\x51\x37\x2d\x7b\x59\xfb\xfb\xbe\x1e\x4b\x08\xd4\xfb\x8c\xd2\xc3\x5b\xc4\x25\xa3\xdf\xdf\xdc\x1c\xa2\xe2\x3e\xf7\x1e\x67\x4a\x4a\x64\x99\xaa\xae\xbe\xbf\xdd\xf1\xb8\xf2\x9d\x0a\x87\x68\xa0\xd7\x12\x1f\x9e\xfd\x58\xdc\xde\x5e\xed\x68\x19\x44\x37\xa2\x0e\x01\x43\x47\xb1\xfa\x8c\xce\x47\xa7\xaf\x11\x74\xb6\x2a\x29\xa7\xad\xe8\xc1\x19\xea\x42\x95\x2a\xbd\x59\xe8\x32\x0f\xa4\x70\xbd\xd7\x56\x30\x78\x21\x7b\x98\x31\x63\x59\x5d\x2d\x38\x76\xf8\x76\xd2\xe3\xb7\xe2\x4a\xd8\xbd\x17\x2b\x4e\xa2\xdf\xc8\xe3\xb4\xde\x8d\xd2\xd7\xbd\x58\x6f\xae\xa6\xd1\x6a\x75\xab\x35\xc1\xbc\xd1\x70\x1b\x20\x42\xae\xb2\xd1\x76\xbf\xa9\xed\x56\x33\x8b\x6f\x83\x1c\x1e\xb7\x39\xe7\x5a\xfc\x23\xe0\x88\x71\xb1\x20\xf3\x62\xb9\x28\xcf\xe2\x00\xb5\x5c\x5d\x64\x7d\x54\xfd\x93\x2a\x03\x56\xb9\xf3\x6f\x5a\x55\xf4\xdb\x50\xda\x2e\x7d\xd9\xa0\x99\x25\xc3\x2b\x34\x3b\x1a\x28\x24\xc4\x97\x5f\x94\x64\x30\x1f\x08\xab\xcf\x14\xc8\xcf\xcf\x04\x81\xfe\x45\x82\xca\xb1\x01\xc3\x1d\xaa\xba\x07\x32\xe8\xa3\x4d\x65\xab\xf2\xf5\x3c\xd7\xa2\x8d\xf1\x6e\xfb\x62\xa5\xc6\x79\x27\xa1\xf0\x85\x99\x4e\xba\x4e\x62\x1c\x32\x60\x84\xdf\x5f\x9b\x44\x5e\x14\x98\x3f\x6d\x6b\x50\x94\x75\xb1\x17\x72\xeb\x24\xd9\x70\xbf\xcd\xff\x66\xfc\x1f\x37\xea\x4f\x36\x69\x2d\x69\xcd\x5e\xb1\x8d\x95\x1d\x23\x1e\xbb\x16\x77\xdc\xca\xbf\x12\xf9\xd7\x86\x6e\xb4\xef\x9b\xed\x4a\xe4\xa2\x7d\xed\xe8\xbb\x17\xa8\x07\x62\xe3\x50\xed\x83\x39\xf2\x53\x3e\x3f\x3d\xd0\xd9\xe9\x8c\xbc\xd9\x13\x0f\x7a\x5e\x7a\x37\x54\x79\x9f\x7e\xb0\x33\xd2\xb9\x97\x34\x6d\x30\x0e\x79\x2e\x3a\x1f\xa8\xb2\xad\x3f\xd8\x59\xe8\x6c\x88\x9a\x23\x04\xc3\x9f\x7f\xce\xe5\xaa\x1c\x04\x18\xec\xcc\x73\x55\xae\x41\xc7\x32\x9f\x4a\x28\xcc\x57\xe5\xfc\xc0\xf0\x67\x3f\x0b\x5a\xdc\xeb\x68\xfa\x59\xcf\x5d\x40\x2e\x9f\x6a\x18\xec\x0c\x73\x41\x8d\x95\x83\x53\xfb\xd4\xe2\x3e\x07\x33\x2b\xd1\x7c\x70\x62\xd0\xf3\xc9\xf9\x42\x08\x87\x33\xfe\xb0\x6c\xf0\x3a\xee\xd8\xf7\x86\x7d\xab\xd7\x8a\x74\x7c\xfb\x84\xff\xff\xee\x7f\x01\x00\x00\xff\xff\x0a\xd1\xa8\x87\x15\x8b\x00\x00")

func v2SchemaJSONBytes() ([]byte, error) {
	return bindataRead(
		_v2SchemaJSON,
		"v2/schema.json",
	)
}

func v2SchemaJSON() (*asset, error) {
	bytes, err := v2SchemaJSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v2/schema.json", size: 35605, mode: os.FileMode(420), modTime: time.Unix(1433877279, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"jsonschema-draft-04.json": jsonschemaDraft04JSON,
	"v2/schema.json": v2SchemaJSON,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"jsonschema-draft-04.json": &bintree{jsonschemaDraft04JSON, map[string]*bintree{
	}},
	"v2": &bintree{nil, map[string]*bintree{
		"schema.json": &bintree{v2SchemaJSON, map[string]*bintree{
		}},
	}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

