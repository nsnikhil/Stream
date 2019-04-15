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
)

type BufferStream struct {
	elements   []interface{}
	operations []interface{}
}

func OfBuffer(e interface{}) *BufferStream {
	return &BufferStream{elements: interfaceSlice(e)}
}

func (bfs *BufferStream) BufferFilter(predicate *functions.Predicate) *BufferStream {
	bfs.operations = append(bfs.operations, predicate)
	return bfs
}

func (bfs *BufferStream) BufferMap(function *functions.Function) *BufferStream {
	bfs.operations = append(bfs.operations, function)
	return bfs
}

func (bfs *BufferStream) BufferForEach(consumer *functions.Consumer) {
	bfs.operations = append(bfs.operations, consumer)
	bfs.runTermination()
}

func (bfs *BufferStream) runTermination() {
	for _, o := range bfs.operations {
		switch o.(type) {
		case *functions.Predicate:
			bfs.runPredicate(o.(*functions.Predicate))
		case *functions.Function:
			bfs.runFunction(o.(*functions.Function))
		case *functions.Consumer:
			bfs.runConsumer(o.(*functions.Consumer))
		}
	}
}

func (bfs *BufferStream) runPredicate(predicate *functions.Predicate) {
	var el []interface{}
	for _, e := range bfs.elements {
		if predicate.Test(e) {
			el = append(el, e)
		}
	}
	bfs.elements = el
}

func (bfs *BufferStream) runFunction(function *functions.Function) {
	var el []interface{}
	for _, e := range bfs.elements {
		el = append(el, function.Apply(e))
	}
	bfs.elements = el
}

func (bfs *BufferStream) runConsumer(consumer *functions.Consumer) {
	for _, e := range bfs.elements {
		consumer.Accept(e)
	}
}
