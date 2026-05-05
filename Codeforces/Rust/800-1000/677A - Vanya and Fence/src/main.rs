use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let args: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse().expect("It's not number"))
        .collect();

    let _n = args[0];
    let h = args[1];

    input.clear();

    let mut l: i32 = 0;
    let _ = io::stdin().read_line(&mut input);

    let heights = input
        .split_whitespace()
        .map(|s| s.parse::<i64>().expect("It's not number"));

    for height in heights {
        if height <= h { l += 1 } else { l += 2 }
    }

    println!("{}", l)
}
