use std::io;

fn main() {
    let mut t = String::new();
    let _ = io::stdin().read_line(&mut t);

    let mut s = String::new();
    let _ = io::stdin().read_line(&mut s);

    let t_trim = t.trim();
    let s_trim = s.trim();

    let reverse: String = t_trim.chars().rev().collect();

    if reverse == s_trim {
        println!("YES")
    } else {
        println!("NO")
    }
}
