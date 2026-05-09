use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let t: i64 = input.trim().parse().unwrap();

    for _ in 0..t {
        let mut input = String::new();
        let _ = io::stdin().read_line(&mut input);

        let n: i64 = input.trim().parse().unwrap();
        println!("{}", (n - 1) / 2)
    }
}
