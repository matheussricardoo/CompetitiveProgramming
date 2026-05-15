use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let values: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let k: i64 = values[0];
    let r: i64 = values[1];

    for shovels in 1..=10 {
        let total_price = shovels * k;
        if total_price % 10 == 0 || total_price % 10 == r {
            println!("{}", shovels);
            break;
        }
    }
}
