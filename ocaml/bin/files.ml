open Printf

module P = Progress
module F = File

type t = {
  mutable files : F.t list;
  self : F.t;
  mutable longest : int;
  mutable progress : P.t;
}

let init path =
  let path = if String.length path > 0 && path.[String.length path - 1] <> '/' then path ^ "/" else path in
  if not (Sys.is_directory path) then
    failwith "Files can't be listed"
  else
    let self = F.create path in
    let progress = P.init 0 in
    { files = []; self; longest = 0; progress }

let get_sizes t =
  let contents = Sys.readdir t.self.F.path in
  let len = Array.length contents in
  t.progress <- P.init len;
  P.display t.progress;
  Array.iter (fun name ->
      let f = F.create (t.self.F.path ^ name) in
      t.longest <- max t.longest (String.length name);
      t.files <- f :: t.files;
      P.update t.progress 1;
      P.display t.progress
    ) contents;
  t.files <- List.rev t.files;
  t.files <- List.sort (fun a b -> compare b.F.size a.F.size) t.files

let sort t =
  t.files <- List.sort (fun a b -> compare b.F.size a.F.size) t.files

let list t order =
  Printf.printf "\027[A\027[K";
  let sep = String.make (t.longest + 19) '=' in
  Printf.printf "%s\n%s\n" t.self.F.path sep;
  let files = if order then t.files else List.rev t.files in
  List.iter
    (fun v ->
       let file_type = if v.F.is_file then "[File]  " else "[Folder]" in
       printf "%s  %-*s  (%s)\n" file_type t.longest v.F.name v.F.real_size)
    files;
  Printf.printf "%s\nTotal: %s\n" sep t.self.F.real_size

let list_raw t order =
  let files = if order then t.files else List.rev t.files in
  List.iter (fun f -> Yojson.Safe.to_string (F.to_yojson f) |> print_endline) files
