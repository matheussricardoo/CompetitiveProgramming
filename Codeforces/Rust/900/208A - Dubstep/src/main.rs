use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    stdin.lock().read_line(&mut line).unwrap();
    let input = line.trim();

    let words: Vec<&str> = input.split("WUB").filter(|s| !s.is_empty()).collect();

    println!("{}", words.join(" "))
}
