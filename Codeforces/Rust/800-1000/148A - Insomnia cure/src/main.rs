use std::io;

fn main() {
    let mut values = Vec::new();

    for _ in 0..5 {
        let mut line = String::new();
        let _ = io::stdin().read_line(&mut line).unwrap();
        if let Ok(val) = line.trim().parse::<i32>() {
            values.push(val);
        }
    }

    let (k, l, m, n, d) = (values[0], values[1], values[2], values[3], values[4]);

    let damaged_count = (1..=d)
        .filter(|i| i % k == 0 || i % l == 0 || i % m == 0 || i % n == 0)
        .count();

    println!("{}", damaged_count)
}
