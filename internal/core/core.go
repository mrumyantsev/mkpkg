package core

const (
	AppName = "mkpkg"
	Version = "1.0.0"

	Help = `Usage: %s [OPTION]... PKG_PATH
Creates the named types like structs or interfaces within Go packages.

Mandatory arguments to long options are mandatory for short options too.
  -c, --ctor                   name constructor for struct ('New()' is default)
      --eol                    override EOLs for new code ('\n' is default)
  -f, --filename               specify file name
  -h, --help                   display this help and exit
      --iface                  switch struct creating to interface creating
  -m, --methods                apply list of methods (semicolon separated)
  -n, --name                   specify name for the new object
  -v, --version                output version information and exit
`

	Hint = "Try '%s --help' for more information.\n"
)
