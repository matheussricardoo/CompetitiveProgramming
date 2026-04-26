use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let mut iter = input.split_whitespace();

    let m: usize = iter.next().unwrap().parse().expect("It's not a number");
    let n: usize = iter.next().unwrap().parse().expect("It's not a number");

    let count = m * n / 2;
    println!("{}", count)
}
