impl Solution {
    pub fn compare_version(version1: String, version2: String) -> i32 {
        let parts1: Vec<&str> = version1.split('.').collect();
        let parts2: Vec<&str> = version2.split('.').collect();
        
        let max_len = parts1.len().max(parts2.len());
        
        for i in 0..max_len {
            let v1 = if i < parts1.len() {
                parts1[i].parse::<i32>().unwrap_or(0)
            } else {
                0
            };
            
            let v2 = if i < parts2.len() {
                parts2[i].parse::<i32>().unwrap_or(0)
            } else {
                0
            };
            
            match v1.cmp(&v2) {
                std::cmp::Ordering::Less => return -1,
                std::cmp::Ordering::Greater => return 1,
                std::cmp::Ordering::Equal => continue,
            }
        }
        
        0
    }
}

/*
 * TC: O(n + m) for parsing and comparison
 * SC: O(n + m) for storing split parts
 */