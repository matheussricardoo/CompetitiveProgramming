use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let n: usize = input.trim().parse().unwrap();

    let mut line = String::new();
    io::stdin().read_line(&mut line).unwrap();
    let heights: Vec<i32> = line
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();

    let mut max_idx = 0;
    let mut min_idx = 0;

    for i in 0..n {
        if heights[i] > heights[max_idx] {
            max_idx = i;
        }
        if heights[i] <= heights[min_idx] {
            min_idx = i;
        }
    }

    let mut ans = max_idx + (n - 1 - min_idx);
    if max_idx > min_idx {
        ans -= 1;
    }

    println!("{}", ans);
}
