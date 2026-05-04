use std::io;

fn main() {
    let mut n = String::new();
    let _ = io::stdin().read_line(&mut n);

    let mut s = String::new();
    let _ = io::stdin().read_line(&mut s);

    let mut count_a = 0;
    let mut count_d = 0;

    for c in s.trim().chars() {
        if c == 'A' {
            count_a += 1;
        } else {
            count_d += 1;
        }
    }

    if count_a > count_d {
        println!("Anton");
    } else if count_d > count_a {
        println!("Danik");
    } else {
        println!("Friendship");
    }
}
