use std::{cmp::min, io};

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let numbers: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let n: i64 = numbers[0];
    let k: i64 = numbers[1];
    let l: i64 = numbers[2];
    let c: i64 = numbers[3];
    let d: i64 = numbers[4];
    let p: i64 = numbers[5];
    let nl: i64 = numbers[6];
    let np: i64 = numbers[7];

    let ml_drink = k * l;
    let tostas = ml_drink / nl;
    let lemons = c * d;
    let salt = p / np;
    let take_it_light = min(tostas, min(lemons, salt)) / n;
    println!("{}", take_it_light)
}
