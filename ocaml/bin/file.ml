open Unix

let sizes = [|'B'; 'K'; 'M'; 'G'; 'T'|]

let parse_size line =
  let parts = String.split_on_char '\t' line in
  let size_part = match parts with
    | h :: _ -> String.trim h
    | [] -> "0B"
  in
  let len = String.length size_part in
  if len = 0 then ("0B", 0.)
  else
    let unit = size_part.[len - 1] in
    let num_str =
      String.sub size_part 0 (len - 1)
      |> String.map (fun c -> if c = ',' then '.' else c)
    in
    let v =
      try float_of_string num_str with _ -> 0.
    in
    let rec aux i =
      if i >= Array.length sizes then v
      else if sizes.(i) = unit then v *. (1024. ** float_of_int (i + 1))
      else aux (i + 1)
    in
    (size_part, aux 0)

let run_du path =
  let cmd =
    if path = "/" then "sudo du -sh /" else Printf.sprintf "du -sh %s" (Filename.quote path)
  in
  let ic = Unix.open_process_in cmd in
  let line = input_line ic in
  ignore (Unix.close_process_in ic);
  line

let chop_trailing_slash s =
  if String.length s > 1 && s.[String.length s - 1] = '/' then
    String.sub s 0 (String.length s - 1)
  else s

let basename path = Filename.basename (chop_trailing_slash path)

type t = {
  path : string;
  name : string;
  size : float;
  real_size : string;
  is_file : bool;
}

let create path =
  let line = run_du path in
  let real_size, size = parse_size line in
  let name = let n = basename path in if n = "" then path else n in
  let is_file = not (Sys.is_directory path) in
  { path; name; size; real_size; is_file }

let to_yojson t =
  `Assoc
    [
      ("path", `String t.path);
      ("name", `String t.name);
      ("size", `Float t.size);
      ("real_size", `String t.real_size);
      ("file", `Bool t.is_file);
    ]
