package gb

// Project represents a gb project. A gb project has a simlar layout to
// a $GOPATH workspace. Each gb project has a standard directory layout
// starting at the project root, which we'll refer too as $PROJECT.
//
//     $PROJECT/                       - the project root
//     $PROJECT/.gogo/                 - used internally by gogo and identifies
//                                       the root of the project.
//     $PROJECT/src/                   - base directory for the source of packages
//     $PROJECT/bin/                   - base directory for the compiled binaries
type Project struct {
	rootdir string
}

// NewContext returns a new build context from this project.
func (p *Project) NewContext() *Context {
	return &Context{
		Project: p,
	}
}