use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().unwrap();

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let mut numbers = input.split_whitespace();

        let a: i64 = numbers.next().unwrap().parse().expect("It's not number");
        let b: i64 = numbers.next().unwrap().parse().expect("It's not number");

        if a % b == 0 {
            println!("0");
        } else {
            let reminders = a % b;
            println!("{}", b - reminders)
        }
    }
}
