/**
 * @param {string[]} words
 * @return {number[]}
 */

class TrieNode{
    constructor() {
        this.children = {}
        this.prefixCount = 0
    }
}

class Trie{
    constructor(){
        this.root = new TrieNode()
    }

    insert(word){
        let node = this.root
        for(const char of word){
            if(!node.children[char]){
                node.children[char] = new TrieNode()
            }
            node = node.children[char]
            node.prefixCount++
        }
    }

    getPrefixCount(word){
        let node = this.root
        let count = 0
        for(const char of word){
            node = node.children[char]
            count += node.prefixCount
        }
        return count
    }
}

var sumPrefixScores = function(words) {
    //Time limit exceeded
    // let result = []
    // let prefixes = {}
    // for(let x = 0; x < words.length; x++){
    //     let prefixSum = 0
    //     for(let y = 1; y <= words[x].length; y++){
    //         prefix = words[x].slice(0,y)
    //         if(!prefixes[prefix]){
    //             prefixSum += words.filter((word) => word.slice(0,y) === prefix).length
    //             prefixes[prefix] = prefixSum
    //         } else {
    //             prefixSum= prefixes[prefix]
    //         }
    //     }
    //     result.push(prefixSum)
    // }
    // return result

    let trie = new Trie()
    for(const word of words){
        trie.insert(word)
        console.log("at word: ", word, " trie is : " ,trie)
    }
    //console.log(trie.root.children.a.children.b)
    const result = words.map(word => trie.getPrefixCount(word))
    return result
};