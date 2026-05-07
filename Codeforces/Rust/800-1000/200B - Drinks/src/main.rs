use std::io;

fn main() {
    let mut input = String::new();
    let _  = io::stdin().read_line(&mut input);

    let n:f64 = input.trim().parse().expect("It's not number");

    let mut input_p = String::new();
    let _  = io::stdin().read_line(&mut input_p);

    let sum:f64 = input_p
        .split_whitespace()
        .map(|j| j.parse::<f64>().expect("It's not number"))
        .sum();

    println!("{}",sum / n)
}
