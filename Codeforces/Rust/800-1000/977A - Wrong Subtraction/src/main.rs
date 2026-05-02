use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let mut numbers = input.split_ascii_whitespace();

    let mut n: i32 = numbers.next().unwrap().parse().expect("It's not number");
    let k: i32 = numbers.next().unwrap().parse().expect("It's not number");

    for _ in 0..k {
        if n % 10 == 0 {
            n /= 10;
        } else {
            n -= 1;
        }
    }
    println!("{}", n)
}
