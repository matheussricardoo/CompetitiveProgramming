use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let _n: i64 = input.trim().parse().unwrap();

    println!("{}", 25)
}
