use std::io::{self, BufRead, Write, BufWriter};

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();
    let out = io::stdout();
    let mut writer = BufWriter::new(out.lock());

    let mut line = String::new();

    handle.read_line(&mut line).unwrap();
    let t: usize = line.trim().parse().unwrap();

    for _ in 0..t {
        line.clear(); 
        handle.read_line(&mut line).unwrap();
        
        let n: i64 = line.trim().parse().unwrap();

        if n & (n - 1) != 0 {
            writeln!(writer, "YES").unwrap();
        } else {
            writeln!(writer, "NO").unwrap();
        }
    }
}