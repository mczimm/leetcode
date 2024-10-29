package main

import "fmt"

func main() {
	//fmt.Println(findAllRecipes([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}))
	//fmt.Println(findAllRecipes([]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}))
	fmt.Println(findAllRecipes2([]string{"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy"},
		[][]string{{"wbjr"}, {"otr", "fzr", "g"}, {"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"}, {"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"}, {"wi", "xgp", "wbjr"}, {"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"}, {"xgp", "otr", "wbjr"}, {"wbjr", "otr"}, {"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"}, {"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"}, {"bxgnm"}, {"wi", "fzr", "otr", "wbjr"}, {"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"}, {"otr", "fzr", "xgp", "wbjr"}, {"xgp", "wbjr", "q", "vpio", "tokfq", "we"}, {"wbjr", "wi", "xgp", "we"}, {"wbjr"}, {"wi"}},
		[]string{"wi", "otr", "wbjr", "fzr", "xgp"}))
}

//["xevvq","izcad","bxgnm","i","hjvu","tokfq","z","g","b","hthy"]

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var res []string
	canMake := true
	supp := make(map[string]struct{})

	for i := 0; i < len(supplies); i++ {
		supp[supplies[i]] = struct{}{}
	}

	for i := 0; i < len(ingredients); i++ {
		canMake = true
		for j := 0; j < len(ingredients[i]); j++ {
			if _, ok := supp[ingredients[i][j]]; !ok {
				canMake = false
				break
			}
		}
		if canMake {
			supp[recipes[i]] = struct{}{}
			res = append(res, recipes[i])
		}
	}

	//if canMake {
	//	res = append(res, recipes[len(recipes)-1])
	//}

	return res
}

func findAllRecipes2(recipes []string, ingredients [][]string, supplies []string) []string {
	inDegree := make(map[string]int)
	g := make(map[string][]string)
	var result []string

	for i, recipie := range recipes {
		inDegree[recipie] = len(ingredients[i])
		for _, component := range ingredients[i] {
			g[component] = append(g[component], recipie)
		}
	}

	queue := supplies
	for len(queue) > 0 {
		ingredient := queue[0]
		queue = queue[1:]

		for _, recipe := range g[ingredient] {
			inDegree[recipe]--
			if inDegree[recipe] == 0 {
				queue = append(queue, recipe)
				result = append(result, recipe)
			}
		}
	}
	return result
}
