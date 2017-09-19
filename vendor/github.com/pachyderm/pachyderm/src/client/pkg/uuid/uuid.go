package uuid

import (
	"strings"

	"github.com/satori/go.uuid"
)

var (
	// UUIDWithoutDashesLength is the length of the UUIDs returned
	// by NewWithoutDashes.  It's also the length of commit IDs in
	// Pachyderm.
	UUIDWithoutDashesLength int
)

func init() {
	UUIDWithoutDashesLength = len(NewWithoutDashes())
}

// New returns a new uuid.
func New() string {
	return uuid.NewV4().String()
}

// NewWithoutDashes returns a new uuid without no "-".
func NewWithoutDashes() string {
	return strings.Replace(New(), "-", "", -1)
}

// NewWithoutUnderscores returns a new uuid without no "_".
func NewWithoutUnderscores() string {
	return strings.Replace(New(), "_", "", -1)
}
