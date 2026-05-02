use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let numbers: Vec<i32> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let (k, n, w) = (numbers[0], numbers[1], numbers[2]);

    let solutions = k * (w * (w + 1) / 2);
    let borrowed = solutions - n;

    if borrowed < 0 {
        println!("0")
    } else {
        println!("{}", borrowed)
    }
}
