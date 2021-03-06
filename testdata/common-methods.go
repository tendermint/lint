// Test that we don't nag for comments on common methods.
// OK

// Package pkg ...
package pkg

import "net/http"

// Description ...
type T int

func (T) Name() string                                     { return "" }
func (T) Error() string                                    { return "" }
func (T) String() string                                   { return "" }
func (T) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (T) Read(p []byte) (n int, err error)                 { return 0, nil }
func (T) Write(p []byte) (n int, err error)                { return 0, nil }

func New() T          {}
func NewXxx() T       {}
func Newxxx() T       {}
func NewXxxXxxXxx() T {}
