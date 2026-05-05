use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().expect("It's not number");

    let mut current = 0;
    let mut max_capacity = 0;

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let mut parts = input.split_whitespace();

        let exit: i32 = parts.next().unwrap().parse().expect("It's not number");
        let enter: i32 = parts.next().unwrap().parse().expect("It's not number");

        current = current - exit + enter;

        if current > max_capacity {
            max_capacity = current;
        }
    }

    println!("{}", max_capacity);
}
