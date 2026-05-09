use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().unwrap();
    let mut sum = 0;

    for _ in 0..n {
        let mut s = String::new();
        let _ = io::stdin().read_line(&mut s);

        match s.trim() {
            "Tetrahedron" => sum += 4,
            "Cube" => sum += 6,
            "Octahedron" => sum += 8,
            "Dodecahedron" => sum += 12,
            "Icosahedron" => sum += 20,

            _ => (),
        }
    }
    println!("{}", sum)
}
