use std::io;

fn main() {
    let mut input1 = String::new();
    let mut input2 = String::new();

    let _ = io::stdin().read_line(&mut input1);
    let _ = io::stdin().read_line(&mut input2);

    let s1 = input1.trim().to_lowercase();
    let s2 = input2.trim().to_lowercase();

    if s1 < s2 {
        println!("-1")
    } else if s1 > s2 {
        println!("1")
    } else {
        println!("0")
    }
}
