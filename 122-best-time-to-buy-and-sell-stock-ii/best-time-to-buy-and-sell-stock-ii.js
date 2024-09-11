/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let profit =0;
    let trnding = false;
    prices.forEach((p, i) => 
        {
        if(prices[i+1] && (prices[i+1]-p)>0){
            profit+=prices[i+1]-p
        }
    }
    )
    return profit;
};