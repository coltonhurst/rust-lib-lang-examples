#[no_mangle]
pub extern "C" fn add(left: i32, right: i32) -> i32 {
    adder(left, right)
}

pub fn adder(left: i32, right: i32) -> i32 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = adder(2, 2);
        assert_eq!(result, 4);
    }
}
