use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().expect("It's not number");

    if n % 2 == 0 {
        println!("{}", n / 2)
    } else {
        println!("{}", -(n + 1) / 2)
    }
}
