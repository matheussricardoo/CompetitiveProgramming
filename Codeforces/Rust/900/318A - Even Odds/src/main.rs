use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();
    let mut line = String::new();
    handle.read_line(&mut line).unwrap();

    let numbers: Vec<i64> = line
        .split_whitespace()
        .map(|w| w.parse().unwrap())
        .collect();

    let n: i64 = numbers[0];
    line.clear();
    let mut k: i64 = numbers[1];

    let odd = if n % 2 == 0 { n / 2 } else { (n + 1) / 2 };

    if k <= odd {
        println!("{}", 2 * k - 1)
    } else {
        k = k - odd;
        println!("{}", 2 * k)
    }
}
