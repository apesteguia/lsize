let path_arg =
  let open Cmdliner in
  let doc = "Path to list" in
  Arg.(value & opt string "." & info ["path"] ~doc)

let order_arg =
  let open Cmdliner in
  let doc = "Order from bigger to smaller" in
  Arg.(value & flag & info ["order"] ~doc)

let raw_arg =
  let open Cmdliner in
  let doc = "Raw JSON output" in
  Arg.(value & flag & info ["raw"] ~doc)

let run path order raw =
  let files = Files.init path in
  Files.get_sizes files;
  if raw then Files.list_raw files order else Files.list files order

let cmd =
  let open Cmdliner in
  let doc = "OCaml lsize" in
  Term.(const run $ path_arg $ order_arg $ raw_arg),
  Term.info "lsize_ocaml" ~version:"0.1" ~doc

let () =
  let open Cmdliner in
  exit (Cmd.eval cmd)
