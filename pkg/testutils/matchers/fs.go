package matchers

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/onsi/gomega/types"
	"github.com/spf13/afero"
)

/*
	BeNamedFileOrDir succeeds if actual is a path with the base of the expected name. Actual must be a string representing a path being checked.
*/
func BeNamedFileOrDir(name string) types.GomegaMatcher {
	return &beNamedFileOrDirMatcher{
		expected: name,
	}
}

type beNamedFileOrDirMatcher struct {
	expected string
}

func (b *beNamedFileOrDirMatcher) Match(actual interface{}) (success bool, err error) {
	pathString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("actual %v of type %T must be of type string", actual, actual)
	}
	actualName := filepath.Base(pathString)
	return actualName == b.expected, nil
}
func (b *beNamedFileOrDirMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path or dir named\n\t%#v", actual, b.expected)
}

func (b *beNamedFileOrDirMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path or dir not named\n\t%#v", actual, b.expected)
}

/*
	BeTempPath succeeds if actual is a path contained under the TempDir of the given Fs. Actual must be a string representing a path being checked.
*/
func BeTempPath(fs afero.Fs) types.GomegaMatcher {
	return &beTempPathMatcher{
		expected: os.TempDir(),
	}
}

type beTempPathMatcher struct {
	expected string
}

func (b *beTempPathMatcher) Match(actual interface{}) (success bool, err error) {
	pathString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("actual %v of type %T must be of type string", actual, actual)
	}
	return strings.HasPrefix(pathString, b.expected), nil
}
func (b *beTempPathMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path contained in TempDir \n\t%#v", actual, b.expected)
}

func (b *beTempPathMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path not contained under TempDir\n\t%#v", actual, b.expected)
}

/*
	BeSubPath succeeds if actual is a path contained under expected. Actual must be a string representing a path being checked.
*/
func BeSubPath(expected ...string) types.GomegaMatcher {
	return &beSubPathMatcher{
		expected: filepath.Join(expected...),
	}
}

type beSubPathMatcher struct {
	expected string
}

func (b *beSubPathMatcher) Match(actual interface{}) (success bool, err error) {
	pathString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("actual %v of type %T must be of type string", actual, actual)
	}
	return strings.HasPrefix(pathString, b.expected), nil
}
func (b *beSubPathMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path contained under path\n\t%#v", actual, b.expected)
}

func (b *beSubPathMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be a path contained under path\n\t%#v", actual, b.expected)
}

/*
	BeFile succeeds if actual points to the same pointer as expected. Actual must be a *os.File being checked.
*/
func BeFile(expected *os.File) types.GomegaMatcher {
	return &beFileMatcher{
		expected: expected,
	}
}

type beFileMatcher struct {
	expected *os.File
}

func (matcher *beFileMatcher) Match(actual interface{}) (success bool, err error) {
	file, ok := actual.(*os.File)
	if !ok {
		return false, fmt.Errorf("BeFile matcher expects an *os.File")
	}

	actualPtr := reflect.ValueOf(file).Pointer()
	expectedPtr := reflect.ValueOf(matcher.expected).Pointer()

	return actualPtr == expectedPtr, nil
}

func (matcher *beFileMatcher) FailureMessage(actual interface{}) (message string) {
	actualPtr := reflect.ValueOf(actual).Pointer()
	expectedPtr := reflect.ValueOf(matcher.expected).Pointer()

	return fmt.Sprintf("Expected\n\t%#v\nto point to Address\n\t%#v", actualPtr, expectedPtr)
}

func (matcher *beFileMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	actualPtr := reflect.ValueOf(actual).Pointer()
	expectedPtr := reflect.ValueOf(matcher.expected).Pointer()
	return fmt.Sprintf("Expected\n\t%#v\nnot to point to Address\n\t%#v", actualPtr, expectedPtr)
}
