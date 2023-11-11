# johncmanuel

```python
>>> from goated_programmers import johncmanuel
>>> import json
>>> john = johncmanuel()
>>> print(json.dumps(john.bio), indent=2)
{
  "name": "John Carlo Manuel",
  "occupation": "Software Engineer",
  "pronouns": "he/him",
  "schools": [
    "Skyline College",
    "California State University, Fullerton"
  ],
  "interests": [
    "Full-Stack Development",
    "Game Development",
    "Distributed Systems"
  ],
  "github_stats": {
    "public_repos": "{{.PublicReposCount}}",
    "stargazers": "{{.StarGazersCount}}"
  },
  "languages": {
    {{- range $i, $lang := .Languages}}
    "{{.Language | html}}": {
      "usagePercent": {{.Percentage}}
    }{{if notLastElement $i $.Languages}},{{end}}{{end}}
  },
  "hobbies": [
    "Video Games",
    "Anime",
    "Code",
    "Writing",
    "Manga",
  ]
}
```

## Socials

<p>
<a href="https://www.linkedin.com/in/johncarlomanuel/" target="_blank">
    <img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />
  </a>
  <a href="mailto:johncarlomanuel@csu.fullerton.edu">
    <img alt="Gmail" src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white" />
  </a>
 <a href="https://www.johncarlomanuel.com/" target="_blank">
    <img alt="Website" src="https://img.shields.io/badge/website-000000?style=for-the-badge&logo=About.me&logoColor=white" />
  </a>
<a href=>
</p>
