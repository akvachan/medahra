# medahra - Mediocre Dataset for HR Applications

## Motivation and Goal

I need to evaluate BM25-HR, my custom BM25 algorithm fine-tuned for HR
information retrieval / recommender system.

Goal is to create an information retrieval dataset, where each candidate is
mapped to top 10 most suitable positions, thus providing a gold standard for
Hit@10, Hit@1 or MRR evaluations.

You can think of a `candidate` as a "query" and `position` as a "document" in
an information retrieval system.

The dataset needs to be in 3 files:
- `candidates.jsonl`
- `positions.jsonl`
- `matches.jsonl`

## Structure of `candidates.jsonl`

```json
{
 "candidate_id": "serial id of a candidate",
 "about": "about a candidate",
 "short_about": "short, AI-generated description of a candidate",
 "general_skills": [
   "a",
   "list",
   "of",
   "general",
   "skills"
 ],
 "preferred_positions": [
   "a",
   "list",
   "of",
   "preferred",
   "positions"
 ],
 "preferred_skills": [
   "a",
   "list",
   "of",
   "preferred",
   "skills"
 ],
 "preferred_min_salary": "preferred minimal salary, as integer",
 "preferred_modes": [
   "office",
   "remote",
   "hybrid"
 ],
 "preferred_contract_types": [
   "full_time",
   "part_time",
   "contract",
   "freelance",
   "internship"
 ],
 "working_experience": [
   {
     "date_from": "date in DDMMYYYY format, for example 21052001",
     "date_to": "date in DDMMYYYY format, for example 21052001",
     "days_duration": "duration in days",
     "description": "description of the experience",
     "seniority": "intern / junior / medior / senior / manager / executive",
     "skills": [
       "a",
       "list",
       "of",
       "relevant",
       "skills"
     ],
     "title": "the latest title",
     "type": "full_time / part_time / contract / freelance / internship"
   },
   {
     "date_from": "date in DDMMYYYY format, for example 21052001",
     "date_to": "date in DDMMYYYY format, for example 21052001",
     "description": "description of the experience",
     "hours_duration": "duration in hours",
     "seniority": "intern / junior / medior / senior / manager / executive",
     "skills": [
       "a",
       "list",
       "of",
       "relevant",
       "skills"
     ],
     "title": "the second latest title",
     "type": "full_time / part_time / contract / freelance / internship"
   }
 ],
 "qualifications": [
   {
     "date_from": "date in DDMMYYYY format",
     "date_to": "date in DDMMYYYY format",
     "description": "description of the education",
     "hours_duration": "duration in hours",
     "institution": "instituion that provided the education",
     "name": "program / education name",
     "nqf": "qualification gained in NQF framework, ranges from 1 to 10",
     "qualification_custom": "qualification gained using custom label",
     "skills": [
       "a",
       "list",
       "of",
       "skills"
     ]
   }
 ],
 "languages": [
   {
     "cefr": "language level according to CEFR (A2, B1, C2)",
     "language": "language"
   }
 ]
}
```

## Structure of `positions.jsonl`

```json
{
 "position_id": "a serial position identifier",
 "description": "a full job description",
 "short_description": "an AI-generated, short description",
 "modes": [
   "remote",
   "office",
   "hybrid"
 ],
 "contract_types": [
   "full_time",
   "part_time",
   "contract",
   "freelance",
   "internship"
 ],
 "required_skills": [
   "a",
   "list",
   "of",
   "required",
   "skills"
 ],
 "required_qualifications": [
   {
     "name": "program / education name",
     "description": "description of the education",
     "nqf": "qualification gained in NQF framework, ranges from 1 to 10",
     "qualification_custom": "qualification gained using custom label",
     "skills": [
       "a",
       "list",
       "of",
       "skills"
     ]
   }
 ],
 "required_languages": [
   {
     "cefr": "language level according to CEFR (A2, B1, C2)",
     "language": "language"
   }
 ],
 "min_salary": "an integer for minimum salary in range",
 "max_salary": "an integer for maximum salary in range",
 "exact_salary": "an integer for exact salary"
}
```

## Structure of `matches.jsonl`

```json
{
 "match_id": "a serial id of a match",
 "candidate_id": "a serial id of a candidate",
 "positions": [
   {
     "position_id": "id of the 1st position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 2nd position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 3rd position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 4th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 5th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 6th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 7th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 8th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 9th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   },
   {
     "position_id": "id of the 10th position in the ranking",
     "score": "float",
     "reasoning": "reasoning"
   }
 ],
 "methods": [
   {
     "type": "type of a method (reranker / manual / llm)",
     "model": "exact model that was utilized (if applicable)"
   }
 ],
 "created_at": "unix timestamp of the creation"
}
```

## Precursor datasets

- For positions: [https://www.kaggle.com/datasets/arshkon/linkedin-job-postings](https://www.kaggle.com/datasets/arshkon/linkedin-job-postings)
- For candidates: [https://www.kaggle.com/datasets/snehaanbhawal/resume-dataset](https://www.kaggle.com/datasets/snehaanbhawal/resume-dataset)

## Scale

- Positions (Job Postings): 123848 rows
- Candidates (Resumes): 2483 rows
- Matches (Candidate-Positions tuples): 2483 rows * top 10 results = 24830 items

## Methodology

1. Parse as much information as possible from precursors, keep it as raw as possible.
2. Prompt a small language model to fill in missing information.
3. Use three different reranker models from three different cross-encoder families as annotators for the matches.
4. Use majority vote as the source of truth (2 rerankers should have the same position, else choose position from the best reranker).
5. Cross-check matched positions with multiple fields from the original datasets.
