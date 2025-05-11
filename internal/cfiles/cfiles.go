package cfiles

// Go's CGO toolchain has a limitation that it only compiles C files in the package directory. In order to have the Go
// toolchain compile C files in a subdirectory, we have to create a new Go package in that subdirectory and import it
// somewhere in our code.

import "C"
