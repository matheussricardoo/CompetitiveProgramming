use std::collections::HashSet;
use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input).unwrap();
    let n: usize = input.trim().parse().unwrap();

    let mut levels = HashSet::new();

    for _ in 0..2 {
        let mut line = String::new();
        let _ = io::stdin().read_line(&mut line).unwrap();
        let mut parts = line.split_whitespace();

        parts.next();

        for num_str in parts {
            let level: i32 = num_str.parse().unwrap();
            levels.insert(level);
        }
    }

    if levels.len() == n {
        println!("I become the guy.")
    } else {
        println!("Oh, my keyboard!")
    }
}
