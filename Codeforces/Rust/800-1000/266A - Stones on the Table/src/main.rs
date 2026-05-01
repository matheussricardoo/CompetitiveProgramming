use std::io;

fn main() {
    let mut n_str = String::new();
    let _ = io::stdin().read_line(&mut n_str);
    let n: usize = n_str.trim().parse().expect("It's not a number");

    let mut s = String::new();
    let _ = io::stdin().read_line(&mut s);
    let bytes = s.trim().as_bytes();

    let mut remove = 0;

    for i in 0..n - 1 {
        if bytes[i] == bytes[i + 1] {
            remove += 1;
        }
    }

    println!("{}", remove)
}
