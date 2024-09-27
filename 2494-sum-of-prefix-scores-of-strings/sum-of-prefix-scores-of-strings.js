/**
 * @param {string[]} words
 * @return {number[]}
 */


class TrieNode {
    constructor(){
        this.children = {}
        this.prefixCount = 0
    }
}

class Trie {
    constructor(){
        //when trie is instantiated, always start first TrieNode for all words
        this.root = new TrieNode()
    }

    insert(word){
        //start from the same trie node, to compare if new word follows the same path or new path (order) is created
        let node = this.root
        for(const char of word){
            if(!node.children[char]){
                //check for existence of char, in order
                node.children[char] = new TrieNode()
            }
            //reset root to the latest char being inspected to maintain order of char, hence suitable to count and track prefix
            node = node.children[char]
            node.prefixCount++
        }
    }

    getPrefixCount(word){
        //after all insert has been done, all word has been filtered and check for order and repetition, hence, if same prefix has occured, will be counted, hence we just have to tally up
        let node = this.root
        let count = 0
        for(const char of word){
            //this is following the exact order of prefix, hence any prefix that has occurred to this exact char being inspected has been recorded, now we tally up
            node = node.children[char]
            count += node.prefixCount
        }
        return count
    }
}

var sumPrefixScores = function(words) {
    trie = new Trie()
    for(const word of words){
        trie.insert(word)
    }

    const res = words.map(word => trie.getPrefixCount(word))

    return res;
};