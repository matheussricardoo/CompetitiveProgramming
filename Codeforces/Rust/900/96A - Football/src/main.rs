fn main() {
    let mut player = String::new();
    let _ = std::io::stdin().read_line(&mut player);

    let players = player.trim().as_bytes();
    if players.len() < 7 {
        println!("NO");
        return;
    }
    let mut count = 1;
    let mut dangerous = false;

    for i in 1..players.len() {
        if players[i] == players[i - 1] {
            count += 1;
            if count == 7 {
                dangerous = true;
                break;
            }
        } else {
            count = 1;
        }
    }

    if dangerous {
        println!("YES")
    } else {
        println!("NO")
    }
}
