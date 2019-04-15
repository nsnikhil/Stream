/*
 *     Stream for Go  Copyright (C) 2019  Nikhil Soni
 *     This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
 *     This is free software, and you are welcome to redistribute it
 *     under certain conditions; type `show c' for details.
 *
 * The hypothetical commands `show w' and `show c' should show the appropriate
 * parts of the General Public License.  Of course, your program's commands
 * might be different; for a GUI interface, you would use an "about box".
 *
 *   You should also get your employer (if you work as a programmer) or school,
 * if any, to sign a "copyright disclaimer" for the program, if necessary.
 * For more information on this, and how to apply and follow the GNU GPL, see
 * <https://www.gnu.org/licenses/>.
 *
 *   The GNU General Public License does not permit incorporating your program
 * into proprietary programs.  If your program is a subroutine library, you
 * may consider it more useful to permit linking proprietary applications with
 * the library.  If this is what you want to do, use the GNU Lesser General
 * Public License instead of this License.  But first, please read
 * <https://www.gnu.org/licenses/why-not-lgpl.html>.
 */

package stream

import (
	"Collections/functions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf2(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{
			name:  "Test convert int slice to LazyStream",
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Test convert string slice to LazyStream",
			input: []string{"one", "two"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, &LazyStream{elements: interfaceSlice(test.input)}, OfLazy(interfaceSlice(test.input)))
		})
	}
}

func TestLazyStream_LazyFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     LazyStream
		predicate *functions.Predicate
		output    LazyStream
	}{
		{
			name:  "Filter even number from stream",
			input: LazyStream{elements: interfaceSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})},
			predicate: functions.NewPredicate(func(i interface{}) bool {
				return i.(int)%2 == 0
			}),
			output: LazyStream{elements: interfaceSlice([]int{0, 2, 4, 6, 8})},
		},
		{
			name:  "Filter string whose len are greater than 4",
			input: LazyStream{elements: interfaceSlice([]string{"one", "two", "three", "four"})},
			predicate: functions.NewPredicate(func(i interface{}) bool {
				return len(i.(string)) > 4
			}),
			output: LazyStream{elements: interfaceSlice([]string{"three"})},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.input.LazyFilter(test.predicate), &test.output)
		})
	}
}

func TestLazyStream_LazyMap(t *testing.T) {
	tests := []struct {
		name     string
		input    LazyStream
		function *functions.Function
		output   LazyStream
	}{
		{
			name:  "Map int slice to double each element",
			input: LazyStream{elements: interfaceSlice([]int{0, 1, 2, 3, 4})},
			function: functions.NewFunction(func(i interface{}) interface{} {
				return i.(int) * 2
			}),
			output: LazyStream{elements: interfaceSlice([]int{0, 2, 4, 6, 8})},
		},
		{
			name:  "Append each string in slice with *",
			input: LazyStream{elements: interfaceSlice([]string{"a", "b", "c", "d"})},
			function: functions.NewFunction(func(i interface{}) interface{} {
				return i.(string) + "*"
			}),
			output: LazyStream{elements: interfaceSlice([]string{"a*", "b*", "c*", "d*"})},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.input.LazyMap(test.function), &test.output)
		})
	}
}

