use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().expect("It's not number");

    let mut count = 0;
    let mut previous = 0;

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let current: i64 = input.trim().parse().expect("It's not number");

        if current != previous {
            count += 1;
        }

        previous = current;
    }
    println!("{}", count)
}
