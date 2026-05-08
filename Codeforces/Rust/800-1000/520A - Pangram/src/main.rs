use std::{collections::HashMap, io};

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let _n: i64 = input.trim().parse().unwrap();

    input.clear();

    let _ = io::stdin().read_line(&mut input);
    let s = input.trim();

    let mut letters: HashMap<char, bool> = HashMap::new();
    for c in s.to_lowercase().chars() {
        if c >= 'a' && c <= 'z' {
            letters.insert(c, true);
        }
    }

    if letters.len() == 26 {
        println!("YES");
    } else {
        println!("NO");
    }
}
