package zstd

import "github.com/klauspost/compress/zstd"

func Decompress(buf []uint8) ([]uint8, error) {
	d, err := zstd.NewReader(nil, zstd.WithDecoderMaxMemory(1<<30))
	if err != nil {
		return nil, err
	}
	return d.DecodeAll(buf, nil)
}

func Compress(buf []uint8) ([]uint8, error) {
	e, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}
	return e.EncodeAll(buf, nil), nil
}
