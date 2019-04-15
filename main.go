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

package main

import (
	"fmt"
	"github.com/nsnikhil/Stream/functions"
	"github.com/nsnikhil/Stream/stream"
)

type Person struct {
	name string
	rank string
}

func main() {

	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var f, m []int
	for _, e := range d {
		if e%2 == 0 {
			f = append(f, e)
		}
	}
	for _, e := range f {
		m = append(m, e*2)
	}
	for _, e := range m {
		fmt.Println(e)
	}

	a := stream.Of([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	a.Filter(functions.NewPredicate(func(i interface{}) bool {
		return i.(int)%2 == 0
	})).Maps(functions.NewFunction(func(i interface{}) interface{} {
		return i.(int) * 2
	})).ForEach(functions.NewConsumer(func(i interface{}) {
		fmt.Println(i)
	}))

	b := stream.OfLazy([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	b.LazyFilter(functions.NewPredicate(func(i interface{}) bool {
		return i.(int)%2 == 0
	})).LazyMap(functions.NewFunction(func(i interface{}) interface{} {
		return i.(int) * 2
	})).LazyForEach(functions.NewConsumer(func(i interface{}) {
		fmt.Println(i)
	}))

	e := stream.Of([]string{"one", "two", "three", "four", "five", "six", "seven"})
	e.Filter(functions.NewPredicate(func(i interface{}) bool {
		return len(i.(string)) >= 4
	})).Maps(functions.NewFunction(func(i interface{}) interface{} {
		return i.(string) + "*"
	})).ForEach(functions.NewConsumer(func(i interface{}) {
		fmt.Println(i)
	}))

	g := stream.OfLazy([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	g.LazyMap(functions.NewFunction(func(i interface{}) interface{} {
		return Person{name: fmt.Sprintf("%s %d", "Name", i.(int)), rank: fmt.Sprintf("%s %d", "Rank ", i.(int))}
	})).LazyForEach(functions.NewConsumer(func(i interface{}) {
		fmt.Println(i)
	}))

	stream.Of([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}).Filter(evenPredicate()).Maps(mapToDouble()).ForEach(printIt())
	stream.OfLazy([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}).LazyFilter(evenPredicate()).LazyMap(mapToDouble()).LazyForEach(printIt())
	stream.OfBuffer([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}).BufferFilter(evenPredicate()).BufferMap(mapToDouble()).BufferForEach(printIt())

	//stream.Generate(0, it + 1).Filter(it % 2 == 0).Map(it * 2).ForEach(println(it))

	//stream.Of(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Filter(it % 2 == 0).Map(it * 2).ForEach(println(it))

}

func evenPredicate() *functions.Predicate {
	return functions.NewPredicate(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
}

func mapToDouble() *functions.Function {
	return functions.NewFunction(func(i interface{}) interface{} {
		return i.(int) * 2
	})
}

func printIt() *functions.Consumer {
	return functions.NewConsumer(func(i interface{}) {
		fmt.Println(i)
	})
}
