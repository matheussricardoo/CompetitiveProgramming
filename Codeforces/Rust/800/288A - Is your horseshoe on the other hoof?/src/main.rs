use std::{collections::HashSet, io};

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let numbers: HashSet<i32> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let result = 4 - numbers.len();

    println!("{}", result)
}
