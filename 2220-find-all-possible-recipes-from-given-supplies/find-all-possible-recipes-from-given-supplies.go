func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    sup := make(map[string]bool)
    rti := make(map[string][]string)
    visit := make(map[string]int)
    res := make([]string, 0)

    for _, supp := range supplies {
        sup[supp] = true
    }

   
    for i, rec := range recipes {
        rti[rec] = ingredients[i]
    }
   
    var canMake func(string) bool
    canMake = func(recipe string) bool {
        if val,there := visit[recipe]; there {
            return val == 1
        }

        if _,there := sup[recipe]; there {
            return true
        }

        if _, there := rti[recipe]; !there {
            return false
        }
        visit[recipe] = 0
        for _,i := range rti[recipe] {
            if !canMake(i) {
                visit[i] = -1
                return false
            }
        }

        visit[recipe] = 1
        res = append(res, recipe)
        return true

    }

    for _, recipe := range recipes {
        canMake(recipe)
    }

    return res
}