use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let params: Vec<usize> = input
        .split_whitespace()
        .map(|s| s.parse().expect("Error for parser n or k"))
        .collect();

    let _n = params[0];
    let k = params[1];

    input.clear();
    let _ = io::stdin().read_line(&mut input);

    let a: Vec<i32> = input
        .split_whitespace()
        .map(|s| s.parse().expect("The pontuation has to a number"))
        .collect();

    let target_score = a[k - 1];
    let mut count = 0;

    for score in a {
        if score >= target_score && score > 0 {
            count += 1;
        }
    }

    println!("{}", count)
}
