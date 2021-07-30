package modzy

import "github.com/spf13/afero"

// AppFs is exposed for possible mocking
var AppFs = afero.NewOsFs()
