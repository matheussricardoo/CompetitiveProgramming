use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let mut points: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    points.sort();

    let distance = points[2] - points[0];
    println!("{}", distance)
}
