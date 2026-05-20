use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();
    let mut line = String::new();
    handle.read_line(&mut line).unwrap();

    let s = line.trim().as_bytes();
    for i in 0..s.len() {
        if s[i] == b'H' || s[i] == b'Q' || s[i] == b'9' {
            println!("YES");
            return
        }
    }
    println!("NO")
}
