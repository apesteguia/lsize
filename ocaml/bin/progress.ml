type t = {
  mutable chars : char array;
  mutable len : int;
  mutable percentage : int;
}

let init len =
  let chars = Array.make len '.' in
  Printf.printf "Reading filesystem\n";
  { chars; len; percentage = 0 }

let update t n =
  t.percentage <- min t.len (t.percentage + n);
  for i = 0 to t.percentage - 1 do
    t.chars.(i) <- '='
  done

let display t =
  Printf.printf "\027[F\027[K(%d/%d)[" t.percentage t.len;
  Array.iter (fun c -> Printf.printf "%c" c) t.chars;
  Printf.printf "]\n%!"
