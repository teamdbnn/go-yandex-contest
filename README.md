# go-yandex-contest
yandex contest api client (built from swagger https://api.contest.yandex.net/api/public/swagger-ui.html)



## Methods
| &#9744; | Method | Path  | Description |
|-----|--------|-------|-------------|
| :white_check_mark: | GET    | `/competitions/{competitionId}/participants` | Get registered participants of competition  |
| :white_check_mark: | GET | `/competitions/{competitionId}/participants` | Get registered participants of competition |
| :white_check_mark: | POST | `/competitions/{competitionId}/participants` | Register participants into competition |
| :white_check_mark: | GET | `/contests/{contestId}` | Get information about contest by id |
| :white_check_mark: | GET | `/contests/{contestId}/clarifications` | Get clarifications in contest by contest id |
| :white_check_mark: | GET | `/contests/{contestId}/messages` | Get messages in contest by contest id |
| &#9744; | POST | `/contests/{contestId}/messages` | Post new question to contest |
| :white_check_mark: | GET | `/contests/{contestId}/participants` | Get contest participants |
| :white_check_mark: | POST | `/contests/{contestId}/participants` | Register for contest |
| &#9744; | GET | `/contests/{contestId}/participants/{participantId}` | Get information about participant by id |
| &#9744; | PUT | `/contests/{contestId}/participants/{participantId}` | Start the contest for participant |
| &#9744; | DELETE | `/contests/{contestId}/participants/{participantId}` | Unregister user from contest by id |
| &#9744; | PATCH | `/contests/{contestId}/participants/{participantId}` | Update participant in contest |
| &#9744; | GET | `/contests/{contestId}/participants/{participantId}/stats` | Get participant stats |
| &#9744; | GET | `/contests/{contestId}/participation` | Get information about your status |
| &#9744; | PUT | `/contests/{contestId}/participation` | Start the contest |
| &#9744; | DELETE | `/contests/{contestId}/participation` | Unregister yourself from contest |
| &#9744; | GET | `/contests/{contestId}/problems` | Get contest's problems by contest id |
| &#9744; | GET | `/contests/{contestId}/standings` | Get contest's standings by contest id |
| &#9744; | GET | `/contests/{contestId}/standings/my` | Get user's standings by contest id |
| &#9744; | GET | `/contests/{contestId}/submissions` | Get your submissions to contest by contest id |
| &#9744; | POST | `/contests/{contestId}/submissions` | Send a submission to contest |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}` | Get detailed information about submission |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}/{testName}/answer` | Get full answer file in contest by test name |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}/{testName}/input` | Get full input file in contest by test name |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}/{testName}/output` | Get output file in contest by test name |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}/full` | Get full information about submission |
| &#9744; | GET | `/contests/{contestId}/submissions/{runId}/source` | Get source code of submission |
| &#9744; | HEAD | `/contests/{contestId}/submissions/{runId}/source` | Get source code of submission |
| &#9744; | POST | `/contests/{contestId}/submissions/lazy` | Send a submission to contest lazy |
| &#9744; | GET | `/contests/{contestId}/submissions/multiple` | Get full information about multiple submissions |
| &#9744; | GET | `/contests/neurips/{contestId}/submissions/all` | Get all contest submissions report |
| &#9744; | GET | `/contests/neurips/{contestId}/submissions/my` | Get your submissions to contest by contest id |
| &#9744; | GET | `/participants/{participantId}/stats` | Get participant stats |
| &#9744; | GET | `/service/capacity Get contest submissions` | queue capacity |
| :white_check_mark: | POST | `/submissions/{submissionId}/rejudge` | Rejudge submission |
| :white_check_mark: | POST | `/user/{userId}/generate-password` | Generate new password for internal user |
