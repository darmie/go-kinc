// +build cgo,windows

package kinc

// #cgo windows CFLAGS: -I${SRCDIR}/../../Backends/System/Windows/Sources -I${SRC_DIR}/../../Backends/System/Windows/Libraries/DirectShow/BaseClasses
// #cgo windows CFLAGS: -I${SRCDIR}/../../Backends/System/Microsoft/Sources -I${SRC_DIR}/../../Backends/System/Windows/Libraries/DirectShow/BaseClasses
// #cgo windows CFLAGS: -I${SRCDIR}/../../Audio2/WASAPI
// #cgo windows CPPFLAGS: -I${SRCDIR}/../../Audio2/WASAPI
// #cgo windows CPPFLAGS: -DKORE_WINDOWS=1 -DKORE_MICROSOFT=1 -D_CRT_SECURE_NO_WARNINGS=1 -D_WINSOCK_DEPRECATED_NO_WARNINGS=1 -O2 -g
// #cgo windows LDFLAGS: -ldxguid -ldsound -ldinput8 -lws2_32 -lWinhttp -lstrmiids -lwinmm
import "C"
