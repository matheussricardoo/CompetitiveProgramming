use std::io;

fn main() {
    let mut input = String::new();

    let _ = io::stdin().read_line(&mut input);

    let w: u32 = input.trim().parse().expect("Not a number");
    if w > 2 && w %2 == 0 {
        println!("YES")
    } else {
        println!("NO")
    }
}
