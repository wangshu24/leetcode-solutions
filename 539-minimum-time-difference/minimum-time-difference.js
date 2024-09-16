/**
 * @param {string[]} timePoints
 * @return {number}
 */
var findMinDifference = function(timePoints) {
    let minGap = 0;
    let freqMap = {}
    for(let i of timePoints){
        if(freqMap[i] === 1) {return minGap;}
        freqMap[i] = (freqMap[i] || 0) + 1;
    }

    let remains = []
    for(let j in freqMap){
        if(freqMap[j] === 1){remains.push(convertToMinnute(j))}
    }
    remains.sort((a,b)=> a-b);
    console.log(remains)
    const maxGap = 1440;
    for(let i = 0; i < remains.length; i++){
        if(i===0){minGap = remains[i+1]-remains[i]}
        if(remains[i+1]-remains[i] < minGap) {minGap = remains[i+1]-remains[i]}
        if(i===remains.length-1) {
            if(maxGap-remains[i]+remains[0] < minGap) {
                minGap = maxGap-remains[i]+remains[0]
                }              
        }
    }
    return minGap;
};

var convertToMinnute = function(string){
    let parsed = string === "00:00" ? 1440 : parseInt(string.substring(0,2))*60 + parseInt(string.substring(3,5))
     return parsed;
}