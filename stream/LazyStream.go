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
	"sync"
)

type LazyStream struct {
	elements []interface{}
	wg       sync.WaitGroup
}

func OfLazy(e interface{}) *LazyStream {
	return &LazyStream{elements: interfaceSlice(e)}
}

func (ls *LazyStream) LazyFilter(predicate *functions.Predicate) *LazyStream {
	ls.wg.Add(1)
	v := <-ls.lazyFilter(predicate)
	return &v
}

func (ls *LazyStream) lazyFilter(predicate *functions.Predicate) chan LazyStream {
	c := make(chan LazyStream)
	t := LazyStream{}
	go func() {
		for _, e := range ls.elements {
			if predicate.Test(e) {
				t.elements = append(t.elements, e)
			}
		}
		c <- t
		ls.wg.Done()
	}()
	return c
}

func (ls *LazyStream) LazyMap(function *functions.Function) *LazyStream {
	ls.wg.Add(1)
	v := <-ls.lazyMap(function)
	return &v
}

func (ls *LazyStream) lazyMap(function *functions.Function) chan LazyStream {
	c := make(chan LazyStream)
	t := LazyStream{}
	go func() {
		for _, e := range ls.elements {
			t.elements = append(t.elements, function.Apply(e))
		}
		c <- t
		ls.wg.Done()
	}()
	return c
}

func (ls *LazyStream) LazyForEach(consumer *functions.Consumer) {
	ls.wg.Wait()
	for _, e := range ls.elements {
		consumer.Accept(e)
	}
}
