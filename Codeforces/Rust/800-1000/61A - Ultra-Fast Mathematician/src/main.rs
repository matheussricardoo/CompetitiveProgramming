use std::io;

fn main() {
    let mut n1 = String::new();
    let mut n2 = String::new();
    let _ = io::stdin().read_line(&mut n1);
    let _ = io::stdin().read_line(&mut n2);

    let n1 = n1.trim();
    let n2 = n2.trim();

    let result: String = n1
        .chars()
        .zip(n2.chars())
        .map(|(c1, c2)| if c1 != c2 { '1' } else { '0' })
        .collect();

    println!("{}", result)
}
