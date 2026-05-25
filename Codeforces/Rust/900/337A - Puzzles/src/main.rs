use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();

    let mut line1 = String::new();
    handle.read_line(&mut line1).unwrap();
    let mut tokens1 = line1.split_whitespace();

    let n: usize = tokens1.next().unwrap().parse().unwrap();
    let _m: usize = tokens1.next().unwrap().parse().unwrap();

    let mut line2 = String::new();
    handle.read_line(&mut line2).unwrap();

    let mut puzzles: Vec<i32> = line2
        .split_whitespace()
        .map(|w| w.parse().unwrap())
        .collect();

    puzzles.sort_unstable();

    let min_diff = puzzles
        .windows(n)
        .map(|window| window[n - 1] - window[0])
        .min()
        .unwrap();

    println!("{}", min_diff);
}
