use std::io;

fn main() {
    let mut s1 = String::new();
    let mut s2 = String::new();
    let mut pile = String::new();

    let _ = io::stdin().read_line(&mut s1).unwrap();
    let _ = io::stdin().read_line(&mut s2).unwrap();
    let _ = io::stdin().read_line(&mut pile).unwrap();

    let mut names: Vec<char> = (s1.trim().to_string() + s2.trim()).chars().collect();
    let mut pile_chars: Vec<char> = pile.trim().chars().collect();

    names.sort();
    pile_chars.sort();

    if names == pile_chars {
        println!("YES")
    } else {
        println!("NO")
    }
}
