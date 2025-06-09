# Lsize
Little program to list directories and their current size in disk (via du -sh).
## Here an example in the current project.
![Lsize example](/screenshots/example.png)

## OCaml implementation

An OCaml version of `lsize` is available in the `ocaml` directory. To build it, run:

```bash
cd ocaml
# Requires dune, yojson and cmdliner
dune build
```

Example usage:

```bash
dune exec bin/main.exe -- --path . --order
```
