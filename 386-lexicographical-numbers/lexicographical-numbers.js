/**
 * @param {number} n
 * @return {number[]}
 */

var lexicalOrder = function(n) {
    // let result = []
    // temp = n
    // let lexInt = 1
    // let index = 0
    // let num = 0
    // while(temp-1>0){
    //     let tempNum = num +1
    //     if(getLexInt(tempNum) === lexInt ){
    //         if(tempNum < n){
    //             result.push(tempNum)
    //             num = tempNum
    //         }
    //     }
    // }
    let result = []
    for(let i = 1; i < n+1; i++){
        result.push(i)
    }
    return result.sort()
};

var getLexInt = function(n){
    let firstInt = parseInt(n.toString()[0])
    return firstInt
}
// var lexicalOrder = function(n) {
//     let result = []
//     switch (true){
//         case 10 <= num <= 99:

//         case 100 <= num <= 999:
//         case 1000 <= num <= 9999:
//         case 10000 <= num <= 50000:
//         default:
//     }
// };

// var sortLex = function(a,b){
//     if()
// }