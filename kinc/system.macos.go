// +build cgo,darwin

package kinc

// #cgo darwin CFLAGS: -I${SRCDIR}/../Backends/System/Apple/Sources
// #cgo darwin CFLAGS: -I${SRCDIR}/../Backends/System/macOS/Sources
// #cgo darwin CFLAGS: -I${SRCDIR}/../Backends/System/POSIX/Sources
// #cgo darwin CPPFLAGS: -DKORE_MACOS=1 -DKORE_POSIX=1
// #cgo darwin LDFLAGS: -framework Foundation -framework AVFoundation -framework IOKit -framework Cocoa -framework AppKit -framework CoreAudio -framework CoreMedia -framework CoreVideo
import "C"
