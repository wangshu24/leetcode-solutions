/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const open = ["(", "{", "["]
    const close = [")", "}", "]"]
    let opening = []
    let valid = true
    for(let i = 0; i < s.length; i++){
        if(open.includes(s[i])) {
            opening.push(s[i])
        } else {
            let lastChar = opening.pop()
            let indexOpen = open.indexOf(lastChar)
            let indexClose = close.indexOf(s[i])
            if(indexOpen === indexClose) {
                valid = true
            } else {valid = false}
        }
        if(!valid) { break; }
    }
    if(opening.length > 0) { valid = false}
    return valid;
};