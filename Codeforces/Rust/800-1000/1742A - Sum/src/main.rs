use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let t: i64 = input.trim().parse().unwrap();

    for _ in 0..t {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let numbers: Vec<i64> = input
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        let a: i64 = numbers[0];
        let b: i64 = numbers[1];
        let c: i64 = numbers[2];

        if a + b == c || a + c == b || b + c == a {
            println!("YES")
        } else {
            println!("NO")
        }
    }
}
