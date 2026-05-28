use std::io::{self, BufRead, BufWriter, Write};

fn solve(mut n: i32) -> i32 {
    let mut c2 = 0;
    let mut c3 = 0;

    while n % 2 == 0 {
        c2 += 1;
        n /= 2;
    }

    while n % 3 == 0 {
        c3 += 1;
        n /= 3;
    }

    if n != 1 || c2 > c3 {
        return -1;
    }

    (c3 - c2) + c3
}

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
        let n: i32 = line.trim().parse().unwrap();

        writeln!(writer, "{}", solve(n)).unwrap();
    }
}
