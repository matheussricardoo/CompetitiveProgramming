use std::collections::HashSet;
use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let name = input.trim();

    let names: HashSet<char> = name.chars().collect();

    if names.len() % 2 == 0 {
        println!("CHAT WITH HER!")
    } else {
        println!("IGNORE HIM!")
    }
}
