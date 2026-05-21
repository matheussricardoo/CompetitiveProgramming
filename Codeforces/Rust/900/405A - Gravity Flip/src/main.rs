use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = stdin.lock().lines();

    let n: usize = line.next().unwrap().unwrap().trim().parse().unwrap();

    let mut columns: Vec<i64> = line
        .next()
        .unwrap()
        .unwrap()
        .split_whitespace()
        .map(|n| n.parse().unwrap())
        .take(n)
        .collect();

    columns.sort_unstable();

    let result: Vec<String> = columns.iter().map(|x| x.to_string()).collect();
    println!("{}", result.join(" "))
}
