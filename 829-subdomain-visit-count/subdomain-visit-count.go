func subdomainVisits(cpdomains []string) []string {
	counts := make(map[string]int)

	for _, cpdomain := range cpdomains {
		parts := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		for {
			counts[domain] += count
			
			dotIdx := strings.Index(domain, ".")
			if dotIdx == -1 {
				break
			}
			domain = domain[dotIdx+1:]
		}
	}

	result := make([]string, 0, len(counts))
	for domain, count := range counts {
		result = append(result, strconv.Itoa(count)+" "+domain)
	}

	return result
}