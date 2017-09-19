package goth_test

import (
	"testing"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/faux"
	"github.com/stretchr/testify/assert"
)

func Test_UseProviders(t *testing.T) {
	a := assert.New(t)

	provider := &faux.Provider{}
	goth.UseProviders(provider)
	a.Equal(len(goth.GetProviders()), 1)
	a.Equal(goth.GetProviders()[provider.Name()], provider)
	goth.ClearProviders()
}

func Test_GetProvider(t *testing.T) {
	a := assert.New(t)

	provider := &faux.Provider{}
	goth.UseProviders(provider)

	p, err := goth.GetProvider(provider.Name())
	a.NoError(err)
	a.Equal(p, provider)

	p, err = goth.GetProvider("unknown")
	a.Error(err)
	a.Equal(err.Error(), "no provider for unknown exists")
	goth.ClearProviders()
}
