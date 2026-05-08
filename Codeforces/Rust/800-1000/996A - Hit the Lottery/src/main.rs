use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input).unwrap();

    let mut n: i64 = input.trim().parse().unwrap();

    let bills = vec![100, 20, 10, 5, 1];
    let mut count = 0;
    for i in bills {
        count += n / i;
        n %= i;
    }
    println!("{}", count)
}
