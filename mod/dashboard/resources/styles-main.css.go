package resources

import (
	"bytes"
	"compress/gzip"
	"io"
	"reflect"
	"unsafe"
)

var _styles_main_css = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x7b\x6f\xe3\xb8\x11\xff\x2a\x42\x0e\x8b\x8b\x0b\x2b\x90\xe5\x47\x6e\x25\x1c\xd0\xbd\x5c\xd3\x16\xe8\x15\x05\xee\xfe\x28\x50\xf4\x0f\x5a\xa2\x6d\x62\x65\x51\xa5\xe8\x3c\x56\xc8\x77\xef\x90\x14\x25\x52\xa2\x62\xc9\xd9\x2d\x8a\x8b\x17\x8b\xd0\x14\x39\x9c\xf7\xfc\x86\xf2\x96\xa6\xcf\xd5\x11\xb1\x3d\xc9\xa3\xe0\xe5\x06\xf3\x24\xf5\x13\x9a\x73\x44\x72\xcc\xaa\x2d\x4a\x3e\xef\x19\x3d\xe5\x62\x32\xa3\x2c\xfa\x6e\xb7\xdb\xc5\x5b\xca\x52\xcc\xa2\x45\xf1\xe4\x95\x34\x23\xa9\xf7\x5d\x9a\xa6\xf5\xac\xcf\x50\x4a\x4e\x65\xb4\x2e\x9e\x60\xe6\xc9\x2f\x0f\x28\xa5\x8f\x11\xdb\x6f\xd1\x75\x30\x17\xff\x6e\x16\xab\x8f\x41\x38\xf3\x02\x4f\x10\x58\xc2\xba\x1d\x9c\xe7\xef\xd0\x91\x64\xcf\xd1\xd5\x5f\x70\xf6\x80\x39\x49\x90\xf7\x77\x7c\xc2\x57\xf3\xe6\xfb\xfc\x13\x23\x28\x9b\x97\x28\x2f\xfd\x12\x33\xb2\x8b\xe9\x03\x66\xbb\x0c\xa8\x1f\x48\x9a\xe2\x5c\x9d\x47\xbe\x90\x7c\x1f\xd5\xcc\xc0\x4c\xec\x1f\xe9\x17\xdf\xfd\xa8\xa0\x25\xe1\x84\xe6\x11\xc3\x19\xe2\xe4\x01\xc7\x27\xa0\x0c\xd4\x33\x9c\xf0\x28\xa7\x39\x8e\xfd\x47\xbc\xfd\x4c\xb8\xdf\x7f\x20\xc8\x3a\x66\xcb\xde\xe4\x0b\xaa\x6a\xe5\x85\x8b\xdb\xcd\xa7\xbb\x98\xe3\x27\xee\xa7\x38\xa1\x0c\xc9\xd3\xd5\xa2\xe8\x20\xe4\x99\xa3\x08\x25\x82\x95\xaa\xbb\x0a\x8c\x80\x59\x06\x66\x79\x21\x79\x71\xe2\xff\xe2\xcf\x05\xfe\x51\x2c\xfa\x77\x65\x28\x9a\xe4\x25\xe6\xb5\x6e\x43\xf8\x6f\x2a\x7e\x3d\xd3\xa6\x0b\x3a\xd6\x6a\xac\x00\x3a\xc2\xd1\x42\x7c\x2d\x50\x9a\x82\xba\xfc\x0c\xef\xb8\xb4\xa6\x9e\x60\x64\x7f\x50\x33\x07\x2c\x87\x21\x8c\x5f\x0e\x61\xd5\x12\x08\xc3\xae\x55\x7f\xa5\x27\x96\x60\xef\x57\x30\x9e\xf7\x0f\x46\xaf\xe6\x53\xec\x2c\x29\x3d\xaa\xc3\xd6\x41\x10\x6b\x77\x05\x31\xc3\xa0\xe5\xac\xf1\xdf\xed\x89\x73\x9a\x57\x29\x29\x8b\x0c\x3d\x83\x4a\x84\xd6\xfc\x6d\x46\x93\xcf\xcd\xda\x0d\xe8\x66\x21\xd8\x54\xc4\xc0\x1d\x60\xcf\x11\x14\x63\xa8\x61\xa5\xa5\xa8\xcf\x5e\xc1\xd9\x92\x54\x2d\xf8\xe2\x66\x15\xfe\xb0\xbe\x5d\xac\xc2\x8f\xca\xa6\x28\x23\xfb\x3c\x4a\x70\xce\x31\x8b\x1f\x0f\x84\x63\xbf\x2c\x50\x82\xc1\xc2\x8f\x0c\x15\x31\xd8\x57\x48\x98\xd5\x0b\x8f\xe0\xb6\x19\x8e\x93\x13\x2b\xc1\x3b\x0a\x4a\xe4\xbe\x5e\x74\x71\x06\xaa\x28\x10\x03\xb2\x1d\xb3\x09\x06\xbf\x82\x83\xc6\x3e\xed\xcf\xf5\x26\xb4\xd6\x4d\x1f\x72\xba\x5d\xa0\x9c\x2e\x5c\xaf\xe7\xfa\xff\x4d\x38\x9b\xbb\x7d\x32\x5c\xcf\x2c\xab\x99\x63\xbf\x3c\xa2\x2c\xab\x0c\x37\x6b\x4d\xed\x2d\x82\x9e\xcf\x5a\x84\xfc\x82\x11\x60\xf9\xd9\x91\xc3\xc0\x6a\x3f\xdd\x7d\x8a\x8d\x8c\x26\xad\x57\x8b\x22\xdc\x6a\x79\x86\x4b\x4d\x5c\xc7\x6a\xff\x8c\x30\xbc\xdd\xa0\xb4\xde\x54\xd0\x42\xc4\xb6\xb1\x2c\xfa\x6e\xb9\x5c\x3a\x82\x50\xcb\xb7\x90\xc2\xea\xf4\x84\xb6\xe0\x09\x27\x8e\x63\x4e\x8b\x68\xf9\x11\x1e\x7d\xf1\x09\xa4\x83\xa7\xe8\x23\x7c\x4c\x41\x3a\x41\x6c\x58\x28\x90\xaa\x17\x5a\xeb\xe6\x04\x1d\x27\x32\x0d\x59\x0c\xfb\x98\x31\xca\x3c\x7b\x4e\x54\x07\x70\x45\x9d\xd3\xee\xef\x97\x77\xab\xa5\x15\x27\xb7\x10\x27\xa6\xfb\x08\xed\x3a\xfd\x54\x3d\xe8\xfa\x69\x3d\x5b\xf6\x26\x3b\xcc\xe5\x94\x27\x87\xea\x91\xa4\xfc\xa0\x62\x55\x87\xa5\x19\x16\x32\x78\x76\x94\x1d\x23\x46\x39\xe2\xf8\x7a\xb5\x4e\xf1\x7e\xe6\xd0\x6d\x9d\x08\x84\x8a\xfd\x75\x9b\x18\x64\xfe\x93\xba\xec\x18\x4f\x2c\x0c\x62\x95\x0d\x85\xb9\x6c\xee\x6c\x56\xe5\xaa\x4a\x92\xba\xbd\xed\x2e\x55\x4f\x3d\x97\x70\x72\xc7\x39\xe2\x62\x51\x55\xf3\x11\xf4\x88\x8b\xa7\x4e\xda\x43\x9c\x0b\x03\xef\x08\x3b\x56\x86\x42\x0c\xba\x07\x8c\xc0\x67\xb5\xda\x83\xe0\x83\xa3\x8a\x4e\x2d\xc4\x16\xe9\x1b\x99\xf5\xac\x58\xc1\x18\x9b\xea\xd7\xb6\xdd\x8b\xc0\x01\x57\xbc\x16\x49\x19\xb1\xb9\x14\x15\xd8\x55\x03\x95\xd1\xe7\xd2\x4d\xfd\x12\xa6\xaf\x83\x0f\x73\x41\x6a\x66\xce\x09\x09\xe6\x02\xbf\xcc\x66\xae\x23\x14\xe5\xf6\x24\x41\x5d\xd0\xf0\x02\xb9\xc9\x13\xdb\xed\x8d\xe0\xb7\xd3\x37\x5d\x7a\x5a\x7f\x83\x16\xbb\xbf\x6d\x47\x32\x28\x2f\x51\xc1\xe8\x9e\xa4\xd1\xcf\xff\xfc\xeb\x11\xed\xf1\x6f\x3a\x3c\x6e\x7e\x21\x09\xa3\x25\xdd\xf1\x9b\x86\x5a\xc9\x11\xe3\x77\x42\x57\x25\x67\x3f\x7e\x2f\x48\xc2\xe7\xfb\xb9\x87\xf3\xd4\x98\x4e\xe5\x07\xa6\xff\x5c\x6f\xfc\x4d\x20\x93\xa0\xc9\x98\x02\x61\xaa\x60\x91\xee\x1c\x38\x1c\x46\x23\x39\x1f\x32\xea\x89\xd3\xf6\xfb\x93\xc6\x76\x3a\xb6\x85\xc3\x19\xbe\x77\xa1\xa7\x09\x9e\x3c\x8e\xb6\x19\xae\xbe\x26\x31\x8f\x0b\x17\xf6\x78\xaa\xc0\x5b\x9b\x7b\x4e\x45\x81\x59\x82\x4a\x6c\x26\x68\x28\x87\x16\x9c\x90\x58\xa6\xa1\x2e\xe1\x48\x07\x62\x6b\x08\xa6\xb4\x69\x03\xb2\x85\x81\x84\x5a\x34\x63\xa1\xb8\x40\x57\x8a\xcd\x66\xe3\xe0\x5d\x8d\xd3\xca\x82\x38\x3f\x4c\xe1\x69\xd3\x83\x89\x2e\xae\x36\x5d\x74\x19\x74\x81\x91\xd2\x1e\x2d\xbe\x65\x15\x31\x24\x57\x13\x9c\x67\x76\x72\x5b\x2e\x3f\x0c\x69\xa9\xd9\x51\xb9\x11\xc8\xe0\x7a\xe3\x2c\x51\x6e\x75\x11\x15\x55\xdc\xac\xa0\xe0\x8e\x43\x6c\x0a\xc4\x41\x01\x1a\xdb\xac\x06\xce\xb3\x6b\x6f\x8c\x20\x9d\x97\xdc\x4f\x0e\x24\x4b\xe7\x83\x76\x37\x57\x55\x96\x7d\x16\xaf\x53\xcf\xd0\x08\xe2\xed\xa2\xaa\xef\x21\xdd\xbe\x53\xd7\x2b\x86\x1f\x08\x7e\xac\xbf\xa5\xd0\x15\xd0\x7d\x0f\x43\xf5\x0b\xb9\xa2\xdb\x66\x9b\x1a\x52\xc9\xf0\x72\x43\xa5\xd5\x1b\x9a\xd0\x26\xd4\x28\x20\x7d\xc2\x9f\x61\x74\x84\xaa\xa9\xe3\x47\x06\x80\x05\x46\x24\xb7\x20\x1b\x85\x94\x00\xcb\x01\xe1\xc6\x43\xf3\x8e\x7d\xe9\xa9\x6e\x07\x17\xeb\xe0\x58\xc6\xc3\x4f\x1c\x7b\x39\x39\x0a\xbd\xef\x4e\xb9\xf4\xa1\x08\x43\x42\x02\x1c\xe9\xd3\x13\x8f\xc7\x2d\x9b\x60\x28\xeb\x8b\x7f\xc4\x65\x09\xe5\xa6\xb2\xbb\xad\x70\x9a\xed\x6d\x92\x0a\x89\x97\x1e\xb2\x12\xd6\x72\x65\x5b\xd9\xdd\x72\x99\xd8\xae\xeb\xdd\x6a\x98\x91\x92\x57\x96\xeb\x4c\xbe\x5d\x68\xae\x27\x64\x49\x33\x0b\x58\xcf\x67\xdd\xe7\xd7\x43\x95\xbb\x70\xea\xe8\x32\xfe\xf4\xe9\x7e\x79\x7f\x3f\xb0\x1b\xdd\xa4\x84\xc1\x4e\x0a\x4d\x50\x07\x9e\x0f\xec\xe0\x4c\x5d\x42\x68\x45\xc3\xc1\x1c\x7b\xe5\xc3\xbe\x5a\x3c\x90\x92\x6c\x49\x26\x1c\x5c\x0e\x41\x89\x00\x28\x32\x28\xe1\xe9\xab\xcc\x2b\x1a\x95\x59\xe2\x54\xca\x92\xed\xdc\x6b\xd6\x81\x40\x72\x75\xf0\x23\x4e\x93\x1c\x9b\x27\x4a\x4e\x05\x7c\x19\xa5\xe7\x31\xb4\xa0\xf2\x8d\x60\xa4\xd6\xa6\x20\xd1\x69\xf2\x3b\xba\x6b\xbc\xbf\xa6\xc9\xe8\x63\xd9\x98\xa1\xce\xf4\xda\xc1\xdb\xfe\x41\xf8\x83\x39\x5d\x6b\x56\x0e\xcf\x68\xd6\xdd\x5a\xd6\x39\x53\x74\x3b\x8e\xa6\xd0\xcd\xa4\x15\xae\x3e\xfc\xc1\x48\x57\x39\xc9\x9f\x26\x64\x1a\x6f\x0c\x25\xc0\x7d\xf8\x1c\x1d\xf1\xa0\x67\x20\x1d\xb4\x80\x5c\xbc\x8d\xa5\xab\x33\xe6\x10\xf7\x02\xdb\x04\x19\x1b\x74\xbe\x29\x45\x2d\x4b\xe5\x46\xb5\x72\xbb\xdd\xd6\xcb\xe0\xb0\x6f\x62\x82\x96\x73\x38\xc1\x92\x31\x34\x64\xbc\xb5\x64\x84\xc9\x49\x22\xc2\xfa\x11\x12\x8e\xf0\x4d\xcf\x7a\xe4\x17\x88\x1f\xaa\xbe\x78\xaa\xfd\x15\xec\xeb\x0a\xad\x04\xae\x35\x22\x6e\xe4\xd6\xfd\x7a\x30\xe9\x54\x4f\x5e\x8d\x7e\x05\x4c\x3f\xfe\xf8\x12\x3d\x60\x87\xb0\x8a\x85\xf5\xaa\x91\xd6\x5f\xb7\x97\x0a\x52\xd6\xde\x6d\xbb\x23\x18\x04\x71\x3b\xa6\xfa\x67\x2b\xf2\xb7\x67\x34\x37\x8a\x9e\x65\x40\x45\x77\x13\x8e\x27\x6c\x84\xec\x10\x9b\x52\x0b\x6f\xa5\xe7\x60\xf3\x75\x65\xda\x98\x62\xc8\x35\xa5\xbf\x98\xd8\x4d\x7c\x3f\xdf\x99\x2a\x27\x9e\x8c\x0e\xdc\xaf\x56\x2c\xcf\x95\x6b\x15\x6b\x56\xcf\x35\x41\x56\x0f\xfe\xce\x27\x2c\x17\x0d\x12\x62\x18\xd9\xd0\xa7\xc9\x59\x26\x48\xbe\x03\xf0\x45\x01\xd6\xcf\xaf\xfe\x46\xb6\x58\x61\x4f\xef\x17\x9a\xd3\xab\xf9\x1d\x3d\x31\x82\xd9\xfc\x08\xdf\xe4\x05\xf8\x50\xfb\x6e\xde\x91\x03\x0b\xbe\xbc\x25\xef\x9b\xa7\x77\xab\xa9\x6f\xc8\x63\xc0\xa4\x02\x1a\x5c\x60\x80\x4b\x8b\xda\x58\xc5\x8e\x20\xd1\x28\x7b\x7c\x8d\x74\xa2\x63\x7c\x2c\xb8\xbe\x62\x69\xda\x11\x43\xcb\xfd\x77\x13\x43\x00\x74\xb2\x28\xea\xe8\xd1\x60\x61\x98\xfd\x6e\x87\x20\xa4\x09\xd7\x1f\xe2\xb6\x41\x7e\xf9\xe3\x11\x03\xf4\xf7\xae\x8f\xe8\xc9\x57\xd2\x01\x92\x2d\x9e\x66\xd5\x88\x13\x25\x9c\x6f\x55\x72\xa9\xb8\x92\x8c\x8c\x48\x5f\x6a\xf6\x5b\x34\x76\xe1\x60\x63\x17\xfe\xef\x1a\xbb\x71\xc0\xac\x55\x47\xf0\x4e\x54\xd1\xa9\x24\x4d\xdd\x78\x63\x36\xd1\x3e\xf5\x8d\xee\x0a\xfe\xff\xf4\x38\xe8\x52\x3d\xcd\xfe\x8e\x14\xd2\x66\x30\x98\x78\x43\x06\x5b\x9f\xf1\x37\x97\x2a\xd7\x4d\xe1\x3d\xb7\x7b\xac\xb7\xbe\x9d\xce\x65\x35\x64\x9c\xeb\x4c\xe3\xae\x4f\xc5\xc9\xdb\x1b\x7a\xd8\xcb\x60\x83\x83\x40\x0f\x34\x48\x85\xb9\x99\x2a\x39\xe2\xa5\x77\x08\xcd\xa2\xea\xb7\x9d\x82\x78\x35\x81\x78\x7d\xef\x41\x99\x03\x17\xcb\x77\x7f\x6d\xc7\xb6\xd8\xd8\xaf\x9e\xdd\x74\xac\xeb\x14\xca\x7c\x80\x78\x47\xf7\x8f\x31\x34\x2a\x0c\x9b\xc6\x79\x11\x76\x3a\xf8\xd5\x10\xb7\xae\x53\x54\xf7\xdb\xe9\x7c\xc7\x6f\x37\xda\xdf\xe5\x72\xf9\x9a\x4a\xd5\x78\x0f\x68\xf5\x50\x5d\xfe\xab\xa3\x4e\xe3\xdf\x40\xb7\xba\xa2\xea\x2e\xd9\x84\xe2\x23\x99\x32\xc7\xc6\xcf\xb9\xdc\x27\x6f\xd4\xed\x66\x73\x23\xab\xce\x97\x23\x66\xbe\x98\xfa\x1a\xf8\x5a\x71\x6a\xbe\xcf\xc8\x10\xc7\x79\xf2\x7c\x2e\xb5\x99\x22\xaa\xbb\xd9\xdf\x65\x13\xd7\x15\x53\xbb\xe9\x7f\x4e\x10\xf3\x55\x23\x53\x1b\x2f\x81\x71\x57\x67\xc5\x56\x1d\xf2\xcd\xaf\xb5\x2e\x38\x0f\x92\xb6\xeb\xd2\x39\x59\x05\x41\x18\x5e\x42\x8f\x42\x9d\xdc\xbb\x7e\x2d\x73\x7f\x7f\x17\x04\x03\x3d\xfc\xeb\x24\xf7\x0c\xe3\xdc\x41\x31\x08\x7e\xfe\x29\x5c\x4d\xa2\x58\x60\xb0\x94\xf8\x55\x9d\xf9\x56\xce\x7a\x11\xe6\xf8\x99\x9c\xf5\xb2\x74\x10\x08\x0e\x9c\x58\xfb\xbe\xff\x80\xb2\x13\x76\xa6\xc8\x89\x4d\x4f\x2f\x4a\x0c\x57\x3c\x83\x55\xfb\x89\x4d\xa5\xa1\x57\x56\xab\xa1\x0c\x66\xbb\xbe\x1b\xfb\xdf\x45\x9f\xd4\xd3\x86\x51\xb7\x4d\x5d\xbe\x67\x3d\x9c\xa9\xbf\xaa\x6c\x41\x01\xf6\xfe\x50\x39\xae\xaa\x27\x1c\xd4\x36\xa5\xef\xd2\xf5\x06\xee\x2b\xde\x9f\xef\x59\x8a\x38\xe3\x7d\x0a\x90\x38\xbc\x6f\x62\xcb\x36\x90\x7f\xd7\x3a\xfd\x8e\x82\x37\x56\xc6\x98\x00\x8b\xba\x70\xde\x6e\x10\xfe\x1b\x00\x00\xff\xff\xf8\x98\x92\xa5\xe9\x2f\x00\x00"

// styles_main_css returns raw, uncompressed file data.
func styles_main_css() []byte {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&_styles_main_css))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(_styles_main_css)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, gz)
	gz.Close()

	return buf.Bytes()
}


func init() {
	go_bindata["/styles/main.css"] = styles_main_css
}
