use std::io::{self, BufRead};

fn parse_int(bytes: &[u8]) -> (i32, i32) {
    let mut pos = 0;
    let len = bytes.len();

    let get_next = |p: &mut usize| -> i32 {
        while *p < len && bytes[*p] <= b' ' {
            *p += 1;
        }
        let mut res = 0;
        while *p < len && bytes[*p] >= b'0' && bytes[*p] <= b'9' {
            res = res * 10 + (bytes[*p] - b'0') as i32;
            *p += 1;
        }
        res
    };

    let n = get_next(&mut pos);
    let m = get_next(&mut pos);
    (n, m)
}

fn main() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();
    let mut line = String::new();

    handle.read_line(&mut line).unwrap();
    let (n, m) = parse_int(line.as_bytes());

    let min_val = std::cmp::min(n, m);

    if min_val % 2 != 0 {
        println!("Akshat");
    } else {
        println!("Malvika");
    }
}
