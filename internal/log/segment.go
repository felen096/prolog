package log

import (
	"fmt"
	"os"
	"path"
)

type segment struct {
	store                  *store
	index                  *index
	baseOffset, nextoffset uint64
	config                 Config
}

func newSegment(dir string, baseOffset uint64, c Config) (*segment, error) {
	s := &segment{
		baseOffset: baseOffset,
		config:     c,
	}

	var err error
	storeFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprint("%d%s", baseOffset, ".store")),
		os.O_RDWR
	)
}

CDM06PrdApp11- CDM01P
CD1PPF4BECA9AF5