use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();

    let mut line = String::new();

    handle.read_line(&mut line).unwrap();
    let n: usize = line.trim().parse().unwrap();
    line.clear();

    handle.read_line(&mut line).unwrap();

    let mut coins: Vec<i32> = line
        .split_whitespace()
        .map(|w| w.parse().unwrap())
        .collect();

    coins.truncate(n);

    let total_sum: i32 = coins.iter().sum();

    coins.sort_unstable_by(|a, b| b.cmp(a));

    let mut my_sum = 0;
    let mut count = 0;

    for coin in coins {
        my_sum += coin;
        count += 1;
        if my_sum > total_sum / 2 {
            break;
        }
    }

    println!("{}", count)
}
