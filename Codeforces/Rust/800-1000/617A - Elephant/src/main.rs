use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let inpu = input.trim();
    let n: usize = inpu.parse().expect("It's not a number");
    let mut count = n / 5;
    if n % 5 != 0 {
        count += 1
    }

    println!("{}", count)
}
