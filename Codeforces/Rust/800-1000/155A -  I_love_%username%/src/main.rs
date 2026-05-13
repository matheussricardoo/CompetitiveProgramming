use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: usize = input.trim().parse().unwrap();

    let mut scores_input = String::new();
    let _ = io::stdin().read_line(&mut scores_input);

    let score: Vec<i64> = scores_input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let mut max = score[0];
    let mut min = score[0];
    let mut count = 0;

    for i in 1..n {
        if score[i] > max {
            max = score[i];
            count += 1
        } else if score[i] < min {
            min = score[i];
            count += 1
        }
    }
    println!("{}", count)
}
