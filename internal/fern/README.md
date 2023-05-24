# internal/fern

The `internal/fern` directory contains packages generated by `fern`.
This includes the IntermediateRepresentation, which was once written
by hand but is now consumed from this package (i.e. compiler bootstrapping),
as well as the `generator-exec` configuration, which represents the
schema for the generator's `config.json` file.

Note that this generator was written with IR v20, and that is what
is found in `ir.json`.