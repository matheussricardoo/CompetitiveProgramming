use std::collections::HashSet;
use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input).unwrap();

    let set: HashSet<char> = input.chars().filter(|c| c.is_ascii_lowercase()).collect();

    println!("{}", set.len());
}
