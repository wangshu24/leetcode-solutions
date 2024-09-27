/**
 * @param {string[]} strs
 * @return {string}
 */

class TrieNode {
    constructor(){
        this.children = {}
        this.count = 0
    }
}

class Trie {
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
            node.count++
        }
    }

    getLongestCommonPrefix(word, num){
        let node = this.root
        let result = ""
        // console.log(node)
        // console.log(node)
        // while(node[Object.keys(node)[0]].count === num){
        //    result+=Object.keys(node)[0]
        //    node = node[Object.keys(node)[0]]
        //    node = node.children
        // }
        for(const char of word){
            if(node.children[char]){
                if( node.children[char].count === num){
                    result+=char
                    node = node.children[char]
                }
            }
        }

        return result
    }
}

var longestCommonPrefix = function(strs) {
    let trie = new Trie()
    for(const word of strs){
        trie.insert(word)
    }

    const prefix = trie.getLongestCommonPrefix(strs.sort()[strs.length-1], strs.length) 

    return prefix
};