package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	client   = &http.Client{Timeout: 15 * time.Second}
	username = "johncmanuel"
	ghUrl    = "https://api.github.com"
)

func main() {
	templateFile := "README.template.md"
	content, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	stars, repos, langs, err := getGitHubStats()
	topLangs := getTopLangs(langs)

	fmt.Println("top langs:", topLangs)

	data := ReadMeData{
		PublicReposCount: strconv.Itoa(repos),
		StarGazersCount:  strconv.Itoa(stars),
		Languages:        topLangs,
	}

	funcMap := template.FuncMap{
		"notLastElement": notLastElement,
	}

	// Replace the values in the template
	tmpl, err := template.New("markdown").Funcs(funcMap).Parse(string(content))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Save the generated markdown to a new file
	outputFile := "README.md"

	err = os.WriteFile(outputFile, output.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Markdown file generated successfully:", outputFile)
}

type GitHubRepository struct {
	StargazersCount int    `json:"stargazers_count"`
	LanguagesURL    string `json:"languages_url"`
}

type Languages map[string]int

type LanguagePercentage struct {
	Language   string
	Percentage float64
}

// All values must be string in order to insert
// them into the README
type ReadMeData struct {
	PublicReposCount string
	StarGazersCount  string
	Languages        []LanguagePercentage
}

func httpAuthReq(method string, url string, token string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

// Return number of stars, number of public repos, and language
// percentage across repos
func getGitHubStats() (int, int, Languages, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		panic("github token not set")
	}

	repoUrl := fmt.Sprintf("%s/users/%s/repos", ghUrl, username)
	repoReq := httpAuthReq("GET", repoUrl, token)

	var repos []GitHubRepository

	response, err := client.Do(repoReq)
	if err != nil {
		panic("Error getting a response.")
	}
	defer response.Body.Close()

	// Contains a language and the number of bytes of
	// code for that language (across all public repos)
	// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repository-languages
	languages := make(Languages)

	if err := json.NewDecoder(response.Body).Decode(&repos); err != nil {
		panic("Error decoding a response.")
	}

	numOfStargazers := 0

	// Loop through each repo and retrieve its popular languages
	for _, repo := range repos {
		numOfStargazers += repo.StargazersCount

		langReq := httpAuthReq("GET", repo.LanguagesURL, token)
		res, err := client.Do(langReq)
		if err != nil {
			fmt.Println("Error fetching the following language_url:", repo.LanguagesURL)
			continue
		}
		defer res.Body.Close()

		var data Languages

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading the response body:", err)
			continue
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}

		for language, bytes := range data {
			_, ok := languages[language]
			if ok {
				languages[language] += bytes
			} else {
				languages[language] = bytes
			}
		}
	}

	return numOfStargazers, len(repos), languages, err
}

func getTopLangs(languages Languages) []LanguagePercentage {
	totalSize := 0
	for _, size := range languages {
		totalSize += size
	}

	var percentages []LanguagePercentage

	for lang, size := range languages {
		percent := math.Round((float64(size) / float64(totalSize)) * 100)
		percentages = append(percentages, LanguagePercentage{lang, percent})
	}

	// Sort the slice in descending order of percentages
	sort.Slice(percentages, func(i, j int) bool {
		return percentages[i].Percentage > percentages[j].Percentage
	})

	// Extract the top 3 languages
	var topLangs []LanguagePercentage
	var otherPercent float64

	for i, lp := range percentages {
		if i < 3 {
			topLangs = append(topLangs, lp)
			continue
		}
		otherPercent += lp.Percentage
	}

	topLangs = append(topLangs, LanguagePercentage{"Others", otherPercent})

	return topLangs
}

func notLastElement(index int, slice []LanguagePercentage) bool {
	return index < len(slice)-1
}
