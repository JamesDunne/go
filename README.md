# My Changes

This branch adds a ternary conditional expression to Go, equivalent to the ternary `a ? b : c` expression found in the C programming language.

The syntax for Go is:

`a then b else c`

I originally wanted the expression to be of the form `if a then b else c` in a naturally readable way but adding the `if` prefix caused many shift/reduce conflicts in the `go.y` yacc grammar source causing conflicts with the `if` statement. I did not want to spend time trying to resolve these conflicts so I dropped the `if` requirement and that solved all the conflicts.

Unfortunately, `then` is now a reserved keyword. Ideally I want it to be a non-reserved keyword. There was one conflict in the Go standard library in `runtime/mgc.go` which was using `then` as a variable name. Renaming the variable solved the conflict of course, but it would be better form to not reserve the word `then`, thus allowing it to be used as a variable name, provided that doesn't introduce any grammar conflicts. I think Go does not have any non-reserved keywords in the grammar so adding that ability to the lexer is probably more complex in terms of effort than I'm willing to invest.

# Examples

Trivial functions like min and max can now become one-liners without requiring a full if/else statement and two body blocks:

`func max(a, b int) int { return a > b then a else b }`

More complex expressions can be used of course, so long as the `then` expression type exactly matches the `else` expression type.

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](doc/gopher/fiveyears.jpg)

For documentation about how to install and use Go,
visit https://golang.org/ or load doc/install-source.html
in your web browser.

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Please report issues here: https://golang.org/issue/new

Go is the work of hundreds of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

##### Please note that we do not use pull requests.

Unless otherwise noted, the Go source files are distributed
under the BSD-style license found in the LICENSE file.

--

## Binary Distribution Notes

If you have just untarred a binary Go distribution, you need to set
the environment variable $GOROOT to the full path of the go
directory (the one containing this file).  You can omit the
variable if you unpack it into /usr/local/go, or if you rebuild
from sources by running all.bash (see doc/install-source.html).
You should also add the Go binary directory $GOROOT/bin
to your shell's path.

For example, if you extracted the tar file into $HOME/go, you might
put the following in your .profile:

	export GOROOT=$HOME/go
	export PATH=$PATH:$GOROOT/bin

See https://golang.org/doc/install or doc/install.html for more details.
