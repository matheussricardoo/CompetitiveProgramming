use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let t: i64 = input.trim().parse().unwrap();

    for _ in 0..t {
        let mut s = String::new();
        let _ = io::stdin().read_line(&mut s);

        let s_trim = s.trim();

        if s_trim.to_lowercase() == "yes" {
            println!("YES")
        } else {
            println!("NO")
        }
    }
}
