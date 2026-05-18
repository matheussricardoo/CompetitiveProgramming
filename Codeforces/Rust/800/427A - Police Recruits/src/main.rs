use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let _n: i64 = input.trim().parse().unwrap();

    let mut input_events = String::new();

    let _ = io::stdin().read_line(&mut input_events);

    let events: Vec<i64> = input_events
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let mut count = 0;
    let mut polices = 0;

    for e in events {
        if e > 0 {
            polices += e
        } else {
            if polices > 0 {
                polices -= 1;
            } else {
                count += 1;
            }
        }
    }
    println!("{}", count)
}
