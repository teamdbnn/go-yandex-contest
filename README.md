# go-yandex-contest
yandex contest api client (built from swagger https://api.contest.yandex.net/api/public/swagger-ui.html)



## Methods
| &#9744; | Method | Path  | Description |
|-----|--------|-------|-------------|
| &#9744; | GET | `/competitions/{competitionId}/participants` | Get registered participants of competition |
| &#9744; | POST | `/competitions/{competitionId}/participants` | Register participants into competition |
| &#9744; | GET | `/compilers` | Get compilers list |
| :white_check_mark: | GET | `/contests/{contestId}` | Get information about contest |
| &#9744; | POST | `/groups/` | Create a new group |
| &#9744; | GET | `/groups/{groupId}` | View group |
| &#9744; | POST | `/groups/{groupId}/members` | Add group member |
| &#9744; | DELETE | `/groups/{groupId}/members` | Remove group member |
| ? | POST | `/user/{userId}/generate-password` | Generate new password for internal user |
| &#9744; | GET | `/contests/{contestId}/clarifications` | Get jury clarifications |
| &#9744; | GET | `/contests/{contestId}/messages` | Get your questions and jury answers |
| &#9744; | POST | `/contests/{contestId}/messages` | Send question to jury |
| &#9744; | GET | `/contests/neurips/{contestId}/submissions/all` | Get all submissions for contest |
| &#9744; | GET | `/contests/neurips/{contestId}/submissions/my` | Get your submissions for contest |
| :white_check_mark: | GET | `/contests/{contestId}/groups` | List groups registered for contest |
| &#9744; | POST | `/contests/{contestId}/groups/{groupId}` | Register group for contest |
| &#9744; | DELETE | `/contests/{contestId}/groups/{groupId}` | Delete group for contest |
| &#9744; | PATCH | `/contests/{contestId}/groups/{groupId}` | Change group registration info |
| :white_check_mark: | GET | `/contests/{contestId}/participants` | Get contest participants |
| :white_check_mark: | POST | `/contests/{contestId}/participants` | Register for contest |
| &#9744; | GET | `/contests/{contestId}/participants/{participantId}` | Get information about participant |
| &#9744; | PUT | `/contests/{contestId}/participants/{participantId}` | Start the contest for participant |
| &#9744; | DELETE | `/contests/{contestId}/participants/{participantId}` | Unregister participant from contest |
| &#9744; | PATCH | `/contests/{contestId}/participants/{participantId}` | Update participant in contest |
| &#9744; | GET | `/contests/{contestId}/participation` | Get informantion about your participation |
| &#9744; | PUT | `/contests/{contestId}/participation` | Start the contest |
| &#9744; | DELETE | `/contests/{contestId}/participation` | Unregister yourself from contest |
| &#9744; | GET | `/contests/{contestId}/problems` | Get contest problems |
| &#9744; | GET | `/contests/{contestId}/problems/{alias}/statement` | Get problem statement |
| &#9744; | GET | `/problems` | Get problem file |
| &#9744; | GET | `/service/capacity` | Get submissoins queue capacity |
| &#9744; | GET | `/service/introspect` | Get avaible scopes |
| &#9744; | GET | `/contests/{contestId}/standings` | Get contest standings |
| &#9744; | GET | `/contests/{contestId}/standings-extended` | Get contest standings |
| &#9744; | GET | `/contests/{contestId}/standings/my` | Get your position in contest standings |
| &#9744; | GET | `/contests/{contestId}/submissions` | Get submissions for contest |
| &#9744; | POST | `/contests/{contestId}/submissions` | Send submission |
| &#9744; | POST | `/contests/{contestId}/submissions/lazy` | Send submission from URL |
| &#9744; | GET | `/contests/{contestId}/submissions/multiple` | Get report for multiple submissions |
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}` | Get brief report for submission |
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}/full` | Get full report for submission |
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}/source` | Get submission source code |
| &#9744; | HEAD | `/contests/{contestId}/submissions/{submissionId}/source` | Get metadata of submission source code|
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}/{testName}/answer` | Get full answer file for test |
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}/{testName}/input` | Get full input file for test |
| &#9744; | GET | `/contests/{contestId}/submissions/{submissionId}/{testName}/output` | Get participant output for test |
| ? | POST | `/submissions/{submissionId}/rejudge` | Rejudge submission | 
| &#9744; | GET | `/teams` | Get user teams |

