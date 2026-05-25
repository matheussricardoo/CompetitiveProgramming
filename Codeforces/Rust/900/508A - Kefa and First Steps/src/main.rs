use std::io::{self, Read};

fn main() {
    let mut buffer = String::new();
    io::stdin().read_to_string(&mut buffer).unwrap();
    let mut words = buffer.split_whitespace();

    let n: usize = words.next().unwrap().parse().unwrap();

    let mut max_len = 0;
    let mut current_len = 0;
    let mut prev = 0;

    for i in 0..n {
        let current: i32 = words.next().unwrap().parse().unwrap();

        if i == 0 || current >= prev {
            current_len += 1;
        } else {
            current_len = 1;
        }

        if current_len > max_len {
            max_len = current_len
        }
        prev = current
    }

    println!("{}", max_len)
}
