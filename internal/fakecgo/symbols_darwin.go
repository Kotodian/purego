// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

package fakecgo

//go:cgo_import_dynamic purego_malloc malloc "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_free free "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_setenv setenv "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_unsetenv unsetenv "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_attr_init pthread_attr_init "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_create pthread_create "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_detach pthread_detach "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_attr_destroy pthread_attr_destroy "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_attr_getstacksize pthread_attr_getstacksize "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_pthread_sigmask pthread_sigmask "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_abort abort "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_sigfillset sigfillset "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic purego_nanosleep nanosleep "/usr/lib/libSystem.B.dylib"
