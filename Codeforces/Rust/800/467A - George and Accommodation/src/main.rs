use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().expect("It's not number");

    let mut count = 0;

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let mut inputs = input.split_whitespace();

        let p: i64 = inputs.next().unwrap().parse().expect("It's not number");
        let q: i64 = inputs.next().unwrap().parse().expect("It's not number");

        let size_room = p - q;

        if size_room.abs() >= 2 {
            count += 1
        }
    }
    println!("{}", count)
}
