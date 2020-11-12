// +build cgo,linux

package kinc

// #cgo linux CFLAGS: -I${SRCDIR}/../Backends/System/Linux/Sources
// #cgo linux CFLAGS: -I${SRCDIR}/../Backends/System/POSIX/Sources
// #cgo linux CPPFLAGS: -DKORE_LINUX=1 -DKORE_POSIX=1 -O2 -g
// #cgo linux LDFLAGS: -lasound -ldl
import "C"
